package main

import (
	"fmt"
	"mymath"
)
// import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
// 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
// 那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public)
// 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
//go 文件的模式
func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello World!")
	fmt.Println(mathclass.Add(1, 1))
	fmt.Println(mathclass.Sub(1, 1))
}
