//print the text of each line that appears more than
//once in the standard input,preceded by its count
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)

	cn := 0
	//NewScanner returns a new Scanner to read from r
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//Text return the most recent token  generated by a call
		counts[input.Text()]++

		cn++
		if cn >= 5 {
			break
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}