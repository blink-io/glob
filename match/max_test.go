package match

import (
	"reflect"
	"testing"
)

func TestMaxIndex(t *testing.T) {
	for id, test := range []struct {
		limit    int
		fixture  string
		index    int
		segments []int
	}{
		{
			3,
			"abc",
			0,
			[]int{0, 1, 2, 3},
		},
		{
			3,
			"abcdef",
			0,
			[]int{0, 1, 2, 3},
		},
	} {
		p := Max{test.limit}
		index, segments := p.Index(test.fixture, []int{})
		if index != test.index {
			t.Errorf("#%d unexpected index: exp: %d, act: %d", id, test.index, index)
		}
		if !reflect.DeepEqual(segments, test.segments) {
			t.Errorf("#%d unexpected segments: exp: %v, act: %v", id, test.segments, segments)
		}
	}
}

func BenchmarkIndexMax(b *testing.B) {
	m := Max{10}
	in := acquireSegments(len(bench_pattern))

	for i := 0; i < b.N; i++ {
		m.Index(bench_pattern, in[:0])
	}
}

func BenchmarkIndexMaxParallel(b *testing.B) {
	m := Max{10}
	in := acquireSegments(len(bench_pattern))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Index(bench_pattern, in[:0])
		}
	})
}
