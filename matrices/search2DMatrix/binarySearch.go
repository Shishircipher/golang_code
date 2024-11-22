//using iterative approach

package main

import "fmt"

func main() {

	slice1 := []int{23, 54, 68, 79}

	fmt.Println(slice1)

//	slice2 := [...]int{1, 23, 56}

	fmt.Println(binarySearch(slice1, 54))

}

/* func binarySearch(array []int, x int)  bool {
	length := len(array)
        
	low := 0
	high := length
	mid := (low + high) / 2 // logic error led to infinite loop
        
	
	for low <= high {

		if x == array[mid] {
			//return 0
			return true
		} else if x < array[mid] {
			 high = mid - 1
		} else if x > array[mid] {
			low = mid + 1
		} else {
		//	return -1
		        return false
		}
		
	}
	return 
} */

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

