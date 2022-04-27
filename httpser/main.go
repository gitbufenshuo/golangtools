package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	if len(os.Args) == 1 {
		e := echo.New()
		e.Static("/", ".")
		e.Start("0.0.0.0:9998")
	} else {
		e := echo.New()
		e.Static("/", os.Args[1])
		fmt.Println(os.Args[2])
		e.Start(os.Args[2])
	}
}
