package cache

type Node struct {
	Key int
	Value int
	Pre *Node
	Next *Node

}

type LRUCache struct {
	Root *Node
//	Tail *Node
	Capacity int
	Size int
	Cache map[int]*Node
}


func Constructor(capacity int) LRUCache {
	head:=&Node{}
	head.Next=head
	head.Pre=head
	return LRUCache{
		Root:head,
		Capacity:capacity,
		Size:0,
		Cache:make(map[int]*Node),
	}
}

//1.不在cache中；2.在cache中，需要移动到头节点
func (this *LRUCache) Get(key int) int {
	if _,ok:=this.Cache[key];!ok{
		return -1
	}
	node:=this.Cache[key]
	//从链表remove该节点
	remove(node)
	//移动到头节点
	pushFront(this.Root,node)
	return node.Value
}

func remove(node *Node){
	npre:=node.Pre
	nnext:=node.Next
	npre.Next=nnext
	nnext.Pre=npre
}

func pushFront(root *Node,node *Node){
	head:=root.Next
	root.Next=node
	node.Pre=root
	node.Next=head
	head.Pre=node
}

//如果key在cache中，旧数据删除，新数据插入头节点；不在cache中，如果长度达到上限，删除尾数据，新数据插入头部，如果没有，直接插入新数据到头部
func (this *LRUCache) Put(key int, value int)  {
	if node,ok:=this.Cache[key];ok{
		//复用原来的节点
		node.Value=value
		this.Cache[key]=node
		remove(node)
		pushFront(this.Root,node)
		return
	}
	if this.Size==this.Capacity{
		delete(this.Cache,this.Root.Pre.Key)
		remove(this.Root.Pre)
		node:=&Node{Key:key,Value:value,}
		pushFront(this.Root,node)
		this.Cache[key]=node
	}else{
		node:=&Node{Key:key,Value:value,}
		pushFront(this.Root,node)
		this.Size++
		this.Cache[key]=node
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
