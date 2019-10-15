// package main
//
// import (
// 	"fmt"
// )
// // import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
// // 当标识符（包括常量、变量 、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
// // 那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public)
// // 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
// // go 语言的架构模型，go语言都是以包的形式进行管理的，有一个或者是多个工作目录gopath,这个目录下可以创建很多个文件夹，每一个文件下
// // 可以创建很多的go程序，go程序的导入是输入送所在文件夹的文件夹名
// //可以使用go env命令来查看go中环境变量的参数
// // 使用go build来构件经行编译成exe文件
// // 可以使用setx来设置环境变量
// //Get-ChildItem Env:用于显示注册表中环境变量
// //Set-Item -Path Env:Path -Value ($Env:Path + ";C:\Temp")用于修改注册表中环境变量
// //powershell的官网：https://docs.microsoft.com
// //学习之路：https://github.com/developer-learning/learning-golang/blob/master/README.md
// // 浅拷贝只复制指向某个对象的指针，而不复制对象本身，新旧对象还是共享同一块内存。
// // 但深拷贝会另外创造一个一模一样的对象，新对象跟原对象不共享内存，修改新对象不会改到原对象
// type Rect struct {
//     Width int
//     Height int
//
// }
// func main() {
//     var a interface {}
//     var r = Rect{50, 50}
//     a = &r // 指向了结构体指针
//
//     var rx = a.(*Rect) // 转换成指针类型
//     r.Width = 100
//     r.Height = 100
//     fmt.Println(rx)
// }
//
// 	/* 这是我的第一个简单的程序 */
// 	// fmt.Println("Hello World!")
// 	// fmt.Println(mathclass.Add(1, 1))
// 	// fmt.Println(mathclass.Sub(1, 1))
// 	//fmt.Println(stringutil.Reverse("!oG ,olleH"))
//   //fmt.Printf(trans())
// // func trans() string{
// // type Stringer interface {
// // 	String() string
// //  }
// //   var value interface{} // 调用者提供的值。
// //   switch str := value.(type) {
// //   case string:
// //     log.Printf("1")
// //     return str
// //   case Stringer:
// //     log.Printf("2")
// //     return str.String()
// //   }
// //   fmt.Printf(type(str))
// //   return "end"
// // }
