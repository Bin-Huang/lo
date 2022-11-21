package lo

import (
	"testing"

	fuzz "github.com/google/gofuzz"
)

func FuzzSlice(f *testing.F) {
	f.Add([]byte("a0000000700000000000000000000000\x00"))
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var (
			f32 []float32
			i int
			j int
		)
		fuzzer.Fuzz(&f32)
		fuzzer.Fuzz(&i)
		fuzzer.Fuzz(&j)
		Slice(f32, i, j)
	})
}

func FuzzFilter(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var f32 []float32
		f := func (v float32, index int) bool {
			return true
		}
		fuzzer.Fuzz(&f32)
		Filter(f32, f)
	})
}

func FuzzMap(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var f32 []float32
		f := func (v float32, index int) string {
			return "x"
		}
		fuzzer.Fuzz(&f32)
		Map(f32, f)
	})
}

func FuzzFilterMap(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var f32 []float32
		f := func (v float32, index int) (string, bool) {
			return "x", true
		}
		fuzzer.Fuzz(&f32)
		FilterMap(f32, f)
	})
}

func FuzzFlatMap(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var f32 []float32
		f := func (v float32, index int) []string {
			return []string{"x"}
		}
		fuzzer.Fuzz(&f32)
		FlatMap(f32, f)
	})
}

func FuzzReduce(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		f := func(agg int, item int, _ int) int {
			return agg + item
		}
		var arr []int
		var r int
		fuzzer.Fuzz(&arr)
		fuzzer.Fuzz(&r)
		Reduce(arr, f, r)
	})
}

func FuzzReduceRight(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		f := func(agg int, item int, _ int) int {
			return agg + item
		}
		var arr []int
		var r int
		fuzzer.Fuzz(&arr)
		fuzzer.Fuzz(&r)
		ReduceRight(arr, f, r)
	})
}

func FuzzForEach(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		f := func(agg int, item int) {}
		var arr []int
		fuzzer.Fuzz(&arr)
		ForEach(arr, f)
	})
}

func FuzzTimes(f *testing.F) {
	// TODO:
	// f.Add([]byte("0"))
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var count int
		f := func (index int) string {
			return "x"
		}
		fuzzer.Fuzz(&count)
		println("count", count)
		Times(count, f)
	})
}

func FuzzUniq(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr []int
		fuzzer.Fuzz(&arr)
		Uniq(arr)
	})
}

func FuzzUniqBy(f *testing.F) {
	f.Add([]byte("0"))
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr []int
		f := func (index int) string {
			return "x"
		}
		fuzzer.Fuzz(&arr)
		UniqBy(arr, f)
	})
}

func FuzzGroupBy(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr []int
		f := func (index int) string {
			return "x"
		}
		fuzzer.Fuzz(&arr)
		GroupBy(arr, f)
	})
}

func FuzzChunk(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr []int
		var count int
		fuzzer.Fuzz(&arr)
		fuzzer.Fuzz(&count)
		if count > 0 {
			Chunk(arr, count)
		}
	})
}

func FuzzPartitionBy(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr []int
		f := func (index int) string {
			return "x"
		}
		fuzzer.Fuzz(&arr)
		PartitionBy(arr, f)
	})
}

func FuzzFlatten(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr [][]int
		fuzzer.Fuzz(&arr)
		Flatten(arr)
	})
}

func FuzzInterleave(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr [][]int
		fuzzer.Fuzz(&arr)
		Interleave(arr)
	})
}

func FuzzShuffle(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr [][]int
		fuzzer.Fuzz(&arr)
		Shuffle(arr)
	})
}

func FuzzReverse(f *testing.F) {
	f.Add([]byte("0"))
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr []int
		fuzzer.Fuzz(&arr)
		Reverse(arr)
	})
}

func FuzzReplace(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		fuzzer := fuzz.NewFromGoFuzz(buf)
		var arr []int
		var old int
		var new int
		var n int
		fuzzer.Fuzz(&arr)
		fuzzer.Fuzz(&old)
		fuzzer.Fuzz(&new)
		fuzzer.Fuzz(&n)
		Replace(arr, old, new, n)
	})
}
