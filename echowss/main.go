package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

/*
	生成你自己的 pem 和 key 放在相应的路径
	使用 generate-tls-cert 工具
	https://github.com/Shyp/generate-tls-cert
    such as : generate-tls-cert --host=127.0.0.1,localhost
*/
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", hello)
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "h")
	})
	e.Logger.Fatal(e.StartTLS(":443", "./leaf.pem", "./leaf.key"))
}
