package main

import "fmt"

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) //convert int to Bytecounter
	return len(p), nil
}




func main() {
//	type ByteCounter int
	var c ByteCounter
	c.Write([]byte("hello")) //this will take implicit & in case struct type 
	fmt.Println(c)// "5", = len("hello")

	c = 0
	var name  = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) //"12", = len("hello, Dolly")
}
