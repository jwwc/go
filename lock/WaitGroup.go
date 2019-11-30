/*type WaitGroup struct{
	contains filtered or unexported fields
}
func (wg *WaitGroup)Add(delta int)

func(wg *WaitGroup)Done()

func(wg *WaitGroup)Wait()*/

package main

import (
	"fmt"
	"sync"
	"time"
)
func main(){
	var wg sync.WaitGroup
	var str = []string{
		"hello,world",
		"hello,Go",
		"Bye,PHP",
	}
	wg.Add(len(str))
	for _,s:=range str{
//		wg.Add(1)
		go func(s string){
			defer wg.Done()
			read(s)
		}(s)

	}
	wg.Wait()
}
func read(s string){
	time.Sleep(time.Second*1)
	fmt.Println(s)
}
