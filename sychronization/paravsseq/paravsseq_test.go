package paravsseq

import (
	"math/rand"
	"testing"
	"time"
)

// createTestSlice creates a random slice of size n
func createTestSlice(n int) []int {
	rand.Seed(time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC).UnixNano())
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(n)
	}
	return s
}

func BenchmarkSequentialMergeSort(b *testing.B) {
	s := createTestSlice(10240)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SequentialMergeSort(s)
	}
}

func BenchmarkParallelMergeSortV1(b *testing.B) {
	s := createTestSlice(10240)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParallelMergeSortV1(s)
	}
}

func BenchmarkParallelMergeSortV2(b *testing.B) {
	s := createTestSlice(10240)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParallelMergeSortV2(s)
	}
}
