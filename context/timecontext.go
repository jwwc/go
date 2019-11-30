//任何可能被阻塞，或者需要很长的时间来完成的都应该有context.Context
//context主要用来在goroutine之间传递上下文信息，包括取消信号，超时时间，截止时间，
// context几乎是并发控制和超时控制的一种标准做法
/*type Context interface {
    // 当 context 被取消或者到了 deadline，返回一个被关闭的 channel
    Done() <-chan struct{}

    // 在 channel Done 关闭后，返回 context 取消原因
    Err() error

    // 返回 context 是否会被取消以及自动取消时间（即 deadline）
    Deadline() (deadline time.Time, ok bool)

    // 获取 key 对应的 value
    Value(key interface{}) interface{}
}*/
/*type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() <-chan struct{}
}
Done() 方法返回一个只读的 channel，所有相关函数监听此 channel。一旦 channel 关闭，通过 channel 的“广播机制”，所有监听者都能收到。
*/
/*cancelCtx
再来看一个重要的 context：

type cancelCtx struct {
    Context

    // 保护之后的字段
    mu       sync.Mutex
    done     chan struct{}
    children map[canceler]struct{}
    err      error
}
这是一个可以取消的 Context，实现了 canceler 接口。它直接将接口 Context 作为它的一个匿名字段，这样，它就可以被看成一个 Context。

先来看 Done() 方法的实现：

func (c *cancelCtx) Done() <-chan struct{} {
    c.mu.Lock()
    if c.done == nil {
        c.done = make(chan struct{})
    }
    d := c.done
    c.mu.Unlock()
    return d
}*/
package main

import (
	"fmt"
	"context"
	"time"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return ctx.Err()
		}
	}
}
func main() {
	ctx,_:= context.WithTimeout(context.Background(),
		(time.Second/10000))
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	time.Sleep(time.Second)
	//cancel()
	wg.Wait()
}
