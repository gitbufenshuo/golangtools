package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", os.Args[1])
	fmt.Println(os.Args[2])
	e.Start(os.Args[2])
}
