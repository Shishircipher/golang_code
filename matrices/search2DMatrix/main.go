//https://leetcode.com/problems/search-a-2d-matrix/description/

package main 

import "fmt"


func main() {
    twoDarray1 := [][]int{{1, 3, 5, 7},{10, 11, 16, 20}, {23, 30, 34, 60}}

    fmt.Println(twoDarray1)

    fmt.Println(searchMatrix(twoDarray1, 16))
    fmt.Println(searchMatrix(twoDarray1, 60))

}

func searchMatrix(matrix [][]int, target int) bool {
    // m := len(matrix)
     n := len(matrix[0])
     
     for _, val := range matrix {
	     if target == val[n-1] {
		     return true
	     } else if target < val[n-1] {
		     //return binarySearch(val, target)
		     if binarySearch(val, target) == true {
			     return true
		     } else {
			     continue
	             }

	     } else {
		    // return binarySearch(val, target)
		    //return binarySearch(val, target)
                     if binarySearch(val, target) == true {
                             return true
                     } else {
                             continue
                     }
	     }
     }
     return false
}

func binarySearch(array []int, x int) bool {
    low := 0
    high := len(array) - 1

    for low <= high {
        mid := (low + high) / 2

        if x == array[mid] {
            return true
        } else if x < array[mid] {
            high = mid - 1
  //      } else if x > array[mid]{
  //          low = mid + 1
   //     } else {
  //		return false
	} else {
		low = mid + 1
	}
    }

    return false // Return false if the value is not found
}
