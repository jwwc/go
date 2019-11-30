package main

import "fmt"

var negErr = fmt.Errorf("non positive number")

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("error catched"， err)
        }
    }()
    fmt.Println(fact(10))
    fmt.Println(fact(5))
    fmt.Println(fact(-5))
    fmt.Println(fact(15))
}

func fact(a int) int{
    if a <= 0 {
        panic(negErr)
    }
    var r = 1
    for i :=1;i<=a;i++ {
        r *= i
    }
    return r
}
