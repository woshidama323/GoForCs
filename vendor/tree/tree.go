package tree

import (
	"fmt"
	"sync"
)

//Queue is a queue for practicing
type Node struct{

	//item store
	Data string

	LChild *Node
	RChile *Node
}

type Tree struct{

	//item store
	Root *Node
	Lock sync.RWMutex
}

func New() *Node{

	return &Node{
		Data : "",
		LChild: nil,
		RChile: nil,
	}
}

func ( tree *Tree) Insert(input string) {

	tree.Lock.Lock()
	// newnode := tree.New()
	// newnode.Data = "node1"
	tree.Lock.Unlock()

}

func Test(){
	// why := New(100)

	// for i := 0;i<99;i++{
	// 	item := fmt.Sprintf("what_%d",i)
	// 	fmt.Println(".......................Enqueue..",item)
	// 	why.EnQueueRear(item)
	// }
	// why.EnQueueRear("expire")
	// for {
	// 	item,err := why.DeQueueRear()
	// 	if err !=nil {
	// 		fmt.Println("what happened...",err)
	// 		break
	// 	}
	// 	fmt.Println(".....the item is ...",item)
	// }

	// fmt.Println("dequeue....",why.Data)
	fmt.Println("hello there")
}
