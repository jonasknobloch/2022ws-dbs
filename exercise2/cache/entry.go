package cache

import "fmt"

type Entry struct {
	Key, Value int
}

func (e *Entry) String() string {
	return fmt.Sprintf("(%d, %d)", e.Key, e.Value)
}
