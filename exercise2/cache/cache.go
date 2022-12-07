package cache

import "strings"

type Cache interface {
	Get(int) *Entry
	Put(int, int)
	Data() []*Entry
}

func Serialize(c Cache) string {
	var sb strings.Builder

	sb.WriteString("[")

	for i, v := range c.Data() {
		if i > 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(v.String())
	}

	sb.WriteString("]")

	return sb.String()
}
