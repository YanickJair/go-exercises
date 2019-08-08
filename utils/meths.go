package utils

import (
	"errors"
	"fmt"
)

var errInvalidNode = errors.New("Invalid node")

// Node - Define our interface
type Node interface {
	SetValue(v int) error
	GetValue() (int, error)
}

// SSLNode - node structer
type SSLNode struct {
	next         *SSLNode
	value        int
	SNodeMessage string
}

// PowerNode - node structer
type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

// SingleLinkedList - pratice some list
type SingleLinkedList struct {
	head *SSLNode
	tail *SSLNode
}

// SetValue - create a set function
func (sNode *SSLNode) SetValue(v int) error {
	if sNode == nil {
		return errInvalidNode
	}
	sNode.value = v
	return nil
}

// GetValue - retrieve a set function
func (sNode *SSLNode) GetValue() (int, error) {
	if sNode == nil {
		return -1, errInvalidNode
	}
	return sNode.value, nil
}

// NewSSLNode - create new one
func NewSSLNode() *SSLNode {
	return &SSLNode{SNodeMessage: "This is the new SSLNode"}
}

// NewPowerNode - create a new nil PowerNode
func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is the new PowerNode"}
}

// SetValue - set PowerNode value
func (node *PowerNode) SetValue(v int) error {
	if node == nil {
		return errInvalidNode
	}
	node.value = v
	return nil
}

// GetValue - return PowerNode value
func (node *PowerNode) GetValue() (int, error) {
	if node == nil {
		return -1, errInvalidNode
	}
	return node.value, nil
}

// NewSingleLinkedList - create a nil SingleLinkedList
func NewSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

// Add - add item to the list
func (list *SingleLinkedList) Add(v int) {
	node := &SSLNode{value: v}
	if list.head == nil {
		list.head = node
	} else if list.tail == list.head {
		list.head.next = node
	} else if list.tail != nil {
		list.tail.next = node
	}
	list.tail = node
}

// String - implements the string interface for our list
func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		val, _ := n.GetValue()
		s += fmt.Sprintf(" {%d} ", val)
	}
	return s
}

// CreateNode - create a new Node
func CreateNode(v int) Node {
	sn := NewSSLNode()
	sn.SetValue(v)
	return sn
}

/*
	// implimenting the interface
	var node utils.Node
	node = utils.NewSSLNode()
	node.SetValue(10)
	fmt.Println(node.GetValue())

	node = utils.NewPowerNode()
	node.SetValue(20)
	fmt.Println(node.GetValue())

	list := utils.NewSingleLinkedList()
	list.Add(3)
	list.Add(4)
	list.Add(9)
	list.Add(6)

	node := utils.CreateNode(20)
	switch tp := node.(type) {
	case *utils.SSLNode:
		val, _ := tp.GetValue()
		fmt.Println("Type SSLNode: ", val)
	case *utils.PowerNode:
		val, _ := tp.GetValue()
		fmt.Println("Type SSPowerNodeLNode: ", val)
	default:
		fmt.Println("No type assigned to this Node")
	}
*/
