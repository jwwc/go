/* 封装，结构内部对外部不可见
   继承：继承的主要功能是进行功能的扩充
   多态：接口的不同实现方式,子类和父类之间的上下转型
   方法的多态性：方法的重载
                 方法的重写
   对象的多态性：父子之间的上下转型*/
package main
import "fmt"
type Fruitable interface{
	eat()
}
type Fruit struct{
	Name string
	Fruitable //匿名嵌套接口变量
}
func (f Fruit) want(){
	fmt.Println("I like")
	f.eat()
}
type Apple struct{}
func (a Apple)eat(){
	fmt.Println("eating apple")
}
type Banana struct{}
func(b Banana)eat(){
	fmt.Println("eating banana")
}
func main(){
	var f1 = Fruit{"Apple",Apple{}}
	var f2 = Fruit{"Banana",Banana{}}
	f1.want()
	f2.want()
}
