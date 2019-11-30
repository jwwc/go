/*
include <errno.h>
static void seterrno(int v){
	errno = v;
}
*/
import "C"
import "fmt"

func main(){
	_,err := C.Seterrno(9527)
	fmt.Println(err)
	// output
	// errno 9527
}
//即使没有返回值，依然可以通过第二个获取errno
//对应void函数类型，第一个返回值可以使用_占位
