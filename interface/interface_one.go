package main
import "fmt"
type Smellable interface{
	smell()
}
type Eatable interface{
	eat()
}
type Apple struct{}
func(a Apple)smell(){
	fmt.Println("apple can smell")
}
func(a Apple)eat(){
	fmt.Println("apple can eat")
}
type Flower struct{}
func(f Flower)smell(){
	fmt.Println("flower can smell")
}
func main(){
	var s1 Smellable
	var s2 Eatable
	var apple = Apple{}
	var flower = Flower{}
	s1 = apple
	s1.smell()
	s1 = flower
	s1.smell()
	s2 = apple
	s2.eat()
}
