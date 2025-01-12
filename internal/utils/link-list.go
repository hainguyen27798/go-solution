package utils

import "fmt"

// Node represents a single element in a linked list, holding data and a reference to the next node.
type Node struct {
	Data interface{}
	Next *Node
}

// LinkList represents a singly linked list structure with a reference to the head node.
type LinkList struct {
	head *Node
}

// NewLinkList creates and returns a new instance of an empty LinkList.
func NewLinkList() *LinkList {
	return &LinkList{}
}

// Insert adds a new node containing the given data to the beginning of the linked list.
func (ll *LinkList) Insert(data interface{}) {
	newNode := &Node{data, nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.Next = ll.head
		ll.head = newNode
	}
}

// InsertAtTail adds a new node containing the given data at the end of the linked list.
func (ll *LinkList) InsertAtTail(data interface{}) {
	newNode := &Node{Data: data, Next: nil}
	if ll.head == nil {
		// If the list is empty, the new node becomes the head
		ll.head = newNode
		return
	}

	// Traverse to the end of the list
	currentNode := ll.head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	// Attach the new node at the end
	currentNode.Next = newNode
}

// Delete removes the first occurrence of a node containing the specified data from the linked list.
// If the list is empty or the data is not found, it performs no operation.
func (ll *LinkList) Delete(data interface{}) {
	if ll.head == nil {
		return
	}
	if ll.head.Data == data {
		ll.head = ll.head.Next
		return
	}
	currentNode := ll.head
	for currentNode.Next != nil && currentNode.Next.Data != data {
		currentNode = currentNode.Next
	}

	if currentNode.Next.Data == data {
		currentNode.Next = currentNode.Next.Next
	}
}

// Display prints the data of all nodes in the linked list in order, separated by "->", ending with "nil".
func (ll *LinkList) Display() {
	currentNode := ll.head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}

// Size returns the number of nodes currently present in the linked list.
func (ll *LinkList) Size() int {
	size := 0
	currentNode := ll.head
	for currentNode != nil {
		size++
		currentNode = currentNode.Next
	}
	return size
}

// InsertAt inserts a new node with the given data at the specified index in the linked list.
// If the index is out of bounds, it appends the node to the end of the list.
func (ll *LinkList) InsertAt(index int, data interface{}) {
	if index < 0 {
		return // Negative index is not valid
	}

	newNode := &Node{Data: data, Next: nil}

	if index == 0 {
		// Insert at the head
		newNode.Next = ll.head
		ll.head = newNode
		return
	}

	currentIndex := 0
	currentNode := ll.head

	// Traverse the list to find the correct position or the end
	for currentNode != nil && currentIndex < index-1 {
		currentNode = currentNode.Next
		currentIndex++
	}

	if currentNode == nil {
		// If index is out of bounds, do nothing
		return
	}

	// Insert the new node at the specified index
	newNode.Next = currentNode.Next
	currentNode.Next = newNode
}

// Reverse reverses the order of the nodes in the linked list.
func (ll *LinkList) Reverse() {
	if ll.head == nil || ll.head.Next == nil {
		return // No need to reverse if the list is empty or has only one node
	}

	var prevNode, currentNode, nextNode *Node
	currentNode = ll.head

	for currentNode != nil {
		nextNode = currentNode.Next // Store the next node
		currentNode.Next = prevNode // Reverse the current node's Next pointer
		prevNode = currentNode      // Move prevNode to this node
		currentNode = nextNode      // Move to the next node
	}

	ll.head = prevNode // Update the head to the new first node
}
