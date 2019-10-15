// 不同的并行协程之间交流的方式有两种，一种是通过共享变量，另一种是通过队列。Go 语言鼓励使用队列的形式来交流，
// 它单独为协程之间的队列数据交流定制了特殊的语法 —— 通道。
// https://zhuanlan.zhihu.com/p/51710515
//创建通道
//创建通道只有一种语法，那就是 make 全局函数，提供第一个类型参数限定通道可以容纳的数据类型，再提供第二个整数参数作为通道的容器大小。
//大小参数是可选的，如果不填，那这个通道的容量为零，叫着「非缓冲型通道」，非缓冲型通道必须确保有协程正在尝试读取当前通道，
//否则写操作就会阻塞直到有其它协程来从通道中读东西。非缓冲型通道总是处于既满又空的状态。与之对应的有限定大小的通道就是缓冲型通道。
//在 Go 语言里不存在无界通道，每个通道都是有限定最大容量的。

// 缓冲型通道，里面只能放整数
//var bufferedChannel = make(chan int, 1024)
// 非缓冲型通道
//var unbufferedChannel = make(chan int)
// 读写阻塞
//通道满了，写操作就会阻塞，协程就会进入休眠，直到有其它协程读通道挪出了空间，协程才会被唤醒。
//如果有多个协程的写操作都阻塞了，一个读操作只会唤醒一个协程。



//通道空了，读操作就会阻塞，协程也会进入睡眠，直到有其它协程写通道装进了数据才会被唤醒。如果有多个协程的读操作阻塞了，
//一个写操作也只会唤醒一个协程。
// package main
//
// import "fmt"
//
// func main() {
//  var ch chan int = make(chan int, 4)
//  for i:=0; i<cap(ch); i++ {
//   ch <- i   // 写通道
//  }
//  for len(ch) > 0 {
//   var value int = <- ch  // 读通道
//   fmt.Println(value)
//  }
// }
// package main
// import "fmt"
// import "time"
// import "math/rand"
//
//
// func send(ch chan int){
//   for{
//     var value =  rand.Intn(100)
//     ch <- value
//     fmt.Printf("send %d\n",value)
//   }
// }
//
//
// func recv(ch chan int){
//   for{
//      var value = <- ch
//      fmt.Printf("resv %d\n",value)
//      time.Sleep(time.Second)
//   }
// }
//
// func main() {
//  var ch = make(chan int, 1)
//  // 子协程循环读
//  go recv(ch)
//  // 主协程循环写
//  send(ch)
// }
// 关闭通道
// Go 语言的通道有点像文件，不但支持读写操作， 还支持关闭。读取一个已经关闭的通道会立即返回通道类型的「零值」，
// 而写一个已经关闭的通道会抛异常。如果通道里的元素是整型的，读操作是不能通过返回值来确定通道是否关闭的。
//
// package main
//
// import "fmt"
//
// func main() {
//  var ch = make(chan int, 4)
//  ch <- 1
//  ch <- 2
//  close(ch)
//
//  value := <- ch
//  fmt.Println(value)
//  value = <- ch
//  fmt.Println(value)
//  value = <- ch
//  fmt.Println(value)
// }
// for range 语法我们已经见了很多次了，它是多功能的，除了可以遍历数组、切片、字典，还可以遍历通道，取代箭头操作符。
// 当通道空了，循环会暂停阻塞，当通道关闭时，阻塞停止，循环也跟着结束了。当循环结束时，我们就知道通道已经关闭了。
//
// package main
//
// import "fmt"
//
// func main() {
//  var ch = make(chan int, 4)
//  ch <- 1
//  ch <- 2
//  close(ch)
//
//  // for range 遍历通道
//  for value := range ch {
//   fmt.Println(value)
//  }
// }
// 通道写安全
// 上面提到向一个已经关闭的通道执行写操作会抛出异常，这意味着我们在写通道时一定要确保通道没有被关闭。
//
// package main
//
// import "fmt"
//
// func send(ch chan int) {
//  i := 0
//  for {
//   i++
//   ch <- i
//  }
// }
//
// func recv(ch chan int) {
//  value := <- ch
//  fmt.Println(value)
//  value = <- ch
//  fmt.Println(value)
//  close(ch)
// }
//
// func main() {
//  var ch = make(chan int, 4)
//  go recv(ch)
//  send(ch)
// }
// 那如何确保呢？Go 语言并不存在一个内置函数可以判断出通道是否已经被关闭。即使存在这样一个函数，
// 当你判断时通道没有关闭，并不意味着当你往通道里写数据时它就一定没有被关闭，并发环境下，它是可能被其它协程随时关闭的。
//
//
//
// 确保通道写安全的最好方式是由负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。
//
// package main
//
// import "fmt"
//
// func send(ch chan int) {
//  ch <- 1
//  ch <- 2
//  ch <- 3
//  ch <- 4
//  close(ch)
// }
//
// func recv(ch chan int) {
//  for v := range ch {
//   fmt.Println(v)
//  }
// }
//
// func main() {
//  var ch = make(chan int, 1)
//  go send(ch)
//  recv(ch)
// }
//这个方法确实可以解决单写多读的场景，可要是遇上了多写单读的场合该怎么办呢？任意一个读写通道的协程都不可以随意关闭通道，
//否则会导致其它写通道协程抛出异常。这时候就必须让其它不相干的协程来干这件事，这个协程需要等待所有的写通道协程都结束运行后才能关闭通道。
//那其它协程要如何才能知道所有的写通道已经结束运行了呢？这个就需要使用到内置 sync 包提供的 WaitGroup 对象，它使用计数来等待指定事件完成。

