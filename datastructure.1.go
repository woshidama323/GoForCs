package main

import (
	"fmt"
	"errors"
)

//Queue is a queue for practicing
type Queue struct{

	//item store
	Data []string

	Cap uint64
	Head  uint64
	Rear   uint64
	CurrentLen uint64
}

func New(cap int64) *Queue{

	return &Queue{
		Data:make([]string,0,cap),
		Head :0,
		Rear :0,
		Cap: uint64(cap),
		CurrentLen: 0,
	}
}

func ( Que *Queue) EnQueueHead(input string) error{


	
	if Que.CurrentLen == Que.Cap{
		fmt.Println("have get the maximun of the capacity in this queue")
		return errors.New("the queue is expire")
	} 


	Que.Data = append(Que.Data,input)
	if Que.Head == 0 {
		Que.Head = 100
	}
	Que.Head = Que.Head - 1
	
	// Que.CurrentLen = uint64(len(Que.Data))
	Que.CurrentLen ++ 

	return nil
}

func ( Que *Queue) EnQueueRear(input string) error{

	if Que.CurrentLen == Que.Cap{
		fmt.Println("have get the maximun of the capacity in this queue")
		return errors.New("the queue is expire")
	} 


	Que.Data = append(Que.Data,input)
	Que.Rear = (Que.Rear + 1)%Que.Cap
	// Que.CurrentLen = uint64(len(Que.Data))
	Que.CurrentLen ++ 

	return nil
}



func ( Que *Queue) DeQueueHead() (string,error){

	fmt.Println("starting dequeue........")
	if Que.CurrentLen == 0{
		
		return "",errors.New("it's empty now")
	}

	//先取出来值，然后再进行指针的变化
	getitem := Que.Data[Que.Head]
    Que.Head = (Que.Head+1)%Que.Cap
	Que.CurrentLen --
 
	fmt.Println("get the queue string from a queue",getitem)
	return getitem,nil

}

func ( Que *Queue) DeQueueRear() (string,error){

	fmt.Println("starting dequeue........")
	if Que.CurrentLen == 0{
		
		return "",errors.New("it's empty now")
	}

	//先取出来值，然后再进行指针的变化
	
	if Que.Rear == 0{
		Que.Rear = Que.Cap
	}
	Que.Rear --
	getitem := Que.Data[Que.Rear]

	Que.CurrentLen --
 
	fmt.Println("get the queue string from a queue",getitem)
	return getitem,nil

}

// func ( Que *Queue) GetLenth(){
// 	return len(Que.Field1)
// }

func testqueue(){
	why := New(100)

	for i := 0;i<99;i++{
		item := fmt.Sprintf("what_%d",i)
		fmt.Println(".......................Enqueue..",item)
		why.EnQueueRear(item)
	}
	why.EnQueueRear("expire")
	for {
		item,err := why.DeQueueRear()
		if err !=nil {
			fmt.Println("what happened...",err)
			break
		}
		fmt.Println(".....the item is ...",item)
	}

	fmt.Println("dequeue....",why.Data)
	fmt.Println("hello there")
}
