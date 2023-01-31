package bptree

import "fmt"

type Leaf struct {
	items []Item
	next  *Leaf
}

func (l *Leaf) Keys() []int {
	keys := make([]int, len(l.items))

	for i, v := range l.items {
		keys[i] = v.Key()
	}

	return keys
}

func (l *Leaf) Split() *Interior {
	median := len(l.items) / 2

	right := &Leaf{
		items: l.items[median:],
		next:  nil,
	}

	l.items = l.items[:median]
	l.next = right

	return &Interior{
		keys:     []int{right.Keys()[0]},
		children: []Node{l, right},
	}
}

func (l *Leaf) String() string {
	return fmt.Sprint(l.items)
}
