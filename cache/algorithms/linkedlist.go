package algorithms

import "github.com/mastik5h/LLD/cache/models"

type LinkedList[K models.Key] struct {
	Head *Node[K]
}

type Node[K models.Key] struct {
	Data K
	Next *Node[K]
}

func InitializeLinkedList[K models.Key]() LinkedList[K] {
	ll := LinkedList[K]{
		Head: NewNode[K](),
	}
	return ll
}

func NewNode[K models.Key]() *Node[K] {
	return &Node[K]{
		Next: nil,
	}
}

func (ll *LinkedList[K]) AddNode(key K) {
	headStore := ll.Head
	headTemp := ll.Head
	for headTemp.Next != nil {
		headTemp = headTemp.Next
	}
	headTemp.Next = NewNode[K]()
	headTemp = headTemp.Next
	headTemp.Data = key
	ll.Head = headStore
}

func (ll *LinkedList[K]) RemoveNode(key K) {
	headStore := ll.Head
	prevNode := ll.Head

	for prevNode != nil {
		currNode := prevNode.Next
		if currNode != nil && currNode.Data == key {
			prevNode.Next = currNode.Next
			break
		}
		prevNode = prevNode.Next
	}
	ll.Head = headStore
}

func (ll *LinkedList[K]) GetFirstNode() K {
	firstNode := ll.Head.Next
	return firstNode.Data
}