// package main
//
// import "fmt"
// import "time"
// import "sync"
//
// func send(ch chan int, wg *sync.WaitGroup) {
//  defer wg.Done() // 计数值减一
//  i := 0
//  for i < 4 {
//   i++
//   ch <- i
//  }
// }
//
// func recv(ch chan int) {
//  for v := range ch {
//   fmt.Println(v)
//  }
// }
//
// func main() {
//  var ch = make(chan int, 4)
//  var wg = new(sync.WaitGroup)
//  wg.Add(2) // 增加计数值
//  go send(ch, wg)  // 写
//  go send(ch, wg)  // 写
//  go recv(ch)
//  // Wait() 阻塞等待所有的写通道协程结束
//  // 待计数值变成零，Wait() 才会返回
//  wg.Wait()
//  // 关闭通道
//  close(ch)
//  time.Sleep(time.Second)
// }
//多路通道
//在真实的世界中，还有一种消息传递场景，那就是消费者有多个消费来源，只要有一个来源生产了数据，
//消费者就可以读这个数据进行消费。这时候可以将多个来源通道的数据汇聚到目标通道，然后统一在目标通道进行消费。

// package main
//
// import "fmt"
// import "time"
//
// // 每隔一会生产一个数
// func send(ch chan int, gap time.Duration) {
//  i := 0
//  for {
//   i++
//   ch <- i
//   time.Sleep(gap)
//  }
// }
//
// // 将多个原通道内容拷贝到单一的目标通道
// func collect(source chan int, target chan int) {
//  for v := range source {
//   target <- v
//  }
// }
//
// // 从目标通道消费数据
// func recv(ch chan int) {
//  for v := range ch {
//   fmt.Printf("receive %d\n", v)
//  }
// }
//
//
// func main() {
//  var ch1 = make(chan int)
//  var ch2 = make(chan int)
//  var ch3 = make(chan int)
//  go send(ch1, time.Second)
//  go send(ch2, 2 * time.Second)
//  go collect(ch1, ch3)
//  go collect(ch2, ch3)
//  recv(ch3)
// }
// 但是上面这种形式比较繁琐，需要为每一种消费来源都单独启动一个汇聚协程。Go 语言为这种使用场景带来了「多路复用」语法糖，
// 也就是下面要讲的 select 语句，它可以同时管理多个通道读写，如果所有通道都不能读写，它就整体阻塞，只要有一个通道可以读写，它就会继续。
// 下面我们使用 select 语句来简化上面的逻辑
// p非阻塞读写
// 前面我们讲的读写都是阻塞读写，Go 语言还提供了通道的非阻塞读写。当通道空时，读操作不会阻塞，
// 当通道满时，写操作也不会阻塞。非阻塞读写需要依靠 select 语句的 default 分支。
// 当 select 语句所有通道都不可读写时，如果定义了 default 分支，那就会执行 default 分支逻辑，这样就起到了不阻塞的效果。
// 下面我们演示一个单生产者多消费者的场景。生产者同时向两个通道写数据，写不进去就丢弃。
// package main
//
// import "fmt"
// //import "time"
//
// func send(ch1 chan int, ch2 chan int) {
//  i := 0
//  for {
//   i++
//   select {
//    case ch1 <- i:
//     fmt.Printf("send ch1 %d\n", i)
//    case ch2 <- i:
//     fmt.Printf("send ch2 %d\n", i)
//    default:
//   }
//  }
// }
//
// func recv(ch chan int,name string) {
//  for v := range ch {
//   fmt.Printf("receive %s %d\n", name, v)
//   //time.Sleep(gap)
//  }
// }
//
// func main() {
//  // 无缓冲通道
//  var ch1 = make(chan int)
//  var ch2 = make(chan int)
//  // 两个消费者的休眠时间不一样，名称不一样
//  go recv(ch1,"ch1")
//  go recv(ch2,"ch2")
//  send(ch1, ch2)
// }
//从输出中可以明显看出有很多的数据都丢弃了，消费者读到的数据是不连续的。如果将 select 语句里面的 default 分支干掉，再运行一次
//发送和接收的数据就连续了
