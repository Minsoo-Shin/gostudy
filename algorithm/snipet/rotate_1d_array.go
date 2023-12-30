package snipet

func rotateLeftByOne1DArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	temp := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = temp
	return arr
}

func rotateRightByOne1DArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	temp := arr[len(arr)-1]
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = temp
	return arr
}

/*
(before) [1,2,3,4,5] , count = 3
(after) [4,5,1,2,3]

(before) [1,2,3,4,5] , count = 5
(after) [1,2,3,4,5]
*/
func rotateLeftByCount(arr []int, count int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	count %= len(arr)
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		count := (count + i) % len(arr)
		result[i] = arr[count]
	}
	return result
}
