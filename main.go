package main

import (
	"container/heap"
	"learnaction/backtrack"
	"learnaction/cache"
	"learnaction/hash"
	"learnaction/mysort"
	"learnaction/tree"
	"log"
	"net"
	//"net/http"
	//"runtime"

	//"math"

	///"container/heap"
	"container/list"
	"fmt"
	"learnaction/stackqueue"
	"sync"
	"time"
	"unsafe"
)

func align() {
	type Person struct {
		//flag bool
		_       struct{} `order:"dislike,like,age"`
		dislike uint8
		age     int16
		like    uint8
	}

	type PPerson struct {
		dislike uint8
		like    uint8
		age     int16
		P       Person
	}
	p := Person{}

	fmt.Println(unsafe.Sizeof(p), unsafe.Alignof(p), unsafe.Sizeof(PPerson{}), unsafe.Alignof(PPerson{}))
	var a interface{} = p
	fmt.Println(unsafe.Sizeof("12345"), unsafe.Sizeof([]int{1, 2, 3, 4, 5}), unsafe.Sizeof(a), unsafe.Sizeof(main), unsafe.Sizeof(map[string]int{}))
	fmt.Println(unsafe.Alignof("12345"), unsafe.Alignof([]int{1, 2, 3, 4}), unsafe.Alignof(false), unsafe.Alignof(1))
	fmt.Println(unsafe.Alignof(a), unsafe.Alignof(p))
	fmt.Println(unsafe.Offsetof(p.dislike), unsafe.Offsetof(p.age), unsafe.Offsetof(p.like))
	fmt.Println(unsafe.Alignof(p.dislike), unsafe.Alignof(p.age), unsafe.Alignof(p.like))

	type Part1 struct {
		a bool
		b int32
		c int8
		d int64
		e byte
	}

	part1 := Part1{}

	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	var aa bool
	var bb int16
	var cc int64
	fmt.Printf("%p,%p,%p", &aa, &bb, &cc)
}

func point(){
	type Person struct{
		flag bool
		age int
		name string
	}
	p :=Person{true,19,"adfg"}
	pp:=unsafe.Pointer(&p)
	pptr:=pp
	//Pointer类似于void*
	fmt.Println(pp,pptr)
	pflag:=uintptr(pp)+unsafe.Offsetof(p.flag)
	page:=uintptr(pp)+unsafe.Offsetof(p.age)
	pname:=uintptr(pp)+unsafe.Offsetof(p.name)

	fmt.Printf("%x,%x,%x,%x,%x,%x,%x,%x,%x\n",pflag,&p.flag,unsafe.Offsetof(p.flag),page,&p.age,unsafe.Offsetof(p.age),pname,&p.name,unsafe.Offsetof(p.name))
	fmt.Println(unsafe.Pointer(page),*(*int)(unsafe.Pointer(page)))
	var i uint=1022
	fmt.Printf("%x,%p\n",i,&i)
	ip:=unsafe.Pointer(&i)
	i8:=(*uint8)(ip)
	fmt.Println(ip,i8,*i8)
	i82:=(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&i))+1))
	fmt.Println(i82,*i82)
	*i82=22
	fmt.Printf("%d,%x,%p\n",i,i,&i)
	pt:=uintptr(unsafe.Pointer(new(int)))
	ptp:=(*int)(unsafe.Pointer(pt))
	*ptp=10
	fmt.Println(pt,ptp,*ptp)

}

func chan1(){
	c:=make(chan int,2)
	//chan 可以看作 *struct chan{}，是一个指针
	fmt.Println(cap(c),unsafe.Sizeof(c))
	fmt.Println(c,unsafe.Pointer(&c))
	cptr:=uintptr(unsafe.Pointer(&c))
	fmt.Printf("%x,%x,%x\n",cptr,(*int)(unsafe.Pointer(cptr)),*(*uintptr)(unsafe.Pointer(cptr)))
	ccptr:=*(*uintptr)(unsafe.Pointer(cptr))
	qcount:=*(*uint)(unsafe.Pointer(ccptr))
	dataqsiz:=*(*uint)(unsafe.Pointer(ccptr+8))
	fmt.Println(qcount,dataqsiz)

	c2:=make(chan int,1)
	c2<-10
}

func lock2(){
	lock:=sync.Mutex{}
	var i=10
	go func(lock2 *sync.Mutex){
		lock2.Lock()
		//time.Sleep(4*time.Second)
		fmt.Println("lock")
		lock2.Unlock()
	}(&lock)
	//time.Sleep(1*time.Second)
	lock.Lock()

	time.Sleep(10*time.Second)
	//	lock.Lock()
	i++
	fmt.Println(i)
	lock.Unlock()
	time.Sleep(4*time.Second)
}

