package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// append by line
// default two lines and with ' '
func main() {
	lineSpec := 2
	splitSym := " "
	scanner := bufio.NewScanner(os.Stdin)
	buffer := new(bytes.Buffer)
	buffer.Reset()
	suc := 0
	for scanner.Scan() {
		text := scanner.Text()
		buffer.WriteString(text + splitSym)
		suc++
		if suc == lineSpec {
			fmt.Println(buffer.String())
			buffer.Reset()
			suc = 0
		}
	}
	if buffer.Len() != 0 {
		fmt.Println(buffer.String())
	}
}
