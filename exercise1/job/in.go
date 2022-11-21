package job

import "github.com/go-gota/gota/series"

// Deprecated: Use series.In comparator instead.
func in(values ...string) func(el series.Element) bool {
	stack := make(map[string]struct{}, len(values))

	for _, val := range values {
		stack[val] = struct{}{}
	}

	return func(el series.Element) bool {
		if el.Type() == series.String {
			if val, ok := el.Val().(string); ok {
				if _, ok := stack[val]; ok {
					return ok
				}
			}
		}
		return false
	}
}
