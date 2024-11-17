package main

import (
	"fmt"
	"os"
	"bufio"
)
//prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
// It also print the names of all files in which each duplicated line occurs.

func main() {

	counts := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		// if files is not provided then input the text in terminal
		countLines(os.Stdin, counts)
	} else{
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
	            	countLines(f, counts)
			f.Close()
		}
	}

	for line, files := range counts{
		if len(files) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(files),line, files)
		}

	}

}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
}
