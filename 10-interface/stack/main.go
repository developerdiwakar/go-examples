package main

import (
	"errors"
	"fmt"
)

/*
Implement a basic stack data structure that supports the following operations:
-Push(item interface{}): Add an item to the top of the stack.
-Pop() (interface{}, error): Remove and return the item from the top of the stack. Return an error if the stack is empty.
-IsEmpty() bool: Return true if the stack is empty; otherwise, return false.
*/
var ErrStackIsEmpty = errors.New("stack is empty")

type IStack interface {
	Push(item interface{})
	Pop() (interface{}, error)
	IsEmpty() bool
}

type IntergerStack struct {
	data []int
}

func (s *IntergerStack) Push(item interface{}) {
	s.data = append(s.data, item.(int))
}

func (s *IntergerStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}

	s.data = s.data[:len(s.data)-1]
	return s.data, nil
}

func (s *IntergerStack) IsEmpty() bool {
	return len(s.data) <= 0
}

func CreateNewStack(s IStack) IStack {
	return s
}

func main() {
	intStack := CreateNewStack(&IntergerStack{})

	intStack.Push(2)
	intStack.Push(23)
	intStack.Push(2234)
	stack, err := intStack.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Stack:", stack)
}
