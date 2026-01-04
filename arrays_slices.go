package main

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 100 // Modifies the copy
}

func modifySlice(slice []int) {
	slice[0] = 100 // Modifies the original
}

func arraySlices() {
	// Array
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println(arr) // [1 2 3] - unchanged

	// Slice
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println(slice) // [100 2 3] - changed

	// Growing a slice
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println(nums) // [1 2 3 4 5]
}
