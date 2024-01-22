package dataStructures

import (
	"errors"
	"fmt"
)

type Node[T comparable] interface {
	Next() Node[T]
	Value() T
}

type LinkedList[T comparable] interface {
	Head() Node[T]
	Size() int
	AddFirst(value T)
	Add(value T)
	RemoveAll(value T)
	RemoveFirst(value T)
	Contains(value T) bool
	ContainsAll(value ...T) bool
	Get(index int) (T, bool, error)
}

func NewLinkedList[T comparable]() LinkedList[T] {
	return &linkedList[T]{}
}

type linkedList[T comparable] struct {
	head *node[T]
}

type node[T comparable] struct {
	next  *node[T]
	value T
}

func (n *node[T]) Next() Node[T] {
	return n.next
}

func (n *node[T]) Value() T {
	return n.value
}

func (r *linkedList[T]) Head() Node[T] {
	return r.head
}

func (r *linkedList[T]) Size() int {
	if r.head == nil {
		return 0
	} else {
		size := 1
		current := r.head
		for current.next != nil {
			size++
			current = current.next
		}
		return size
	}
}

func (r *linkedList[T]) AddFirst(value T) {
	if r.head == nil {
		r.head = &node[T]{value: value}
		return
	}

	oldHead := r.head
	r.head = &node[T]{value: value}
	r.head.next = oldHead
}

// Add warning : extremely slow
func (r *linkedList[T]) Add(value T) {
	if r.head == nil {
		r.head = &node[T]{value: value}
		return
	}

	last := node[T]{value: value}
	head := r.head
	for head.next != nil {
		head = head.next
	}
	head.next = &last
}

func (r *linkedList[T]) RemoveAll(value T) {
	prev := (*node[T])(nil)
	current := r.head
	for current != nil {
		if current.value == value {
			if prev != nil {
				prev.next = current.next
			} else {
				r.head = current.next
				current = r.head
			}
		}
		prev = current
		current = current.next
	}
}

func (r *linkedList[T]) RemoveFirst(value T) {
	prev := (*node[T])(nil)
	current := r.head
	for current != nil {
		if current.value == value {
			if prev != nil {
				prev.next = current.next
			} else {
				r.head = current.next
				current = r.head
			}
			return
		}
		prev = current
		current = current.next
	}
}

func (r *linkedList[T]) Contains(value T) bool {
	if r.Size() == 0 {
		return false
	}
	current := r.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (r *linkedList[T]) Get(index int) (T, bool, error) {
	size := r.Size()
	if index > size-1 && size > 0 {
		return *new(T), false, errors.New(fmt.Sprintf("Index %v out of bounds", index))
	}

	current := r.head
	currentIndex := 0
	for currentIndex < index {
		current = current.next
		currentIndex++
	}
	return current.value, true, nil
}

func (r *linkedList[T]) ContainsAll(value ...T) bool {
	if r.Size() == 0 {
		return false
	}

	var lookup = make(map[T]int)
	for index, element := range value {
		lookup[element] = index
	}

	var current = r.head
	for current != nil {
		if lookup[current.value] >= 0 {
			delete(lookup, current.value)
		}
		current = current.next
	}
	return len(lookup) == 0
}
