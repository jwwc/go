package main
import(
	"fmt"
	"time"
)
func send(ch chan int,gap time.Duration){
	i:=0
	for{
		i++
		ch<-i
		time.Sleep(gap)
	}
}
func collect(source chan int,target chan int){
	for v:= range source{
		target<-v
	}
}
func recv(ch chan int){
	for v :=range ch{
		fmt.Printf("receive %d\n",v)
	}
}
func main(){
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	go send(ch1,time.Second)
	go send(ch2,time.Second)
	go collect(ch1,ch3)
	go collect(ch2,ch3)
	recv(ch3)

}
//改进
/*select {
  case ch1 <- v:
      fmt.Println("send to ch1")
  case ch2 <- v:
      fmt.Println("send to ch2")
}*/
