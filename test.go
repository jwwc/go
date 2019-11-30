package main
import(
	"fmt"
//	"time"
)
func main() {
    done := make(chan int) // 带缓存通道

    go func(){
        fmt.Println("你好, 世界")
        <-done 
    }()

    done<-1
}

