package main

import (
	"fmt"
)

// Creating a Node
type Node struct {
	Message string
	Next *Node
}

// Creating a Queue
type Queue struct {
	Head *Node // pointer to front of Queue
	Tail *Node // pointer to the back of the queue
	Size int   // automatically defualts to 0
}

// How to create methods for a struct
// func (s Struct) methodName(variable variableType) returnType { /* code /*}

//// Functions for Queueing
// Enqueue -- add new message to the back of the queue -- O(1)
func (q *Queue) Enqueue(node *Node) {
	node.Next = q.Tail
	q.Tail = node
	q.Size += 1
}

// Dequeue -- remove message from the front of the queue -- O(n)
func (q *Queue) Dequeue() *Node {

	if (q.IsEmpty()) {
		return &Node{"Queue is empty", nil}
	}

	var node = q.Tail
	var nextUp = Node{}
    for {

		if node.Next == nil { // empty queue
			nextUp = *node
			nextUp.Next = nil
			q.Tail = nil
			q.Head = nil
			q.Size -= 1
			break
		}

		if node.Next == q.Head {
			nextUp = *q.Head
			node.Next = nil
			q.Head = node
			q.Size += 1
			break
		}
		node = node.Next
	}
	return &nextUp
}

// Length -- returns length of queue
func (q Queue) Length() int {
	return q.Size
}

// IsEmpty -- checks if queue is empty
func (q Queue) IsEmpty() bool {
	if (q.Size == 0) {
		return true
	}
	return false
}
// Display -- returns current list of messages in queue

// TODO: this function did not work, wouldn't return the right head value to reset
// Traversal -- traverse the queue until we reach Head, return node before head
// func Traversal(q *Queue, head *Node) *Node {
// 	var node = q.Tail
// 	var nextUp = Node{}
//     for {
// 		// TODO: check for if not head then nil - or in case of empty queue
// 		fmt.Printf("traverse %s\n", node.Next.Message)
// 		if node.Next == head {
// 			nextUp = *node
// 			break
// 		}
// 		node = node.Next
// 	}
// 	return &nextUp
// }

func main() {
	var node1 = Node{"Hello", nil}
	var queue = Queue{&node1, &node1, 0}

	queue.Enqueue(&Node{"World", nil})
	queue.Enqueue(&Node{"Look At Me", nil})
	queue.Enqueue(&Node{"Go!", nil})

	fmt.Printf("Dequeue Hello? %s\n", queue.Dequeue().Message)
	fmt.Printf("Dequeue World? %s\n", queue.Dequeue().Message)
	fmt.Printf("Size of Queue? %d\n", queue.Length())
	fmt.Printf("Dequeue Look At Me? %s\n", queue.Dequeue().Message)
	fmt.Printf("Dequeue Go! %s\n", queue.Dequeue().Message)

}