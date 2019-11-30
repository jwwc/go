package main
import "C"
// int sum(int a,int b);
//export sum
func sum(a,b,C.int) C.int{
	return a+b
}
// C.xxx类型不能跨越多个包
// 1：因为C.xxx最终对应的_Ctype_xxx是一个内部类型,因此不同包之间的C.int并不是相同的类型
