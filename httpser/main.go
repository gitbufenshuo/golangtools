package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", ".")
	fmt.Println("0.0.0.0:9999")
	e.Start("0.0.0.0:9999")
}
