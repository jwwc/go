// package main
//
// import (
//   "fmt"
//   "unsafe"
// )
// https://zhuanlan.zhihu.com/p/52756600
// unsafe.Pointer
//Pointer 代表着变量的内存地址，可以将任意变量的地址转换成 Pointer 类型，
//也可以将 Pointer 类型转换成任意的指针类型，它是不同指针类型之间互转的中间类型。Pointer 本身也是一个整型的值。
//在 Go 语言里不同类型之间的转换是要受限的。普通的基础变量转换成不同的类型需要进行内存浅拷贝，
//而指针变量类型之间是禁止直接转换的。要打破这个限制，unsafe.Pointer 就可以派上用场，它允许任意指针类型的互转。

//指针的加减运算
// Pointer 虽然是整型的，但是编译器禁止它直接进行加减运算。如果要进行运算，需要将 Pointer 类型转换 uintptr 类型进行加减，
// 然后再将 uintptr 转换成 Pointer 类型。uintptr 其实也是一个整型。
// type Rect struct{
//    width int
//    height int
// }
//
// func main(){
//   var r = Rect{50,50}
//   var width = *(*int)(unsafe.Pointer(&r))
//   var height = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&r))+uintptr(8)))
//   fmt.Println(width,height)
// }
// package main
//
// import "fmt"
// import "unsafe"
//
// type Rect struct {
//     Width int
//     Height int
// }
//
// func main() {
//     var r = Rect {50, 50}
//     // var pw *int
//     var pw = (*int)(unsafe.Pointer(&r))
//     // var ph *int
//     var ph = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&r)) + uintptr(8)))
//     *pw = 100
//     *ph = 100
//     fmt.Println(r.Width, r.Height)
// }
// 代码中的 uintptr(8) 很不优雅，可以使用 unsafe 提供了 Offsetof 方法来替换它，它可以直接得到字段在结构体内的偏移量
//
// var ph = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&r)) + unsafe.Offsetof(r.Height))
// package main
//
// package main
// import "fmt"
// import "unsafe"
//
// func main() {
//     // head = {address, 10, 10}
//     // body = [1,2,3,4,5,6,7,8,9,10]
//     var s = []int{1,2,3,4,5,6,7,8,9,10}
//     var address = (**[10]int)(unsafe.Pointer(&s))
//     var lens = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
//     var cap = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
//     fmt.Println(address, *lens, *cap)
//     var body = **address
//     for i:=0; i< len(body); i++ {
//         fmt.Printf("%d ", body[i])
//     }
// }

// 字符串与字节切片的高效转换
// 在字符串小节我们提到字节切片和字符串之间的转换需要复制内存，如果字符串或者字节切片的长度较大，
// 转换起来会有较高的成本。下面我们通过 unsafe 包提供另一种高效的转换方法，让转换前后的字符串和字节切片共享内部存储。
// 字符串和字节切片的不同点在于头部，字符串的头部 2 个 int 字节，切片的头部 3 个 int 字节

// package main
//
// import "fmt"
// import "unsafe"
//
// func main() {
//     fmt.Println(bytes2str(str2bytes("hello")))
// }
//
// func str2bytes(s string) []byte {
//     var strhead = *(*[2]int)(unsafe.Pointer(&s))
//     var slicehead [3]int
//     slicehead[0] = strhead[0]
//     slicehead[1] = strhead[1]
//     slicehead[2] = strhead[1]
//     return *(*[]byte)(unsafe.Pointer(&slicehead))
// }
//
// func bytes2str(bs []byte) string {
//     return *(*string)(unsafe.Pointer(&bs))
// }
//切记通过这种形式转换而成的字节切片千万不可以修改，因为它的底层字节数组是共享的，修改会破坏字符串的只读规则。
//其次使用这种形式得到的字符串或者切片只可以用作临时的局部变量，因为被共享的字节数组随时可能会被回收，
//原字符串或者字节切片的内存由于不再被引用，让垃圾回收器解决掉了。
// package main
//
// import "fmt"
//
// func main(){
// 	x,y := 1, 2
// 	var arr =  [...]int{5:2}
// 	//数组指针
// 	var pf *[6]int = &arr
//
// 	//指针数组
// 	pfArr := [...]*int{x,y}
// 	fmt.Println(*pf)
// 	fmt.Println(pfArr)
// }
package main

import "fmt"
import "unsafe"

type Rect struct {
    Width int
    Height int
}

func main() {
    var r = Rect{50, 50}
    // {typeptr, dataptr}
    var s interface{} = r

    var sptrs = *(*[2]*Rect)(unsafe.Pointer(&s))
    //var dataptr *Rect
    var sdataptr = sptrs[1]
    fmt.Println(sdataptr.Width, sdataptr.Height)


    //修改原对象，看看接口指向的对象是否受到影响
    r.Width = 100
    fmt.Println(sdataptr.Width, sdataptr.Height)
}
// package main
//
// import "fmt"
// import "unsafe"
//
// type Rect struct {
//     Width int
//     Height int
// }
//
// func main() {
//     // {typeptr, dataptr}
//     var s interface{} = Rect{50, 50}
//     var r = s
//
//     var rptrs = *(*[2]*Rect)(unsafe.Pointer(&r))
//     var rdataptr = rptrs[1]
//     var sptrs = *(*[2]*Rect)(unsafe.Pointer(&s))
//     var sdataptr = sptrs[1]
//
//     fmt.Println(sdataptr.Width, sdataptr.Height)
//     fmt.Println(rdataptr.Width, rdataptr.Height)
//
//     // 修改原对象S
//     r.Width = 100
//     // 再对比一下原对象和目标对象
//     fmt.Println(sdataptr.Width, sdataptr.Height)
//     fmt.Println(rdataptr.Width, rdataptr.Height)
// }
