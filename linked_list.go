package ll

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value, Next: nil}
}

func (Node *Node) PrintNode() {
	fmt.Printf("{ Value: %d, Next: %p }\n", Node.Value, Node.Next)
}

func (Node *Node) PrintList() {
	if Node != nil {
		Node.PrintNode()
		Node.Next.PrintList()
	}
}

func (head *Node) FindNodeByValue(Value int) *Node {
	if head.Value == Value {
		return head
	}

	return head.Next.FindNodeByValue(Value)
}

func (head *Node) FindNodeByReference(target *Node) *Node {
	if head == target {
		return head
	}

	return head.Next.FindNodeByReference(target)
}

func (head *Node) DeleteNode(target *Node) *Node {
	if head.Next == target {
		head.Next = target.Next
		return target
	}

	return head.Next.DeleteNode(target)
}

func (prev *Node) AddNode(Value int) *Node {
	newNode := CreateNode(Value)
	newNode.Next = prev.Next
	prev.Next = newNode

	return newNode
}

func (head *Node) AppendNode(Value int) *Node {
	if head.Next == nil {
		Node := head.AddNode(Value)
		return Node
	}

	return head.Next.AppendNode(Value)
}

func (head *Node) ListLength(index int) int {
	if head == nil {
		return index
	}
	return head.Next.ListLength(index + 1)
}

func (head *Node) LastNode() *Node {
	if head.Next == nil {
		return head
	}

	return head.Next.LastNode()
}

func (head *Node) Pop() *Node {
	if head.Next.Next == nil {
		temp := head.Next
		head.Next = nil
		return temp
	}

	return head.Next.Pop()
}

func (act *Node) ReverseList(prev *Node) *Node {
	Next := act.Next
	act.Next = prev

	if Next == nil {
		return act
	}

	return Next.ReverseList(act)
}

func CreateListFromSlice(arr []int, index int, head, act *Node) *Node {
	if index == 0 {
		head = CreateNode(arr[index])
		act = head
	} else {
		act.AddNode(arr[index])
		act = act.Next
	}

	if index == len(arr)-1 {
		return head
	}

	return CreateListFromSlice(arr, index+1, head, act)
}

func CreateList(Values ...int) *Node {
	var head, act *Node

	for idx, val := range Values {
		if idx == 0 {
			head = CreateNode(val)
			act = head
			continue
		}

		act.AddNode(val)
		act = act.Next
	}

	return head
}
