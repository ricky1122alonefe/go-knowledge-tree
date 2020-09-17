golang http server 中每一个请求都有一个对应的gouroutine 当一个请求超时或者取消,所有的goroutine都应该要退出 系统去释放他们

1不要将 Contexts 放入结构体，相反context应该作为第一个参数传入，命名为ctx。 func DoSomething（ctx context.Context，arg Arg）error { // ... use ctx ... }
2即使函数允许，也不要传入nil的 Context。如果不知道用哪种 Context，可以使用context.TODO()。
3使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数
4相同的 Context 可以传递给在不同的goroutine；Context 是并发安全的。


type Context interface {
    // Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
   	Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}

默认实现了 emptyCtx 的 contextInterface

1.web编程中，一个请求对应多个goroutine之间的数据交互
2.超时控制
3.上下文控制


为了更方便的创建Context，包里头定义了Background来作为所有Context的根，它是一个emptyCtx的实例。
var (
    background = new(emptyCtx)
    todo       = new(emptyCtx) // 
)

func Background() Context {
    return background
}

所有的Context是树的结构，Background是树的根，当任一Context被取消的时候，那么继承它的Context 都将被回收。