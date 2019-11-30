package main
// #include "hello.h"
import "C"
import "fmt"
func main() {
C.hello(C.int(12))
fmt.Println("Hello Go");
}
