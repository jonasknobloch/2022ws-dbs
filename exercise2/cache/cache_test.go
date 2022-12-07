package cache

import (
	"fmt"
	"testing"
)

func TestLRUK(t *testing.T) {
	lruk := NewLRUK(8)
	fifo := NewFIFO(8)
	lifo := NewLIFO(8)

	for i := 0; i < 8; i++ {
		lruk.Put(i, i+1)
		fifo.Put(i, i+1)
		lifo.Put(i, i+1)
	}

	fmt.Println(Serialize(lruk))
	fmt.Println(Serialize(fifo))
	fmt.Println(Serialize(lifo))

	lruk.Get(1)
	fifo.Get(1)
	lifo.Get(1)

	fmt.Println(Serialize(lruk))
	fmt.Println(Serialize(fifo))
	fmt.Println(Serialize(lifo))

	lruk.Put(9, 10)
	fifo.Put(9, 10)
	lifo.Put(9, 10)

	fmt.Println(Serialize(lruk))
	fmt.Println(Serialize(fifo))
	fmt.Println(Serialize(lifo))
}
