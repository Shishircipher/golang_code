package main

import "fmt"
import "time"

func main() {
	go spinner(100 * time.Millisecond)

	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/`{
			fmt.Printf("\r %c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	
	return fib(x - 1) + fib(x - 2) // if here logic error happens then function can run infinitetly then it shows signal killed , may be golang implements stacks , resource limited ,check  what happen if this happen in c or rust .
}

