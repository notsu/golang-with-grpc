package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/client", client)
	e.GET("/server", server)

	e.Logger.Fatal(e.Start(":8080"))
}

func client(c echo.Context) error {
	resp, err := http.Get("http://localhost:8080/server")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"got": string(body),
	})
}

func server(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Ground Control to Major Tom",
	})
}
