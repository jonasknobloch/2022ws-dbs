package job

import (
	"github.com/go-gota/gota/series"
	"regexp"
	"strings"
)

func like(pattern string) func(el series.Element) bool {
	if ok, err := regexp.MatchString(`%[\w\d\s]+%`, pattern); !ok || err != nil {
		panic("unsupported pattern")
	}

	needle := pattern[1 : len(pattern)-1]

	return func(el series.Element) bool {
		if el.Type() == series.String {
			if val, ok := el.Val().(string); ok {
				return strings.Contains(val, needle)
			}
		}
		return false
	}
}
