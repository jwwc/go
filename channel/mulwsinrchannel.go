package main
import "fmt"
import "time"
import "sync"

func send(ch chan int,wg *sync.WaitGroup){
	defer wg.Done()//计数值减1
	i :=0
	for i<4{
		i++
		ch<-i
	}
}
func recv(ch chan int){
	for v:= range ch{
		fmt.Println(v)
	}
}
func main(){
	var ch = make(chan int ,4)
	var wg = new(sync.WaitGroup)
	wg.Add(2)
	go send(ch,wg)
	go send(ch,wg)
	go recv(ch)
	// Wait()阻塞等待所有的写通道携程结束之后
	//待计数值变为零，Wait()才返回
	wg.Wait()
	close(ch)
	time.Sleep(time.Second)
}
