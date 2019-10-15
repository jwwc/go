package main

import "fmt"
import "unsafe"
//字典变量里存的只是一个地址指针，这个指针指向字典的头部对象。所以字典变量占用的空间是一个字，
//也就是一个指针的大小，64 位机器是 8 字节，32 位机器是 4 字节
func main() {
    var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,
    }

    // var names = make([]string, 0, len(fruits))
    // var scores = make([]int, 0, len(fruits))
    //
    // for name, score := range fruits {
    //     names = append(names, name)
    //     scores = append(scores, score)
    // }
    for name, score := range fruits {
        fmt.Println(name, score)
    }

    for _ ,name := range fruits {
        fmt.Println(name)
    }

    //fmt.Println(names, scores)
    fmt.Println(unsafe.Sizeof(fruits))
}
