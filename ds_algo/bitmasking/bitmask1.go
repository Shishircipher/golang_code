package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%08b\n", 13)
	fmt.Printf("%08b\n", ^13)
	n := 13
	fmt.Printf("%08b\n", n >> 1)
	fmt.Println(n)
	fmt.Printf("%08b\n", n << 1)
	fmt.Println(n)

//	mask := 5 //  101
//	i := 2  // 10    , 101 | 100 = 100
//	fmt.Println(mask | (1 << i))

//	if ( (mask & (1 << i)) == 0) {
//		fmt.Println("done")
//		mask = mask | (1 << i)
//		if ( (mask & (1 << i)) != 0) {
//			fmt.Println(" Now it's done")
//		}
//	
//	}
	 mask := 5 // Binary: 101
  	 i := 2    // Binary: 10 (position 2, counting from 0)

   	 if (mask & (1 << i)) == 0 {
        	fmt.Println("The i-th bit is unset")
        	mask = mask | (1 << i)
   	 } else {
        	fmt.Println("The i-th bit is already set")
    	}

    	fmt.Println("Final mask:", mask)
	// Initial mask
	mask1 := 0b1101 // Binary representation: 13 in decimal
	j := 2         // We want to unset the 2nd bit (counting from 0)

	// Unmask the 2nd bit
	mask1 = mask1 & ^(1 << j)

	fmt.Printf("Updated mask: %b\n", mask1) // Output: 1001 (binary), 9 (decimal)


}
