package main

import "fmt"

func insertion_sort(arr []int, n int) {
	for i := 1; i <= n; i++ {
		key := arr[i]
		j := i - 1
		for j > 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	//return arr[]
}

func main() {
	a := [...]int{3, 4, 2, 1}
	insertion_sort(a[:], 3)
	fmt.Println(insertion_sort(a[:], 3))
}