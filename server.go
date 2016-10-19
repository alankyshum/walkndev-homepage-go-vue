package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {

	// CONFIG
	_configByte := []byte(`{"SERVER_PORT": "8000"}`)
	var config map[string]interface{}
	if err := json.Unmarshal(_configByte, &config); err != nil {
		panic(err)
	}
	var port string = config["SERVER_PORT"].(string)

	// SERVER INIT
	e := echo.New()
	// routing
	e.GET("/todo/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	e.PUT("/todo/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
	e.DELETE("/todo/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Tasks") })
	// start server
	fmt.Printf("Server running at port %s", port)
	e.Run(standard.New(":" + port))
}
