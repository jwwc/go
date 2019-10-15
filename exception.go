package main

import "fmt"
//import "os"
//import "errors"
// type error interface{
//    Error() string
// }
// type SomeError struct{
//    Reason string
// }
// func(er SomeError) Error() string{
//    return er.Reason
// }
// func main() {
//     var err error = SomeError{"something happened"}
//     fmt.Println(err)
// }
// package errors
//
// func New(text string) error {
//     return &errorString{text}
// }
//
// type errorString struct {
//     s string
// }
//
// func (e *errorString) Error() string {
//     return e.s
// }
// 在 Java 语言里，如果遇到 IO 问题通常会抛出 IOException 类型的异常，在 Go 语言里面它不会抛异常，
//而是以返回值的形式来通知上层逻辑来处理错误。下面我们通过读文件来尝试一下 Go 语言的错误处理，读文件需要使用内置的 os 包
// func main() {
//     // 打开文件
//     var f, err = os.Open("go-doc.txt")
//     if err != nil {
//         // 文件不存在、权限等原因
//         fmt.Println("open file failed reason:" + err.Error())
//         return
//     }
//     // 推迟到函数尾部调用，确保文件会关闭
//     defer f.Close()
//     // 存储文件内容
//     var content = []byte{}
//     // 临时的缓冲，按块读取，一次最多读取 100 字节
//     var buf = make([]byte, 100)
//     for {
//         // 读文件，将读到的内容填充到缓冲
//         n, err := f.Read(buf)
//         if n > 0 {
//             // 将读到的内容聚合起来
//             content = append(content, buf[:n]...)
//         }
//         if err != nil {
//             // 遇到流结束或者其它错误
//             break
//         }
//     }
//     // 输出文件内容
//     fmt.Println(string(content))
// }
// import "strconv"
// import "github.com/go-redis/redis"
//
// func main() {
//  // 定义客户端对象，内部包含一个连接池
//     var client = redis.NewClient(&redis.Options {
//         Addr: "localhost:6379",
//     })
//
//     // 定义三个重要的整数变量值，默认都是零
//     var val1, val2, val3 int
//
//     // 获取第一个值
//     valstr1, err := client.Get("value1").Result()
//     if err == nil {
//         val1, err = strconv.Atoi(valstr1)
//         if err != nil {
//             fmt.Println("value1 not a valid integer")
//             return
//         }
//     } else if err != redis.Nil {
//         fmt.Println("redis access error reason:" + err.Error())
//         return
//     }
//
//     // 获取第二个值
//     valstr2, err := client.Get("value2").Result()
//     if err == nil {
//         val2, err = strconv.Atoi(valstr2)
//         if err != nil {
//             fmt.Println("value1 not a valid integer")
//             return
//         }
//     } else if err != redis.Nil {
//         fmt.Println("redis access error reason:" + err.Error())
//         return
//     }
//
//     // 保存第三个值
//     val3 = val1 * val2
//     ok, err := client.Set("value3",val3, 0).Result()
//     if err != nil {
//         fmt.Println("set value error reason:" + err.Error())
//         return
//     }
//     fmt.Println(ok)
// }
//因为 Go 语言中不轻易使用异常语句，所以对于任何可能出错的地方都需要判断返回值的错误信息。上面代码中除了访问 Redis 需要判断之外，字符串转整数也需要判断。

//另外还有一个需要特别注意的是因为字符串的零值是空串而不是 nil，你不好从字符串内容本身判断出 Redis 是否存在这个 key 还是对应 key 的 value 为空串，
//需要通过返回值的错误信息来判断。代码中的 redis.Nil 就是客户端专门为 key 不存在这种情况而定义的错误对象

// var negErr = fmt.Errorf("non positive number")
//
// func main() {
//     fmt.Println(fact(10))
//     fmt.Println(fact(5))
//     fmt.Println(fact(-5))
//     fmt.Println(fact(15))
// }
//
// // 让阶乘函数返回错误太不雅观了
// // 使用 panic 会合适一些
// func fact(a int) int{
//     if a <= 0 {
//         panic(negErr)
//     }
//     var r = 1
//     for i :=1;i<=a;i++ {
//         r *= i
//     }
//     return r
// }

var negErr = fmt.Errorf("non positive number")

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("error catched", err)
        }
    }()
    fmt.Println(fact(10))
    fmt.Println(fact(5))
    fmt.Println(fact(-5))
    fmt.Println(fact(15))
}

func fact(a int) int{
    if a <= 0 {
        panic(negErr)
    }
    var r = 1
    for i :=1;i<=a;i++ {
        r *= i
    }
    return r
}
//还有个值得注意的地方时，panic 抛出的对象未必是错误对象，而 recover() 返回的对象正是 panic 抛出来的对象，所以它也不一定是错误对象。

//func panic(v interface{})
//func recover() interface{}

// 我们经常还需要对 recover() 返回的结果进行判断，以挑选出我们愿意处理的异常对象类型，
// 对于那些不愿意处理的，可以选择再次抛出来，让上层来处理。
//
// defer func() {
//     if err := recover(); err != nil {
//         if err == negErr {
//             fmt.Println("error catched", err)
//         } else {
//             panic(err)  // rethrow
//         }
//     }
// }()
// 有时候我们需要在一个函数里使用多次 defer 语句。
// 比如拷贝文件，需要同时打开源文件和目标文件，那就需要调用两次 defer f.Close()。
//需要注意的是 defer 语句的执行顺序和代码编写的顺序是反过来的，
//也就是说最先 defer 的语句最后执行，为了验证这个规则，我们来改写一下上面的代码
