### 
type iface struct {
  // 包含接口的静态类型信息、数据的动态类型信息、函数表
  tab  *itab
  // 指向具体数据的内存地址比如slice、map等，或者在接口
  // 转换时直接存放小数据(一个指针的长度)
  data unsafe.Pointer
}

type itab struct {
  // 接口的类型信息
  inter  *interfacetype
  // 具体数据的类型信息
  _type  *_type
  link   *itab
  hash   uint32
  bad    bool 
  inhash bool
  unused [2]byte
  // 函数地址表，这里放置和接口方法对应的具体数据类型的方法地址
  // 实现接口调用方法的动态分派，一般在每次给接口赋值发生转换时
  // 会更新此表，或者直接拿缓存的itab
  fun    [1]uintptr // variable sized
}

在Go中只有值传递(包括接口类型)，与具体的类型实现无关，但是某些类型具有引用的属性。典型的9种非基础类型中:

array传递会拷贝整块数据内存，传递长度为len(arr) * Sizeof(elem)
string、slice、interface传递的是其runtime的实现，所以长度是固定的，分别为16、24、16字节(amd64)
map、func、chan、pointer传递的是指针，所以长度固定为8字节(amd64)
struct传递的是所有字段的内存拷贝，所以长度是所有字段的长度和
详细的测试可以参考[这段程序](pass_by_value_main.go)


接口相关的操作主要在于对其内部字段itab的操作，因为接口转换最重要的是类型信息
接口的类型转换在编译期会生成一个函数调用的语法树节点(OCALL)，调用runtime提供的相应接口转换函数完成接口的类型设置，所以接口的转换是在运行时发生的，其具体类型的方法地址表也是在运行时填写的
由于在运行时转换会产生开销，所以对转换的itab做了缓存


getitab  根据接口类型和实际数据类型生成itab 
先从缓存中查找对应的itab  2次查找
// 缓存中没找到则分配itab的内存: itab结构本身内存 + 末尾存方法地址表的可变长度
  m = (*itab)(persistentalloc(unsafe.Sizeof(itab{})+uintptr(len(inter.mhdr)-1)*sys.PtrSize, 0, &memstats.other_sys))
  m.inter = inter           // 设置接口类型信息
  m._type = typ             // 设置实际数据类型信息
  additab(m, true, canfail) // 设置itab函数调用表
