package main
import(
	"fmt"
	"sync"
)
var total struct{
	sync.Mutex
	value int
}
func worker(wg *sync.WaitGroup,user string){
	defer wg.Done()
	for ;total.value>0;{
		total.Lock()
		total.value= total.value-1
		total.Unlock()
		fmt.Println(total.value,user)
	}
}
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	total.value = 100
	go worker(&wg,"user1")
	go worker(&wg,"user2")
	wg.Wait()
	//fmt.Println(total.value)
}
