package tree

import (
	"container/list"
	"fmt"
)

//需要有visited节点来保存已经遍历过的，在pushback的时候加入没有遍历过的
//这适合图的广度优先遍历
func BFS(root*TreeNode){
	lst:=list.New()
	lst.PushBack(root)
	for root!=nil&&lst.Len()!=0{
		elem:=lst.Front()
		lst.Remove(elem)
		node:=elem.Value.(*TreeNode)
		fmt.Printf("%d\t",node.Val)
		if node.Left!=nil{
			lst.PushBack(node.Left)
		}
		if node.Right!=nil{
			lst.PushBack(node.Right)
		}
	}
	fmt.Println()
}

//深度优先遍历，类似于树的先序遍历
func DFS(root *TreeNode){
	if root!=nil{
		fmt.Printf("%d\t",root.Val)
		DFS(root.Left)
		DFS(root.Right)
	}
}

//深度优先遍历，使用栈来实现
func DFS2(root *TreeNode){
	stack:=list.New()
	stack.PushBack(root)
	for root!=nil && stack.Len()!=0{
		elem:=stack.Back()
		stack.Remove(elem)
		node:=elem.Value.(*TreeNode)
		fmt.Printf("%d\t",node.Val)
		if node.Right!=nil{
			stack.PushBack(node.Right)
		}
		if node.Left!=nil{
			stack.PushBack(node.Left)
		}
	}
	fmt.Println()
}

//层序遍历输出到二维数组，每个数组是树的每一层元素
func DFS3(root *TreeNode)[][]int{
	var res [][]int
	dfs(0,root,&res)
	fmt.Println(res)
	return res
}

func dfs(level int,node *TreeNode,res *[][]int){
	if node!=nil{
		if level>=len(*res){
			*res=append(*res,[]int{})
		}
		(*res)[level]=append((*res)[level],node.Val)
		dfs(level+1,node.Left,res)
		dfs(level+1,node.Right,res)
	}
}
