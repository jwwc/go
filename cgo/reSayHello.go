package main
// #include "./hello.h"
import "C"
import "fmt"
func main(){
    C.SayHello(C.CString("hello,world\n"))
}
//export SayHello
func SayHello(s *C.char){
	fmt.Print(C.GoString(s))
}
