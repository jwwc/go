/*type RWMutex struct {
    w           Mutex  // held if there are pending writers
    writerSem   uint32 // semaphore for writers to wait for completing readers
    readerSem   uint32 // semaphore for readers to wait for completing writers
    readerCount int32  // number of pending readers
    readerWait  int32  // number of departing readers
}

func (*RWMutex) Lock    // 写锁

func (*RWMutex) Unlock

func (*RWMutex) RLock   // 读锁

func (*RWMutex) RUnlock*/
package main
import(
	"fmt"
	"sync"
	"time"
)
var m* sync.RWMutex
var val = 0
func main(){
	m = new(sync.RWMutex)
	go read(1)
	go write(2)
	go read(3)
	time.Sleep(5*time.Second)
}
func read(i int){
	fmt.Println(i,"begin read")
	m.RLock()
	time.Sleep(1*time.Second)
	fmt.Println(i,"val:",val)
	time.Sleep(1*time.Second)
	m.RUnlock()
	fmt.Println(i,"end read")
}
func write(i int){
	fmt.Println(i,"begin write")
	m.Lock()
	val = 10
	fmt.Println(i,"val:",val)
	time.Sleep(1*time.Second)
	m.Unlock()
	fmt.Println(i,"end write")
}
