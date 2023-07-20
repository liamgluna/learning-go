// From Matt Holiday - Go Class: 07

package main

import (
	"fmt"
	"os"
	"io"
)

// go run . a.txt
func main() {
	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		file.Close()
	}
}
