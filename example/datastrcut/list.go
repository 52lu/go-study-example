package datastrcut

import (
	"container/list"
	"fmt"
)

func CreateList()  {
	l := list.New()
	nodeA := l.PushFront("A")
	nodeB := l.InsertAfter("B", nodeA)
	nodeC := l.InsertAfter("C", nodeB)
	nodeD := l.InsertAfter("D", nodeC)
	l.InsertAfter("B",nodeD)
	//l.InsertAfter("C",nodeB)
	//l.InsertAfter("D",nodeB)
	for font := l.Front(); font != nil; {
		
	}
	fmt.Println(l)

}
