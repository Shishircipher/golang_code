package main

import "fmt"

func main() { 
    a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    reverse_slice(a[:])

    fmt.Println(a) // how it change - is copy the slices or pointer of slices

    // let us experiment the with array

    b := [4]int{3, 5, 33, 55}
    //reverse_slice(b[])
   // reverse_array(b)
    fmt.Println(b)
}
func reverse_slice(s []int ) {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
                s[i], s[j] = s[j], s[i]
        }

    }
func reverse_array(s *[]int) {
	t := *s
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
                s[i], s[j] = s[j], s[i] }
}
