//reflect 包定义了十几种内置的「元类型」，每一种元类型都有一个整数编号，这个编号使用 reflect.Kind 类型表示。
//不同的结构体是不同的类型，但是它们都是同一个元类型 Struct。
//包含不同子元素的切片也是不同的类型，但是它们都会同一个元类型 Slice。
// https://zhuanlan.zhihu.com/p/53114706
// type Kind uint
//
// const (
//     Invalid Kind = iota // 不存在的无效类型
//     Bool
//     Int
//     Int8
//     Int16
//     Int32
//     Int64
//     Uint
//     Uint8
//     Uint16
//     Uint32
//     Uint64
//     Uintptr // 指针的整数类型，对指针进行整数运算时使用
//     Float32
//     Float64
//     Complex64
//     Complex128
//     Array // 数组类型
//     Chan // 通道类型
//     Func  // 函数类型
//     Interface  // 接口类型
//     Map // 字典类型
//     Ptr // 指针类型
//     Slice // 切片类型
//     String // 字符串类型
//     Struct // 结构体类型
//     UnsafePointer // unsafe.Pointer 类型
// )
// 反射的基础代码
// reflect 包提供了两个基础反射方法，分别是 TypeOf() 和 ValueOf() 方法，分别用于获取变量的类型和值，定义如下
//
// func TypeOf(v interface{}) Type
// func ValueOf(v interface{}) Value
// 下面是一个简单的例子，对结构体变量进行反射
//
// package main
//
// import "fmt"
// import "reflect"
//
// func main() {
//     var s int = 42
//     fmt.Println(reflect.TypeOf(s))
//     fmt.Println(reflect.ValueOf(s))
// }
// reflect.Type
// 它是一个接口类型，里面定义了非常多的方法用于获取和这个类型相关的一切信息。
// 这个接口的结构体实现隐藏在 reflect 包里，每一种类型都有一个相关的类型结构体来表达它的结构信息。
//
// type Type interface {
//   ...
//   Method(i int) Method  // 获取挂在类型上的第 i'th 个方法
//   ...
//   NumMethod() int  // 该类型上总共挂了几个方法
//   Name() string // 类型的名称
//   PkgPath() string // 所在包的名称
//   Size() uintptr // 占用字节数
//   String() string // 该类型的字符串形式
//   Kind() Kind // 元类型
//   ...
//   Bits() // 占用多少位
//   ChanDir() // 通道的方向
//   ...
//   Elem() Type // 数组，切片，通道，指针，字典(key)的内部子元素类型
//   Field(i int) StructField // 获取结构体的第 i'th 个字段
//   ...
//   In(i int) Type  // 获取函数第 i'th 个参数类型
//   Key() Type // 字典的 key 类型
//   Len() int // 数组的长度
//   NumIn() int // 函数的参数个数
//   NumOut() int // 函数的返回值个数
//   Out(i int) Type // 获取函数 第 i'th 个返回值类型
//   common() *rtype // 获取类型结构体的共同部分
//   uncommon() *uncommonType // 获取类型结构体的不同部分
// }
// package main
//
// import "reflect"
// import "fmt"
//
// func main() {
//     type SomeInt int
//     var s SomeInt = 42
//     var t = reflect.TypeOf(s)
//     var v = reflect.ValueOf(s)
//     // reflect.ValueOf(s).Type() 等价于 reflect.TypeOf(s)
//     fmt.Println(t == v.Type())
//     fmt.Println(v.Kind() == reflect.Int) // 元类型
//     // 将 Value 还原成原来的变量
//     var is = v.Interface()
//     fmt.Println(is.(SomeInt))
// }
package main
import "fmt"
import "reflect"

func main() {
    var s int = 42
    var v = reflect.ValueOf(s)
    v.SetInt(int64(s))
    fmt.Println(s)
}

// ---------
// panic: reflect: reflect.Value.SetInt using unaddressable value
//
// goroutine 1 [running]:
// reflect.flag.mustBeAssignable(0x82)
//     /usr/local/go/src/reflect/value.go:234 +0x157
// reflect.Value.SetInt(0x107a1a0, 0xc000016098, 0x82, 0x2b)
//     /usr/local/go/src/reflect/value.go:1472 +0x2f
// main.main()
//     /Users/qianwp/go/src/github.com/pyloque/practice/main.go:8 +0xc0
// exit status 2
// 尝试通过反射来修改整型变量失败了，程序直接抛出了异常。下面我们来尝试通过反射来修改指针变量指向的值，这个是可行的。
//
// package main
//
// import "fmt"
// import "reflect"
//
// func main() {
//     var s int = 42
//     // 反射指针类型
//     var v = reflect.ValueOf(&s)
//     // 要拿出指针指向的元素进行修改
//     v.Elem().SetInt(43)
//     fmt.Println(s)
// }

//结构体也是值类型，也必须通过指针类型来修改。下面我们尝试使用反射来动态修改结构体内部字段的值。

// package main
//
// import "fmt"
// import "reflect"
//
// type Rect struct {
//     Width int
//     Height int
// }
//
// func SetRectAttr(r *Rect, name string, value int) {
//     var v = reflect.ValueOf(r)
//     var field = v.Elem().FieldByName(name)
//     fmt.Println(field)
//     field.SetInt(int64(value))
// }
//
// func main() {
//     var r = Rect{50, 100}
//     SetRectAttr(&r, "Width", 100)
//     SetRectAttr(&r, "Height", 200)
//     fmt.Println(r)
// }
