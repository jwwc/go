/*package main
import "fmt"
func main(){
	var ch chan int = make(chan int,4)
	for i:=0;i<cap(ch);i++{
		ch<-i
	}
	for len(ch)>0{
		var value int = <-ch
		fmt.Println(value)
	}
}*/
//读写阻塞
/*package main
import "fmt"
import "time"
import "math/rand"

func send(ch chan int){
 for {
	 var value = rand.Intn(100)
	 ch<-value
	 fmt.Printf("send %d\n",value)
	}
}
func recv (ch chan int){
 for{
	 value :=<-ch
	 fmt.Printf("recv %d\n",value)
	 time.Sleep(time.Second)
    }
}
func main(){
	var ch = make(chan int ,1)
	go recv(ch)
	send(ch)
}*/
//关闭通道
package main

import "fmt"

func main() {
 var ch = make(chan int, 4)
 ch <- 1
 ch <- 2
 close(ch)

 value := <- ch
 fmt.Println(value)
 value = <- ch
 fmt.Println(value)
 value = <- ch
 fmt.Println(value)
}
