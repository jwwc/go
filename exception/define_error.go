package main

import(
	"fmt"
)
type SomeError struct{
    Reason string
}
func (er SomeError) Error() string{
	return er.Reason
}
func main(){
	var err = SomeError{"something happend"}
	fmt.Println(err)
}
