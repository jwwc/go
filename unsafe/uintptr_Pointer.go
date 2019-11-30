/*unsafe.Pointer
Pointer 代表着变量的内存地址，可以将任意变量的地址转换成 Pointer 类型，也可以将 Pointer 类型转换成任意的指针类型，它是不同指针类型之间互转的中间类型。Pointer 本身也是一个整型的值。

type Pointer int*/
/*在 Go 语言里不同类型之间的转换是要受限的。普通的基础变量转换成不同的类型需要进行内存浅拷贝，而指针变量类型之间是禁止直接转换的。要打破这个限制，unsafe.Pointer 就可以派上用场，它允许任意指针类型的互转。*/
/*指针的加减运算
Pointer 虽然是整型的，但是编译器禁止它直接进行加减运算。如果要进行运算，需要将 Pointer 类型转换 uintptr 类型进行加减，然后再将 uintptr 转换成 Pointer 类型。uintptr 其实也是一个整型。

type uintptr int*/
package main

import "fmt"
import "unsafe"

type Rect struct {
    Width int
    Height int
}

func main() {
    var r = Rect {50, 50}
    // *Rect => Pointer => *int => int
    var width = *(*int)(unsafe.Pointer(&r))
    // *Rect => Pointer => uintptr => Pointer => *int => int
    var height = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&r)) + uintptr(8)))
    fmt.Println(width, height)
}
