package mains
//#include "/home/wuchaochao/go/clib/hello.c"
import "C"
func main(){
	C.SayHello(C.CString("Hello.World\n"))
}
