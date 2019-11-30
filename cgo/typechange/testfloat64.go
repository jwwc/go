package main
import "fmt"
func main(){
	// []float 强制转换为 []int
	var a = []float64{4,2,5,7,1,1,88,1}
	var b []int = ((*[1 <<20 ]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	// 以int方式给float64排序
	sort.Ints(b)
	// float64遵循IEEE754浮点数标准特性
	// 当浮点数有序时对应的整数也必然有序
}
