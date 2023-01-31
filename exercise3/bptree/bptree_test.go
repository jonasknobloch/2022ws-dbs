package bptree

import (
	"fmt"
	"testing"
)

func TestBPTree_Insert(t *testing.T) {
	bpt := NewBPTree(3)

	bpt.Insert(Int(5))
	bpt.Insert(Int(15))
	bpt.Insert(Int(25))
	bpt.Insert(Int(35))
	bpt.Insert(Int(45))
	bpt.Insert(Int(55))

	fmt.Println(bpt.String())
}