type LinkList struct{
	Elem int
	Next *LinkList
}
//1->2->3->4->5  不改变原来链表，生成新的链表，空间复杂度o(n)
func reverseLinklist(head *LinkList)*LinkList{
	var res *LinkList
	for head!=nil{
		res=&LinkList{head.Elem,res}
		head=head.Next
	}
	return res
}

//空间复杂度o(1),改变原链表每个节点的指向
func reverseLinklist2(head *LinkList)*LinkList{
	var res *LinkList
	for head!=nil{
		next:=head.Next
		head.Next=res
		res=head
		head=next
	}
	return res
}
//交换链表相邻的元素，1->2->3->4  2->1->4->3
//1->2->3->4
//2-1-3-4
func swapPairs(head *LinkList)*LinkList{
	var res *LinkList=head
	if head!=nil&&head.Next!=nil{
		res=head.Next
	}
	var pre *LinkList
	for head!=nil && head.Next!=nil{
		next:=head.Next.Next
		//pre:=head
		if pre!=nil{
			pre.Next=head.Next
		}
		pre=head
		head.Next.Next=head
		head.Next=next
		head=next
	}
	return res
}

//链表是否存在环，快指针和慢置针相遇，龟兔赛跑
//在相遇点之后，一个从头开始，一个从相遇点开始，会在入口处再次相遇（找入口）
//1-2-3-nil
func hasCycle(head *LinkList)bool{
	var slow *LinkList=head
	var fast *LinkList=head
	for slow!=nil && fast!=nil && fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			return true
		}
	}
	return false
}

//链表每k个元素进行翻转，k长度小于等于链表长度
//1-2-3-4-5-6-7  k=3
//3-2-1-6-5-4-7
func reverseKGroup(head *LinkList,k int)*LinkList{
	var pre,res ,slab *LinkList

	for j:=0;head!=nil;j++ {
		slab =nil //每一段长度为k的翻转链表
		last := head         //上一段的最后一个节点
		for i := 0; i < k; i++ {
			if head==nil{//剩余节点不翻转,逆操作
				var slab2 *LinkList
				for slab!=nil{
					next2:=slab.Next
					slab.Next=slab2
					slab2=slab
					slab=next2
				}
				slab=slab2
				break
			}
			next:=head.Next
			head.Next = slab
			slab = head
			head = next
		}
		if pre!=nil {
			pre.Next = slab
		}
		pre=last
		if j == 0 {
			res = slab
		}
		//iterLinklist(slab)
	}
	return res
}

func iterLinklist(head *LinkList){
	for head!=nil{
		fmt.Printf("%d\t",head.Elem)
		head=head.Next
	}
	fmt.Println()
}

func listT(){
	lst:=&LinkList{1,&LinkList{2,&LinkList{3,&LinkList{4,&LinkList{5,nil}}}}}
	iterLinklist(lst)
	res:=reverseLinklist(lst)
	iterLinklist(res)
	res2:=reverseLinklist2(lst)
	iterLinklist(res2)
	iterLinklist(lst)
	res3:=swapPairs(res2)
	iterLinklist(res3)

	lst2:=&LinkList{1,&LinkList{2,&LinkList{3,&LinkList{4,&LinkList{5,&LinkList{6,nil}}}}}}
	res4:=reverseKGroup(lst2,4)
	iterLinklist(res4)
}

//判断是否是有效的括号(([]{}))
func isValidBrackets(s string)bool{
	bracketMap:=map[byte]byte{')':'(',']':'[','}':'{'}
	lst:=list.New()
	for i:=0;i<len(s);i++{
		if val,ok:=bracketMap[s[i]];!ok{
			lst.PushBack(s[i])
		}else{
			elem:=lst.Back()
			if lst.Len()==0||elem.Value.(byte)!=val{
				return false
			}
			lst.Remove(elem)
		}
	}
	return lst.Len()==0
}

func runef(){
	a:="我是中国china"
	//for range 遍历的是rune类型
	for i,s:=range a{
		fmt.Printf("%c,%d\t",s,i)
	}
	fmt.Println(len(a)) //底层字节数组的长度,字符串常量符合utf-8格式
	for i:=0;i<len(a);i++{
		fmt.Printf("%c,%d",a[i],i)
	}
	fmt.Println(string([]byte(a)))
	b:=[]byte(a)
	b[0]=2
	for i:=0;i<len(b);i++{
		fmt.Printf("%c,%d,%T",b[i],i,b[i])
	}
	fmt.Println(string(b))

	r:=[]rune(a)
	fmt.Println(r)
	fmt.Println(string(r))
}

