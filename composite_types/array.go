package main
import "fmt"

func main() {
fmt.Println("hello world")

var a [3]int //array of 3 integer
fmt.Println(a[0]) // print the first element
fmt.Println(a[len(a) -1]) // print the last element, a[2]
// Print the indices and elements
for i, v := range a {  // this format is an example of mapping the fucntion 
	//fmt.Println("%d %d \n", i, v)  what happen this uncomment , read about Println
	fmt.Printf("%d %d \n", i,v)
}
// print the elements 
for _, v := range(a) {
	fmt.Printf("%d \n",v)

}
// array literal to initialize an array with a list of values:
//var b [4]int = [4]int {1, 22, 11,}

//using an ellipsis "..."
c := [...]int {88, 78, 99}
fmt.Printf("%T \n", c )

// q := [3]int{1, 2, 3}
//q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
r := [...]int{98: -1} // look into this syntax ?
fmt.Printf("%T\n", r)
//fmt.Printf("%d %d", r[55], r[99])
}


type Currency int
const (
	USD Currency = iota
	EUR
)
symbol := [...]string{USD: "$", EUR: "EU"}
fmt.Println(USD, symbol[USD])

