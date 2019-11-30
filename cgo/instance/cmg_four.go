import "C"
//export GoAdd
func GoAdd(a,b C.int) C.int{
	return a+b
}
//可以导出私有函数
//导出C函数名名字空间约束，需要保证全剧没有重名
//main包的导出函数会在_cgo_export.h_声明
