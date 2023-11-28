// a simple, singly linked list implementation
package linkedlist

import (
	"fmt"
)

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Val)
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

type Iterator[T comparable] struct {
	Current *Node[T]
}

func (i *Iterator[T]) HasNext() bool {
	return i.Current != nil
}

func (i *Iterator[T]) Next() *Node[T] {
	current := i.Current
	i.Current = i.Current.Next
	return current
}

func (l *LinkedList[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{l.Head}
}

func (l *LinkedList[T]) String() string {
	var s string
	iterator := l.Iterator()
	for iterator.HasNext() {
		s += fmt.Sprintf("%v ", iterator.Next())
	}
	return s
}
