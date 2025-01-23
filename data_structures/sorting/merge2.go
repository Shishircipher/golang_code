package main

import "fmt"

// Main function
func main() {
	a := [...]int{3, 4, 2, 1} // Array of integers
	fmt.Printf("%T \n", a)    // Print the type of 'a' to show it's an array
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// Convert the array to a slice and call MergeSort
	MergeSort(a[:], 0, len(a)-1) // Pass the slice of 'a' from index 0 to the last element

	// Print the sorted array
	fmt.Println("Sorted array:", a)
}

// MergeSort function using slice
func MergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2 // Calculate the middle index

		// Recursively sort the left and right halves
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)

		// Merge the sorted halves
		Merge(arr, left, mid, right)
	}
}

// Merge function to merge two sorted halves of the slice
func Merge(arr []int, left, mid, right int) {
	// Determine the sizes of the two subarrays to be merged
	n1 := mid - left + 1
	n2 := right - mid

	// Create temporary slices for left and right halves
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	// Copy data to temporary slices
	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}

	// Merge the temporary slices back into arr[left...right]
	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	// Copy any remaining elements of leftArr, if any
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	// Copy any remaining elements of rightArr, if any
	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}