func queue1(){
	s:="(([]{})"
	fmt.Println(isValidBrackets(s))

	queue := stackqueue.Constructor2();

	queue.Push(1);
	queue.Push(2);
	fmt.Println(queue.Peek(),queue.Pop(),queue.Empty());  // 返回 1
	queue.Push(3)
	queue.Push(4)
	fmt.Println(queue.Peek(),queue.Pop(),queue.Empty());  // 返回 1
}

func kthlagest(){

	h := &stackqueue.IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h,1)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	k:= 3;
	arr := []int{4,5};
	kthLargest := stackqueue.Constructor(k, arr);
	kthLargest.Add(3)
	kthLargest.Add(5)
	kthLargest.Add(10)
	kthLargest.Add(9)
	kthLargest.Add(4)
	//heap.Remove(h,0)
	//fmt.Println([]int{1,2,3,4}[1:4])
	// Output:
	// minimum: 1
	// 1 2 3 5

}

func twosum(){
	//a:=map[int]int{}
	//b:=map[int]int{}
	//fmt.Println(a==b)
	input:=[]int{-1, 0, 1, 2, -1, -4}
	input=[]int{-4,-1,-1,0,1,2,2,3,4,5}
	fmt.Println(hash.ThreeSum2(input,0))
	fmt.Println(hash.FourSum([]int{1, 0, -1, 0, -2, 2},0))
}

func tree1(){
	root:=tree.TreeNode{2,&tree.TreeNode{1,&tree.TreeNode{4,nil,nil},&tree.TreeNode{5,nil,nil}},&tree.TreeNode{3,nil,nil}}
	//tree.RecursiveInOrderReverse(&root)
	//fmt.Println(tree.IsValidBST(&root))
	tree.LowestCommonAncestor(&root,&root,&root)
	l1:=&stackqueue.ListNode{1,&stackqueue.ListNode{2,&stackqueue.ListNode{3,
		&stackqueue.ListNode{4,&stackqueue.ListNode{5,&stackqueue.ListNode{6,nil}}}}}}
	l2:=&stackqueue.ListNode{4,&stackqueue.ListNode{5,&stackqueue.ListNode{6,nil}}}
	res:=stackqueue.GetIntersectionNode(l1,l2)
	fmt.Println(res)
}

func reverse(){
	//divideconquer.MyPow(2,4)
	/*
					1
			2				5
		3		4		6		7
	*/
	left:= &tree.TreeNode{2,&tree.TreeNode{3,nil,nil},&tree.TreeNode{4,nil,nil}}
	right:= &tree.TreeNode{5,&tree.TreeNode{6,nil,nil},&tree.TreeNode{7,nil,nil}}

	root:=&tree.TreeNode{1,left,right}

	tree.BFS(root)
	tree.DFS(root)
	fmt.Println()
	tree.DFS2(root)
	tree.DFS3(root)
}

func lfru(){
	cache2:= cache.Constructor(1)
	cache2.Put(2,1)
	fmt.Println(cache2.Get(2))
	cache2.Put(3,2)
	fmt.Println(cache2.Get(2))
	fmt.Println(cache2.Get(3))

}

func sort(){
	var arr=[]int{4,2,6,5,10,9}
	mysort.BubbleSort(arr)
	fmt.Println(arr)

	arr=[]int{6,5,4,3,2}
	mysort.BubbleSort(arr)
	fmt.Println(arr)

	arr=[]int{6,5,4,3,2}
	mysort.InsertSort(arr)
	fmt.Println(arr)
}

func net1(){
	//http.ListenAndServe(":8989",nil)
	listener,err:=net.Listen("tcp",":8989")
	if err!=nil{
		fmt.Println(err)
	}
	//go func(){
	//	conn,_:=net.Dial("tcp",":8989")
	//	//defer conn.Close()
	//	time.Sleep(5*time.Second)
	//	conn.Write([]byte("123"))
	//
	//}()
	for{
		conn,err:=listener.Accept()
		fmt.Println(conn.LocalAddr(),conn.RemoteAddr())
		if err!=nil{
			fmt.Println(err)
		}
		go func(c net.Conn){
			defer c.Close()
			var buf []byte=make([]byte,10)
			//for {
			n, err := c.Read(buf)
			if err != nil {
				fmt.Println(err)
				//return
			}
			fmt.Println(n, buf[:n], string(buf[:n]))
			//}
			total:=0

			for{
				data:=make([]byte,1024*1024*10)
				n,err:=c.Write(data)
				if err!=nil{
					fmt.Println(err)
				}
				total += n
				log.Printf("server %s write %d bytes this time, %d bytes in total\n",buf, n, total)
			}
		}(conn)
	}
	time.Sleep(1000*time.Second)
}

