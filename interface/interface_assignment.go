package main
import "fmt"
type Rect struct{
	Width int
	Height int
}
func main(){
	var a interface{}
	var r  =  Rect{50,50}
	a = r
	var rx = a.(Rect)
	r.Width = 100
	r.Height = 100
	fmt.Println(rx)
}
