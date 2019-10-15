package main

import "fmt"
import "time"
import "runtime"
// func main(){
//    fmt.Println("run in main goroutine")
//    go func(){
//       fmt.Println("run in child goroutine")
//         go func(){
//            fmt.Println("run in grand child goroutine")
//             go func(){
//               fmt.Println("run in grand grand child goroutine")
//             }()
//         }()
//    }()
//    //time.Sleep(time.Second)
//    fmt.Println("main goroutine will quit")
// }
// 子协程异常退出
// 在使用子协程时一定要特别注意保护好每个子协程，确保它们正常安全的运行。
// 因为子协程的异常退出会将异常传播到主协程，直接会导致主协程也跟着挂掉，然后整个程序就崩溃了。
// 为了保护子协程的安全，通常我们会在协程的入口函数开头增加 recover()
//语句来恢复协程内部发生的异常，阻断它传播到主协程导致程序崩溃。recover 语句必须写在 defer 语句里面。

// func main() {
//     fmt.Println("run in main goroutine")
//     i := 1
//     for {
//         go func() {
//             for {
//                 time.Sleep(time.Second)
//             }
//         }()
//         if i % 10000 == 0 {
//             fmt.Printf("%d goroutine started\n", i)
//         }
//         i++
//     }
// }
//操作系统对线程的调度是抢占式的，也就是说单个线程的死循环不会影响其它线程的执行，每个线程的连续运行受到时间片的限制。
// Go 语言运行时对协程的调度并不是抢占式的。如果单个协程通过死循环霸占了线程的执行权，那这个线程就没有机会去运行其它协程了，
// 你可以说这个线程假死了。不过一个进程内部往往有多个线程，假死了一个线程没事，全部假死了才会导致整个进程卡死。
func main() {
    fmt.Println("run in main goroutine")
    n := 5
    //获取当前线程数
    fmt.Println(runtime.GOMAXPROCS(0))
    //  // 设置线程数为 10
    runtime.GOMAXPROCS(10)
    for i:=0; i<n; i++ {
        go func() {
            fmt.Println("dead loop goroutine start")
            for {}  // 死循环
        }()
    }
    for {
        time.Sleep(time.Second)
        fmt.Println("main goroutine running")
    }
}
