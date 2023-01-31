package bptree

import (
	"sort"
	"strings"
)

type BPTree struct {
	degree int
	root   Node
}

func NewBPTree(degree int) *BPTree {
	return &BPTree{
		degree: degree,
	}
}

func (bp *BPTree) Insert(item Item) {
	if bp.root == nil {
		bp.root = &Leaf{
			items: []Item{item},
			next:  nil,
		}

		return
	}

	var expand func(node Node) *Interior
	expand = func(node Node) *Interior {
		if leaf, ok := node.(*Leaf); ok {
			leaf.items = append(leaf.items, item)

			sort.Slice(leaf.items, func(i, j int) bool {
				return leaf.items[i].Less(leaf.items[j])
			})

			if len(leaf.items) == bp.degree {
				return leaf.Split()
			}

			return nil
		}

		interior, ok := node.(*Interior)

		if !ok {
			panic("unexpected node implementation")
		}

		next := len(interior.children) - 1

		for i, k := range interior.keys {
			if k > item.Key() {
				next = i
				break
			}
		}

		if insert := expand(interior.children[next]); insert != nil {
			if next == 0 {
				interior.keys = append(insert.keys, interior.keys...)
				interior.children = append(insert.children, interior.children[1:]...)
			} else {
				interior.keys = append(append(interior.keys[:next], insert.keys...), interior.keys[next:]...)
				interior.children = append(append(interior.children[:next], insert.children...), interior.children[next+1:]...)
			}
		}

		if len(interior.keys) == bp.degree {
			return interior.Split()
		}

		return nil
	}

	if root := expand(bp.root); root != nil {
		bp.root = root
	}
}

func (bp *BPTree) String() string {
	var sb strings.Builder

	var walk func(Node)
	walk = func(node Node) {
		sb.WriteString(" ")

		if _, ok := node.(*Leaf); ok {
			sb.WriteString(node.String())
			return
		}

		sb.WriteString("(")
		sb.WriteString(node.String())

		if i, ok := node.(*Interior); ok {
			for _, c := range i.children {
				walk(c)
			}
		}

		sb.WriteString(")")
	}

	walk(bp.root)

	return sb.String()[1:]
}
