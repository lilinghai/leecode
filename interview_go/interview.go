package interview_go

import (
"fmt"
	"runtime"
	"sync"
)

/*
打印后
打印中
打印前
panic: 触发异常
1.在退出函数前defer会按照LIFO顺序打印
2.defer的参数在defer声明的时候确定
3.

 */
func Defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

/*
panic call2
panic_call
panic: call2 panic

panic 会一直向上传递，F调用者（caller）调用panic的时候，main调用F函数的时候，就像直接调用panic一样
参考：https://blog.golang.org/defer-panic-and-recover
 */
func Panic_call(){
	defer fmt.Println("panic_call")
	panic_call2()
}

func panic_call2(){
	defer fmt.Println("panic call2")
	panic("call2 panic")
}

/*
recover能够捕获panic，recover只有在defer里面还有用，因为调用者panic之后，
panic后面的代码不会执行，必须在函数退出的时机内来调用recover
 */
func Recover_call() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	g(0)
	fmt.Println("Calling g.")
	fmt.Println("hello")
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/*
输出key不一样，value一样
range 的时候会声明一个临时变量，然后用slice的每个值给临时变量赋值
 */
type student struct {
	Name string
	Age  int
}

func RangeCall() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return m
}

/*
gomaxporcs设置 P的个数，如果是1，则只能在一个P上调度，同时只能绑定一个cpu
第一个循环输出都是10，第二个循环输出i,从0到9，无序
 */
func ProcsCall(){
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("k: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
showA
showB
teacher showB

 */
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func CombineCall(){
	t:=Teacher{}
	t.ShowA()
	t.ShowB()
}




