package match

import (
	"reflect"
	"testing"
)

func TestPrefixIndex(t *testing.T) {
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
			[]int{2, 3},
		},
		{
			"ab",
			"fffabfff",
			3,
			[]int{2, 3, 4, 5},
		},
	} {
		p := Prefix{test.prefix}
		index, segments := p.Index(test.fixture, []int{})
		if index != test.index {
			t.Errorf("#%d unexpected index: exp: %d, act: %d", id, test.index, index)
		}
		if !reflect.DeepEqual(segments, test.segments) {
			t.Errorf("#%d unexpected segments: exp: %v, act: %v", id, test.segments, segments)
		}
	}
}

func BenchmarkIndexPrefix(b *testing.B) {
	m := Prefix{"qew"}
	in := acquireSegments(len(bench_pattern))

	for i := 0; i < b.N; i++ {
		m.Index(bench_pattern, in[:0])
	}
}

func BenchmarkIndexPrefixParallel(b *testing.B) {
	m := Prefix{"qew"}
	in := acquireSegments(len(bench_pattern))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Index(bench_pattern, in[:0])
		}
	})
}
