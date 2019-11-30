// export int Goadd(int a,int b)
//#include "add.h"
import "C"
func main(){
	C.GoAdd(1,1)
	C.c_add(2,2)
}
//无法在go中头文件,因为还没有生成
//GoAdd 是Go导出函数，无法通过_cgo_export.h引用
//c_add 是C定义的函数，还可以通过add.h头文件引用
//可以手写函数声明，不会形成循环依赖
