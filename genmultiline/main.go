package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// how to use: cat test.txt | go run main.go
func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		io.WriteString(os.Stderr, "read_err")
		os.Exit(1)
	}
	o := strings.Replace(string(raw), "\n", `\n`, -1)
	fmt.Println(o)
}
