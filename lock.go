//在并发环境下共享读写变量必须要使用锁来控制数据结构的安全，Go 语言内置了 sync 包，里面包含了我们平时需要经常使用的互斥锁对象 sync.Mutex。
//Go 语言内置的字典不是线程安全的，所以下面我们尝试使用互斥锁对象来保护字典，让它变成线程安全的字典。
//线程不安全的字典
//Go 语言内置了数据结构「竞态检查」工具来帮我们检查程序中是否存在线程不安全的代码。当我们在运行代码时，
//打开 -run 开关，程序就会在内置的通用数据结构中进行埋点检查。竞态检查工具在 Go 1.1 版本中引入，
// 该功能帮助 Go 语言「元团队」找出了 Go 语言标准库中几十个存在线程安全隐患的 bug，
//这是一个非常了不起的功能。同时这也说明了即使是猿界的神仙，写出来的代码也避免不了有 bug。下面我们来尝试一下

// package main
//
// import "fmt"
//
// func write(d map[string]int) {
//     d["fruit"] = 2
// }
//
// func read(d map[string]int) {
//     fmt.Println(d["fruit"])
// }
//
// func main() {
//     d := map[string]int{}
//     go read(d)
//     write(d)
// }

// package main
//
// import "fmt"
// import "sync"
//
// type SafeDict struct {
//     data  map[string]int
//     mutex *sync.Mutex
// }
//
// func NewSafeDict(data map[string]int) *SafeDict {
//     return &SafeDict{
//         data:  data,
//         mutex: &sync.Mutex{},
//     }
// }
//
// func (d *SafeDict) Len() int {
//     d.mutex.Lock()
//     defer d.mutex.Unlock()
//     return len(d.data)
// }
//
// func (d *SafeDict) Put(key string, value int) (int, bool) {
//     d.mutex.Lock()
//     defer d.mutex.Unlock()
//     old_value, ok := d.data[key]
//     d.data[key] = value
//     return old_value, ok
// }
//
// func (d *SafeDict) Get(key string) (int, bool) {
//     d.mutex.Lock()
//     defer d.mutex.Unlock()
//     old_value, ok := d.data[key]
//     return old_value, ok
// }
//
// func (d *SafeDict) Delete(key string) (int, bool) {
//     d.mutex.Lock()
//     defer d.mutex.Unlock()
//     old_value, ok := d.data[key]
//     if ok {
//         delete(d.data, key)
//     }
//     return old_value, ok
// }
//
// func write(d *SafeDict) {
//     d.Put("banana", 5)
// }
//
// func read(d *SafeDict) {
//     fmt.Println(d.Get("banana"))
// }
//
// func main() {
//     d := NewSafeDict(map[string]int{
//         "apple": 2,
//         "pear":  3,
//     })
//     go read(d)
//     write(d)
// }
//
// 避免锁复制
// 上面的代码中还有一个需要特别注意的地方是 sync.Mutex 是一个结构体对象，这个对象在使用的过程中要避免被复制 —— 浅拷贝。
// 复制会导致锁被「分裂」了，也就起不到保护的作用。所以在平时的使用中要尽量使用它的指针类型。
// 读者可以尝试将上面的类型换成非指针类型，然后运行一下竞态检查工具，会看到警告信息再次布满整个屏幕。
// 锁复制存在于结构体变量的赋值、函数参数传递、方法参数传递中，都需要注意。
// 使用匿名锁字段
// 在结构体章节，我们知道外部结构体可以自动继承匿名内部结构体的所有方法。如果将上面的 SafeDict 结构体进行改造，将锁字段匿名，
// 就可以稍微简化一下代码。

// package main
//
// import "fmt"
// import "sync"
//
// type SafeDict struct {
//     data  map[string]int
//     *sync.Mutex
// }
//
// func NewSafeDict(data map[string]int) *SafeDict {
//     return &SafeDict{data, &sync.Mutex{}}
// }
//
// func (d *SafeDict) Len() int {
//     d.Lock()
//     defer d.Unlock()
//     return len(d.data)
// }
//
// func (d *SafeDict) Put(key string, value int) (int, bool) {
//     d.Lock()
//     defer d.Unlock()
//     old_value, ok := d.data[key]
//     d.data[key] = value
//     return old_value, ok
// }
//
// func (d *SafeDict) Get(key string) (int, bool) {
//     d.Lock()
//     defer d.Unlock()
//     old_value, ok := d.data[key]
//     return old_value, ok
// }
//
// func (d *SafeDict) Delete(key string) (int, bool) {
//     d.Lock()
//     defer d.Unlock()
//     old_value, ok := d.data[key]
//     if ok {
//         delete(d.data, key)
//     }
//     return old_value, ok
// }
//
// func write(d *SafeDict) {
//     d.Put("banana", 5)
// }
//
// func read(d *SafeDict) {
//     fmt.Println(d.Get("banana"))
// }
//
// func main() {
//     d := NewSafeDict(map[string]int{
//         "apple": 2,
//         "pear":  3,
//     })
//     go read(d)
//     write(d)
// }
// 使用读写锁
// 日常应用中，大多数并发数据结构都是读多写少的，对于读多写少的场合，可以将互斥锁换成读写锁，
// 可以有效提升性能。sync 包也提供了读写锁对象 RWMutex，不同于互斥锁只有两个常用方法 Lock() 和 Unlock()，
// 读写锁提供了四个常用方法，分别是写加锁 Lock()、写释放锁 Unlock()、读加锁 RLock() 和读释放锁 RUnlock()。
// 写锁是排他锁，加写锁时会阻塞其它协程再加读锁和写锁，读锁是共享锁，加读锁还可以允许其它协程再加读锁，但是会阻塞加写锁。
//
// 读写锁在写并发高的情况下性能退化为普通的互斥锁。下面我们将代码中 SafeDict 的互斥锁改造成读写锁。
//
// package main
//
// import "fmt"
// import "sync"
//
// type SafeDict struct {
//     data  map[string]int
//     *sync.RWMutex
// }
//
// func NewSafeDict(data map[string]int) *SafeDict {
//     return &SafeDict{data, &sync.RWMutex{}}
// }
//
// func (d *SafeDict) Len() int {
//     d.RLock()
//     defer d.RUnlock()
//     return len(d.data)
// }
//
// func (d *SafeDict) Put(key string, value int) (int, bool) {
//     d.Lock()
//     defer d.Unlock()
//     old_value, ok := d.data[key]
//     d.data[key] = value
//     return old_value, ok
// }
//
// func (d *SafeDict) Get(key string) (int, bool) {
//     d.RLock()
//     defer d.RUnlock()
//     old_value, ok := d.data[key]
//     return old_value, ok
// }
//
// func (d *SafeDict) Delete(key string) (int, bool) {
//     d.Lock()
//     defer d.Unlock()
//     old_value, ok := d.data[key]
//     if ok {
//         delete(d.data, key)
//     }
//     return old_value, ok
// }
//
// func write(d *SafeDict) {
//     d.Put("banana", 5)
// }
//
// func read(d *SafeDict) {
//     fmt.Println(d.Get("banana"))
// }
//
// func main() {
//     d := NewSafeDict(map[string]int{
//         "apple": 2,
//         "pear":  3,
//     })
//     go read(d)
//     write(d)
// }
