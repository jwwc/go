//如果一个接口中没有定义任何方法，那么它就是空接口，任意结构体都能够隐式的实现空
//接口
// go语言为了用户重复定义很多的空接口，它自己内置了一个空接口，interface{}
// 空接口没有任何的能力，可以容纳任意对象，就像一个万能的容器
package main
import "fmt"
func main(){
	var user = map[string]interface{}{
		"age":30,
		"address":"Beijing Tongzhou",
		"married":true,
	}
	fmt.Println(user)
	var age = user["age"].(int)
	var address = user["address"].(string)
	var married = user["married"].(bool)
	fmt.Println(age,address,married)
}
