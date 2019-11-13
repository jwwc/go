package main
import "fmt"
type Smellable interface{
	smell()
}
type Eatable interface{
	eat()
}
type Fruitable interface{
	Smellable
	Eatable
}
type Apple struct{}
func (a Apple) smell(){
	fmt.Println("apple is able to semll")
}
func (a Apple) eat(){
	fmt.Println("apple is able to eat")
}
type Fruit struct{
	Fruitable
}
func main(){
	//var fi = Fruitable{Apple{},Apple{}}
	var f = Fruit{Apple{}}
	f.smell()
	f.eat()
}

