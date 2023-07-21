// From Matt Holiday - Go Class: 07

package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	for _, fileName := range os.Args[1:] {
		var lineCount, wordCount, charCount int

		file, err := os.Open(fileName)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			s := scanner.Text()

			wordCount += len(strings.Fields(s))
			charCount += len(s)
			lineCount++
		}

		fmt.Printf("%7d %7d %7d %s\n", lineCount, wordCount, charCount, fileName)
		file.Close()
	}
}