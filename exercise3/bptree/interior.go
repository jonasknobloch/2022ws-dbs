package bptree

import "fmt"

type Interior struct {
	keys     []int
	children []Node // assert len(children) == len(keys)+1
}

func (i *Interior) Keys() []int {
	return i.keys
}

func (i *Interior) Split() *Interior {
	median := len(i.keys) / 2

	left := &Interior{
		keys:     i.keys[:median],
		children: i.children[:median+1],
	}

	right := &Interior{
		keys:     i.keys[median+1:],
		children: i.children[median+1:],
	}

	return &Interior{
		keys:     []int{i.keys[median]},
		children: []Node{left, right},
	}
}

func (i *Interior) String() string {
	return fmt.Sprint(i.keys)
}
