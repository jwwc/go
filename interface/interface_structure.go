// interface structure
/* type iface struct{
	tab *itab 类型指针
	data unsafe.Pointer 数据指针
}
type itab struct{
	inter *interfacetype 接口类新的信息
	_type *_type 数据类型
}*/
package main
import(
	"fmt"
	"unsafe"
)
func main(){
	var s interface{}
	fmt.Println(unsafe.Sizeof(s))
	var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(unsafe.Sizeof(arr))
	s = arr
	fmt.Println(unsafe.Sizeof(s))
}
