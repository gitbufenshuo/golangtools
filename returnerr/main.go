package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(0)
	}
	a := os.Args[1]
	aint, err := strconv.Atoi(a)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(aint % 256)
}
