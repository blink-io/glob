package match

import (
	"fmt"
	"unicode/utf8"
)

type Super struct{}

func (self Super) Match(s string) bool {
	return true
}

func (self Super) Len() int {
	return -1
}

func (self Super) Index(s string) (int, []int) {
	segments := make([]int, utf8.RuneCountInString(s))
	for i := range s {
		segments = append(segments, i)
	}

	segments = append(segments, len(s))

	return 0, segments
}

func (self Super) Kind() Kind {
	return KindSuper
}

func (self Super) String() string {
	return fmt.Sprintf("[super]")
}