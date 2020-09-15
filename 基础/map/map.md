```
type hmap struct {
    // map 中的元素个数，必须放在 struct 的第一个位置，因为 内置的 len 函数会从这里读取
    count     int 
    // map状态标识，比如是否在被写或者迁移等，因为map不是线程安全的所以操作时需要判断flags
    flags     uint8
    // log_2 of buckets (最多可以放 loadFactor * 2^B 个元素即6.5*2^B，再多就要 hashGrow 了)
    B         uint8  
    // overflow 的 bucket 的近似数
    noverflow uint16 
    // hash seed，随机哈希种子可以防止哈希碰撞攻击
    hash0     uint32
   
    // 存储数据的buckets数组的指针， 大小2^B，如果 count == 0 的话，可能是 nil
    buckets    unsafe.Pointer
    // 一半大小的之前的 bucket 数组，只有在 growing 过程中是非 nil
    oldbuckets unsafe.Pointer
    // 扩容进度标志，小于此地址的buckets已迁移完成。
    nevacuate  uintptr

    // 可以减少GC扫描，当 key 和 value 都可以 inline 的时候，就会用这个字段
    extra *mapextra // optional fields
}

type mapextra struct {
    // 如果 key 和 value 都不包含指针，并且可以被 inline(<=128 字节)
    // 使用 extra 来存储 overflow bucket，这样可以避免 GC 扫描整个 map
    // 然而 bmap.overflow 也是个指针。这时候我们只能把这些 overflow 的指针
    // 都放在 hmap.extra.overflow 和 hmap.extra.oldoverflow 中了
    // overflow 包含的是 hmap.buckets 的 overflow 的 bucket
    // oldoverflow 包含扩容时的 hmap.oldbuckets 的 overflow 的 bucket
    overflow       *[]*bmap
    oldoverflow    *[]*bmap

    // 指向空闲的 overflow bucket 的指针
    nextOverflow *bmap
}

type bmap struct {
    // tophash 是 hash 值的高 8 位
    tophash [bucketCnt]uint8
    // 以下字段没有显示定义在bmap，但是编译时编译器会自动添加
    // keys              // 每个桶最多可以装8个key
    // values            // 8个key分别有8个value一一对应
    // overflow pointer  // 发生哈希碰撞之后创建的overflow bucket
}
```
----
### map 查找
```
 1 根据key计算出哈希值
 2.根据哈希值低位确定所在bucket
 3.根据哈希值高8位确定在bucket中的存储位置
---
```
### map 插入
```
    根据key计算出哈希值
    根据哈希值低位确定所在bucket
    根据哈希值高8位确定在bucket中的存储位置
    查找该key是否存在，已存在则更新，不存在则插入
```
### golang中map的无序性
```
mapiterinit 遍历方法中 加入了fastrand 随机数
	r := uintptr(fastrand())
	if h.B > 31-bucketCntBits {
		r += uintptr(fastrand()) << 31
	}
	it.startBucket = r & bucketMask(h.B)
	it.offset = uint8(r >> h.B & (bucketCnt - 1))
```
### golang中map扩容
```
 当前kv内容存储过多的情况

 将当前的hmap.buckets 赋值给oldbuckets
 oldbuckets := h.buckets

 // 申请一个大数组，作为新的buckets
 newbuckets, nextOverflow := makeBucketArray(t, h.B+bigger, nil)

 // 然后重新赋值map的结构体，oldbuckets 被填充。之后将做搬迁操作
  h.B += bigger
  h.flags = flags
  h.oldbuckets = oldbuckets
  h.buckets = newbuckets
  h.nevacuate = 0
  h.noverflow = 0
申请新的数组后 对原先的 buckets overflow 做交换

高位用于寻找bucket中的哪个key
低位用于寻找当前key属于hmap中的哪个bucket

1.先要判断当前bucket是不是已经转移。 (oldbucket 标识需要搬迁的bucket 对应的位置)
2. 如果没有被转移，那就要迁移数据了
3. 确定bucket位置后，需要按照kv 一条一条做迁移

Map 的赋值难点在于数据的扩容和数据的搬迁操作。
bucket 搬迁是逐步进行的，每进行一次赋值，会做至少一次搬迁工作。
扩容不是一定会新增空间，也有可能是只是做了内存整理。
tophash 的标志即可以判断是否为空，还会判断是否搬迁，以及搬迁的位置为X or Y。
delete map 中的key，有可能出现很多空的kv，会导致搬迁操作。如果可以避免，尽量避免。
golang delete map中的key 并不会直接释放内存 
map = nil 之后 gc 会将对应内存释放
```
golang中map并发并非安全
syncMap