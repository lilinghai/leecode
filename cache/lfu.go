package cache

type FNode struct {
	Key int
	Value int
	Frequency int
	Pre *FNode
	Next *FNode
}

type LFUCache struct {
	Root *FNode
	Cache map[int]*FNode
	Capacity int
	Size int
}


func Constructor2(capacity int) LFUCache {
	root:=&FNode{}
	root.Pre=root
	root.Next=root
	return LFUCache{
		Root:root,
		Cache:make(map[int]*FNode),
		Capacity:capacity,
		Size:0,
	}
}

//如果不在cache直接返回，在cache中，修改频次，移动链表节点
//移动元素的时候可能会出现级联的情况，如root->2(5)->4(5)->6(5)->7(4)->8(4)->1(1),查找节点7，会引起7交换到头部
func (this *LFUCache) Get(key int) int {
	if _,ok:=this.Cache[key];!ok{
		return -1
	}
	node:=this.Cache[key]
	node.Frequency++

	//remove节点node
	remove2(node)

	//比较frequency，移动node。根节点频率是0，移动到大于node的频率为止,npre则是要插入的位置
	compareMove(node)
	return node.Value
}

func remove2(node *FNode){
	//remove节点node
	npre:=node.Pre
	nnext:=node.Next
	npre.Next=nnext
	nnext.Pre=npre
}

func compareMove(node *FNode){
	//比较frequency，移动node。根节点频率是0，移动到大于node的频率为止,npre则是要插入的位置
	npre:=node.Pre
	for npre.Frequency!=0 && npre.Frequency<=node.Frequency{
		npre=npre.Pre
	}
	nprenext:=npre.Next
	node.Pre=npre
	node.Next=nprenext
	npre.Next=node
	nprenext.Pre=node
}

//如果在cache中，更新频次，判断是否移动；如果不在cache中，判断是否超过capacity，没有超过，在尾部插入新的；超过，淘汰尾部,在尾部插入新的
func (this *LFUCache) Put(key int, value int)  {
	if this.Capacity<1{
		return
	}
	if elem,ok:=this.Cache[key];ok{
		elem.Frequency++
		elem.Value=value
		remove2(elem)
		compareMove(elem)
		return
	}
	if this.Size==this.Capacity{
		delete(this.Cache,this.Root.Pre.Key)
		remove2(this.Root.Pre)
		node:=&FNode{Key:key,Value:value,Frequency:1,Pre:this.Root.Pre}
		compareMove(node)
		this.Cache[key]=node
	}else{
		node:=&FNode{Key:key,Value:value,Frequency:1,Pre:this.Root.Pre}
		compareMove(node)
		this.Cache[key]=node
		this.Size++
	}
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
