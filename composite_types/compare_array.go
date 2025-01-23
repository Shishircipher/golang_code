package main

import "fmt"

func main(){
fmt.Println("exmaple of comparing two arrays or more")

a := [3]int{2, 5, 7}
b := [...]int{2, 5, 7}
c_size := 2
c := [c_size]int{4,5} // fix the size of array using variable ?
fmt.Println(a == b)

fmt.Println(a == c)
}
