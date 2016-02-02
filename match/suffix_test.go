package match

import (
	"reflect"
	"testing"
)

func TestSuffixIndex(t *testing.T) {
	for id, test := range []struct {
		prefix   string
		fixture  string
		index    int
		segments []int
	}{
		{
			"ab",
			"abc",
			0,
			[]int{2},
		},
		{
			"ab",
			"fffabfff",
			0,
			[]int{5},
		},
	} {
		p := Suffix{test.prefix}
		index, segments := p.Index(test.fixture, []int{})
		if index != test.index {
			t.Errorf("#%d unexpected index: exp: %d, act: %d", id, test.index, index)
		}
		if !reflect.DeepEqual(segments, test.segments) {
			t.Errorf("#%d unexpected segments: exp: %v, act: %v", id, test.segments, segments)
		}
	}
}

func BenchmarkIndexSuffix(b *testing.B) {
	m := Suffix{"qwe"}
	in := acquireSegments(len(bench_pattern))

	for i := 0; i < b.N; i++ {
		m.Index(bench_pattern, in[:0])
	}
}

func BenchmarkIndexSuffixParallel(b *testing.B) {
	m := Suffix{"qwe"}
	in := acquireSegments(len(bench_pattern))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Index(bench_pattern, in[:0])
		}
	})
}
