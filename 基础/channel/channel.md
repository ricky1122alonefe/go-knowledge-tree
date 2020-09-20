channel 基本结构
type hchan struct {
    qcount   uint // 当前缓存数据的总量  
    dataqsiz uint // 缓存数据的容量      
    buf      unsafe.Pointer // 缓存数据，为一个循环数组，容量大小为 dataqsiz，当前大小为 qcount
    elemsize uint16 // 数据类型的大小，比如 int 为 4
    closed   uint32 // 标记是否关闭
    elemtype *_type // 数据的类型
    sendx    uint  // 发送队列 sendq 的长度
    recvx    uint  // 接收队列 recvq 的长度
    recvq    waitq // 阻塞的接收 goroutine 的队列
    sendq    waitq // 阻塞的发送 goroutine 的队列
    lock mutex     // 锁，用于并发控制队列操作
}
waitq 为双向链表，sudog 代表一个封装的 goroutine，其参数 g 为 goroutine
type waitq struct {
    first *sudog
    last  *sudog
}

func makechan(t *chantype, size int) *hchan {
    elem := t.elem
    // 安全检查，数据项大小不超过 16K
    if elem.size >= 1<<16 {
        throw("makechan: invalid channel element type")
    }
    if hchanSize%maxAlign != 0 || elem.align > maxAlign {
        throw("makechan: bad alignment")
    }
    // 获取要分配的内存
    mem, overflow := math.MulUintptr(elem.size, uintptr(size))
    if overflow || mem > maxAlloc-hchanSize || size < 0 {
        panic(plainError("makechan: size out of range"))
    }
    var c *hchan
    switch {
    case mem == 0:
        // size 为 0 的情况，分配 hchan 结构体大小的内存，64位系统为 96 Byte.
        c = (*hchan)(mallocgc(hchanSize, nil, true))
        c.buf = c.raceaddr()
    case elem.kind&kindNoPointers != 0:
        // 数据项不为指针类型，调用 mallocgc 一次性分配内存大小，hchan 结构体大小 + 数据总量大小
        c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
        c.buf = add(unsafe.Pointer(c), hchanSize)
    default:
        // 数据项为指针类型，hchan 和 buf 分开分配内存，GC 中指针类型判断 reachable and unreadchable.
        c = new(hchan)
        c.buf = mallocgc(mem, elem, true)
    }
    // chan 赋值属性, 数据项大小、数据项类型、缓存数据的容量
    c.elemsize = uint16(elem.size)
    c.elemtype = elem
    c.dataqsiz = uint(size)
    return c
}


channel 发送 与 取值
<- func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool
-> func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool)


退出问题
for range 模式  range 遍历条件可以感知 channel 是否关闭
它在并发中的使用场景是：
    当协程只从1个channel读取数据，然后进行处理，处理后协程退出。下面这个示例程序，当in通道被关闭时，协程可自动退出。