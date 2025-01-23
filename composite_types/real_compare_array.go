package main 

import "fmt"
import "crypto/sha256"

func main() {
   c1 := sha256.Sum256([]byte("X"))
   c2 := sha256.Sum256([]byte("Y"))

   fmt.Printf("%x \n %x \n %t \n %T \n ", c1, c2, c1 == c2, c1)
}
