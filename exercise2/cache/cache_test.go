package cache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	lru1 := NewLRUK(8, 1)
	fifo := NewFIFO(8)
	lifo := NewLIFO(8)

	for i := 0; i < 8; i++ {
		lru1.Put(i, i+1)
		fifo.Put(i, i+1)
		lifo.Put(i, i+1)
	}

	fmt.Println("LRU1", Serialize(lru1))
	fmt.Println("FIFO", Serialize(fifo))
	fmt.Println("LIFO", Serialize(lifo))

	lru1.Get(1)
	fifo.Get(1)
	lifo.Get(1)

	fmt.Println("LRU1", Serialize(lru1))
	fmt.Println("FIFO", Serialize(fifo))
	fmt.Println("LIFO", Serialize(lifo))

	lru1.Put(9, 10)
	fifo.Put(9, 10)
	lifo.Put(9, 10)

	fmt.Println("LRU1", Serialize(lru1))
	fmt.Println("FIFO", Serialize(fifo))
	fmt.Println("LIFO", Serialize(lifo))
}

func TestLRU2(t *testing.T) {
	lru2 := NewLRUK(3, 2)

	const A = 1
	const B = 2
	const C = 3
	const D = 4

	lru2.Put(A, 0)
	lru2.Put(B, 0)
	lru2.Put(C, 0)

	lru2.t = 50
	lru2.Get(C)
	lru2.t = 52
	lru2.Get(A)
	lru2.t = 55
	lru2.Get(B)
	lru2.t = 60
	lru2.Get(B)
	lru2.t = 62
	lru2.Get(A)
	lru2.t = 65
	lru2.Get(C)
	lru2.t = 70
	lru2.Get(B)
	lru2.t = 72
	lru2.Get(B)
	lru2.t = 75
	lru2.Get(C)
	lru2.t = 76
	lru2.Get(A)

	fmt.Println(Serialize(lru2))

	lru2.t = 80
	lru2.Put(D, 0)

	fmt.Println(Serialize(lru2))

	for i, v := range lru2.Log() {
		if v[0] == 0 {
			continue
		}

		var k string

		switch v[0] {
		case 1:
			k = "A"
		case 2:
			k = "B"
		case 3:
			k = "C"
		case 4:
			k = "D"
		default:
			k = "X"
		}

		fmt.Println(i, k)
	}
}
