package main

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "fmt"
//    "os"
)

// Create a Customer type
type Customer struct {
    Name string
    Age  int
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to JSON, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.
func (c *Customer) WriteJSON(w io.Writer) error {
    js, err := json.Marshal(c)
    if err != nil {
        return err
    }

    _, err = w.Write(js)
    return err
}

func main() {
    // Initialize a customer struct.
    c := &Customer{Name: "Alice", Age: 21}

    // We can then call the WriteJSON method using a buffer...
    var buf bytes.Buffer
    fmt.Println(buf)
    err := c.WriteJSON(&buf)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(buf.String())

}
