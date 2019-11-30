/*Cond的作用和WaitGroup是一样的，都是goroutine插入，不同的是WaitGroup是被动，所有goroutine跑完后，wait会会自动释放，而cond是主动中断，我们还必须给cond发送信号，来通知等待释放*/
/*type Cond struct{
	noCopy noCopy
	L Locker
	notify notifyList
	check copyChecker
}*/
/*func NewCond(l Locker) *Cond

func (c *Cond)Signal()

func(c *Cond)Broadcast()

func(c *Cond)Wait()*/
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    locker := new(sync.Mutex)
    cond := sync.NewCond(locker)
    done := false

    cond.L.Lock()

    go func() {
        time.Sleep(time.Second * 1)
        done = true
        cond.Signal()    // 发送信号，通知Wait()释放
    }()

    if !done {         cond.Wait()      // 堵塞主goroutine
    }

    fmt.Println("now done is", done)    //一秒钟后会打印出 now done is true
}
package main

import (
    "fmt"
    "sync"
    "time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func read(x int) {
    cond.L.Lock()    // 获取锁
    cond.Wait()      // 等待通知，暂时阻塞
    fmt.Println(x)
    time.Sleep(time.Second * 1)
    cond.L.Unlock()  // 释放锁，不释放的话将只会有一次输出
}

func main() {
    for i := 0; i < 40; i++ {
        go read(i)
    }
    fmt.Println("start all")
    time.Sleep(time.Second * 1)
    cond.Broadcast() // 下发广播给所有等待的goroutine
    time.Sleep(time.Second * 60)
}
