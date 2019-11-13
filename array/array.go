package main
import "fmt"
type Image struct{
        point Point
}
type Point struct{
        X int
        Y int
}
func main(){
	//var image Image
	var line1 = [2]Point{{0,0},{1,1}}
        for _,Y := range line1 {
		fmt.Printf("%d==\n",Y)
        }
}

