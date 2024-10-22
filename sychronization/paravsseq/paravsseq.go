package paravsseq

import "sync"

const max = 2048

func ParallelMergeSortV2(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2

	if len(s) < max {
		SequentialMergeSort(s)
	} else {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			ParallelMergeSortV2(s[:middle])
		}()

		go func() {
			defer wg.Done()
			ParallelMergeSortV2(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

func ParallelMergeSortV1(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ParallelMergeSortV1(s[:middle])
	}()

	go func() {
		defer wg.Done()
		ParallelMergeSortV1(s[middle:])
	}()

	wg.Wait()
	merge(s, middle)
}

func SequentialMergeSort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2

	SequentialMergeSort(s[:middle])
	SequentialMergeSort(s[middle:])

	merge(s, middle)
}

func merge(s []int, middle int) {
	result := make([]int, len(s))

	i := 0
	j := middle
	for k := 0; k < len(s); k++ {
		if i < middle && (j >= len(s) || s[i] <= s[j]) {
			result[k] = s[i]
			i++
		} else {
			result[k] = s[j]
			j++
		}
	}

	for k := 0; k < len(s); k++ {
		s[k] = result[k]
	}
}
