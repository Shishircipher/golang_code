package main

import "fmt"

func main() { 
fmt.Println(" This is experiment with map data type ")

ages := map[string]int{
	"alice" : 31,
	"sam" : 24,
	"bob" : 0,
}

fmt.Println(ages)
fmt.Printf("%d %d %d\n", ages["alice"], ages["sam"], ages["bob"])  //here bob is not present , but my question is what happen when "bob" have 0 as the value

age, ok := ages["bob"]
fmt.Printf("%d\n", age)
if !ok {
	fmt.Println("bob is not present")

}



}
// Comparing the two map , like slices we can not compare the map using '==' operator
// the only legal comparison is with nil ?
func equal_map(x, y map[string]int ) bool {
        if len(x) != len(y) {
                return false
        }
        for k, xv := range x {
                if yv, ok := y[k]; !ok || yv != xv {
			return false }
                }
        return true
}
