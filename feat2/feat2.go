package feat2

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	leftSorted := mergeSort(arr[:len(arr)/2])
	rightSorted := mergeSort(arr[len(arr)/2:])
	fmt.Println("eieiza")
	fmt.Println("lnwza007")
	return merge(leftSorted, rightSorted)
}

func merge(arr1 []int, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	idx_arr1 := 0
	idx_arr2 := 0
	idx_result := 0
	for idx_arr1 < len(arr1) && idx_arr2 < len(arr2) {
		if arr1[idx_arr1] < arr2[idx_arr2] {
			result[idx_result] = arr1[idx_arr1]
			idx_arr1++
		} else {
			result[idx_result] = arr2[idx_arr2]
			idx_arr2++
		}
		idx_result++
	}
	for idx_arr1 < len(arr1) {
		result[idx_result] = arr1[idx_arr1]
		idx_arr1++
		idx_result++
	}
	for idx_arr2 < len(arr2) {
		result[idx_result] = arr2[idx_arr2]
		idx_arr2++
		idx_result++
	}
	fmt.Println(result)
	return result
}

func bubbleSort(arr []int) []int {
	isSort := false
	for !isSort {
		isSort = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSort = false
			}
		}
	}
	return arr
}
