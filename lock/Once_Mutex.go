/*type Once struct{
	m Mutex
	done uint32
}
func(o *Once)Do(f func())*/
package main
import(
	"fmt"
	"sync"
	"time"
)
func main(){
	var once = new(sync.Once)//var once sync.Once
	for i:=0;i<10;i++{
		go func(){
			once.Do(read)
		}()
	}
	time.Sleep(time.Second)
}
func read(){
	fmt.Println(1)
}
