package bptree

type Item interface {
	Key() int
	Less(than Item) bool
}

type Int int

func (a Int) Key() int {
	return int(a)
}
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}
