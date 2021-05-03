package directory

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Node struct {
	Id     string
	Slug   string
	Parent *Node
	Name   string
}

type Tree struct {
	*Node
}

func (n *Node) generateId() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (n *Node) AddSubNode(name, slug string) *Node {
	node := &Node{
		Id:     n.generateId(),
		Slug:   slug,
		Parent: n,
		Name:   name,
	}

	return node
}

func (n *Node) GetRoot() *Node {
	cur := n
	for cur.Parent != nil {
		cur = cur.Parent
	}
	return cur
}

func (n *Node) Path() string {
	res := n.Name
	cur := n
	for cur.Parent != nil {
		cur = cur.Parent
		res = fmt.Sprintf("%s->%s", cur.Name, res)
	}
	return res
}

func CreateTree() *Tree {
	rand.Seed(time.Now().UnixNano())
	return &Tree{}
}
