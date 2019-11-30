package main

import (
	"fmt"
	"time"
)

func Producer( out chan<- int) {
	for i := 0;i<100 ; i++ {
		out <- i 
	}
}

// 消费者
func Consumer(in <-chan int,user string) {
	for v := range in {
		fmt.Println(v,user)
	}
}
func main() {
	ch := make(chan int, 64) // 成果队列
	go Producer(ch)       // 生成 3 的倍数的序列
        go Producer(5, ch)       // 生成 5 的倍数的序列
	go Consumer(ch,"user1")
	go Consumer(ch,"user2")
	// 消费 生成的队列
	// 运行一定时间后退出
	time.Sleep(5 * time.Second)
}
