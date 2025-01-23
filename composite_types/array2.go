package main

import "fmt"

func main() {
   num := [3]int{1, 2, 3}

   fmt.Println(num)
   changeLocal(&num)
   fmt.Println(num)
   y := 5
   fmt.Println(y)
   changeIntValue(&y)
   fmt.Println(y)


}

func changeLocal(num *[3]int) {
	//*num[1] = 78
	num[1] = 78
	//fmt.Println("inside function ", num)
	fmt.Println("inside function", *num)
}

func changeIntValue(x *int) {
	*x = 678
	//x = 678
	fmt.Println("Value Inside ", *x)
}
