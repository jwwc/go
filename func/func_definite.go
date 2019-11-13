package main
import "fmt"
func Add(a,b int)int{
	return a+b
}
//匿名函数
var add = func(a,b int)int{
	return a+b
}
//多个参数和多个返回
func Swap(a,b int)(int,int){
	return b,a
}
func Sum(a int,more ...int)int{
	for _,v:= range more{
		a+=v
	}
	return a
}
//接口类型的可变参数
func Print(a ...interface{}){
	fmt.Print(a...)
}
//参数返回值命名的函数
func Find(m map[int]int,key int)(value int,ok bool){
	value,ok = m[key]
	return
}
//如果参数返回值已经被命名了，也可以通过defer语句在return//之后修改返回值
func Inc()(v int){
	defer func(){v++}()
	return 42
}
/*通过传入i,defer 语句会马上对调用参数求值*/
func Inci(){
	for i:=0; i<3; i++{
		defer func(i int){
			println(i)}(i)
	}
}

func main(){
        var sum = Add(3,4)
	fmt.Println(sum)
	sum = add(2,4)
	fmt.Println(sum)
	var a,b = Swap(3,4)
	fmt.Println(a,b)
	sum = Sum(2,3,4,5,6)
	fmt.Println(sum)
	Print("dff","dfsa")
	var  vmap = map[int]int{
		2:4,
		3:6,
		4:8,
	}
	var value,ok  = Find(vmap,3)
	fmt.Println(value,ok)
	var v =Inc()
	fmt.Println(v)
	Inci()

}
