package main

import "fmt"
import "unsafe"

func main() {
    // head = {address, 10, 10}
    // body = [1,2,3,4,5,6,7,8,9,10]
    var s = []int{1,2,3,4,5,6,7,8,9,10}
    var address = (**[10]int)(unsafe.Pointer(&s))
    var len1 = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
    var cap = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
    fmt.Println(address, *len1, *cap)
    var body = **address
    for i:=0; i< len(body); i++ {
        fmt.Printf("%d ", body[i])
    }
    fmt.Println(s[0],s[1],s[2])
}

