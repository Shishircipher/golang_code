package main

import (
	"fmt"
)

func main() {
	twoDSlice1 := make([][]int, 3)
	fmt.Printf("Number of rows in slice rows : %d\n", len(twoDSlice1[0]))
	for i := range twoDSlice1{
		twoDSlice1[i] = make([]int, 3)
	}
	fmt.Printf("Number of rows in slice : %d\n", len(twoDSlice1))
        fmt.Printf("Number of rows in slice rows : %d\n", len(twoDSlice1[0]))
	
	for i, _ := range twoDSlice1{
		for _, val := range twoDSlice1[i] {
			fmt.Printf(" %d", val)
		}}
		fmt.Println()
	for _, row := range twoDSlice1 {
        for _, val := range row {
            fmt.Println(val)
        }
    }
}
