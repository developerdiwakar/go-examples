package main

import (
	"container/list"
	"fmt"
)

func main() {
	var inList list.List

	inList.PushBack(12)
	inList.PushBack(23)
	inList.PushBack(43)
	inList.PushBack(56)

	for ele := inList.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value.(int))
	}
}
