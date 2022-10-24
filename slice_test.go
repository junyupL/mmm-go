package mmm

import (
	"testing"

	"github.com/junyupL/mmm-go/memory"
)

func BenchmarkDynArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := Make[int](memory.HeapAllocator, 1, 3)
		slice.Append(5)
		slice.Append(8)
		slice.Slice[0] = 23
		slice.Append(13)
		slice.Append(16)
		slice.Append(234)
		slice.Destruct()
	}
}

func BenchmarkENFDynArray(b *testing.B) {
	memory.HInit()
	enf := memory.InitENF(1000)
	slice := Make[int](&enf, 1, 3)
	for i := 0; i < b.N; i++ {

		slice.Append(5)
		slice.Append(8)
		slice.Slice[0] = 23
		for i := 0; i < b.N; i++ {
			slice.Append(234)
		}
		//slice.Destruct()
	}
	enf.Destruct()
}

func BenchmarkSlice(b *testing.B) {
	slice := make([]int, 1, 3)
	for i := 0; i < b.N; i++ {

		slice = append(slice, 5)
		slice = append(slice, 8)
		slice[0] = 23
		for i := 0; i < b.N; i++ {
			slice = append(slice, 234)
		}
	}
}

func BenchmarkENFDynArrayA(b *testing.B) {
	memory.HInit()
	enf := memory.InitENF(1000)
	slice := Make[int](&enf, 1, 3)
	for i := 0; i < b.N; i++ {

		slice.Append(5)

	}
	slice.Destruct()
	enf.Destruct()
}

func BenchmarkGCSliceA(b *testing.B) {

	slice := make([]int, 1, 3)
	for i := 0; i < b.N; i++ {

		slice = append(slice, 5)

	}
}