func str3(){
	//strings.MyAtoi("123")
	//strings.MyAtoi("a123")
	//strings.yAtoi("-123a")
	//strings.MyAtoi("91283472332")
	//strings.MyAtoi("2147483648")
	//strings.MyAtoi("-6147483648")
	//fmt.Println("123","123"[:0])
	//fmt.Println(strings2.Index("123","12"))
	//fmt.Println(strings.LongestCommonPrefix([]string{"123","453"}))
	//fmt.Println(strings.Bm("abcacabdc","abd"))
	//arr:=[]int{1,9,8,3,4,10,7}
	//fmt.Println(arr)
	//mysort.MergeSort(arr,0,len(arr)-1)
	//fmt.Println(arr)
}

func lock3(){
	lock:=sync.Mutex{}
	lock.Lock()
	lock.Lock()
	fmt.Println("main")
	lock.Unlock()

	go func(){
		lock.Unlock()
		fmt.Println("lock")
		//lock.Lock()
	}()
	time.Sleep(1*time.Second)
}

func select1(){
	ch:=make(chan string)
	go func(){
		time.Sleep(2*time.Second)
		//ch<-"10"
		//close(ch)
	}()
	select {
	case <-ch:
		fmt.Println("chan done")
	case <-time.After(2*time.Second):
		fmt.Println("timeout")
		//default:
		fmt.Println("default")

	}
	select {

	}
}

func t(j int) (i int){
	defer func(){
		i++
	}()
	return 1
}

func main(){
	//lock3()
	//arr:=[]int{7,3,10,6,8,1,5}
	//fmt.Println(mysort.KthLargest(arr,0,len(arr)-1,1),arr)
	//mysort.QuickSort(arr,0,len(arr)-1)
	//fmt.Println(arr)

	//fmt.Println(arr,0,len(arr)-1,3)

	//backtrack.Call8Queens(1)
	//fmt.Println(backtrack.SubSet2([]int{}))
	//fmt.Println("abcd"[4:4])
	//fmt.Println(backtrack.LetterCombinations("2"))

	//backtrack.Permute([]int{1,2,3})
	//myreflect.Reflect1()

	//backtrack.StrSegment("0123456")
	//fmt.Println(backtrack.Combine(5,3))
	//fmt.Println(runtime.NumCPU(),runtime.GOARCH)
	//fmt.Println(t(10))

	//fmt.Println("done" )
	//2,4,6,8
	//8,2,4,6
	//6,8,2,4
	//for i:=1;i<10;i++{
	//	//fmt.Println(i,bianrysearch.Search([]int{1,2,2,2,3},i))
	//	//bianrysearch.Search([]int{1,2,3,4,5,6},i)
	//	//bianrysearch.Search2([]int{1,2,2,3},i)
	//	fmt.Println(bianrysearch.Search2([]int{1,2},i))
	//}
	//bianrysearch.Search2([]int{1,2,2,2,2,3,4,5,6,7},2)
	//
	//matrix:=[][]int{[]int{1,3,5,7},[]int{10,11,16,20},[]int{23,30,34,50}}
	//fmt.Println(bianrysearch.SearchMatrix(matrix,20))
	//fmt.Println(bianrysearch.SearchReverseArr([]int{3,3,1,2,2,3,3,3,3}))
	//fmt.Println(bianrysearch.SearchReverseArr2([]int{3,3,3,3,3,3,4,5,3,3,3,3},4))
	//fmt.Println(bianrysearch.SearchReverseArr3([]int{3,3,3,3,3,3,4,5,3,3,3,3},4))
	//3!/2=3，4！/3！=4
	backtrack.MapConbine(map[string]int{"A":1,"B":2})//3  3！/2！
	backtrack.MapConbine(map[string]int{"A":1,"B":3})//4  4！/3！
	backtrack.MapConbine(map[string]int{"A":1,"B":4})//5  5！/4！

	//sum(m1,m2...)!/m1!*m2!*...mn!
	backtrack.MapConbine(map[string]int{"A":2,"B":2})//6   4！/2！*2！
	backtrack.MapConbine(map[string]int{"A":2,"B":3})//10  5！/3！*2！
	backtrack.MapConbine(map[string]int{"A":2,"B":4})//15  6！/4！*2！


}


