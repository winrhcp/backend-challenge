package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func summaryHandler(c echo.Context) error {
	text, err := fetchBacon()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
	}
	fmt.Println("text = ", text)

	counter := countMeats(text)

	response := Response{Beef: counter}
	return c.JSON(http.StatusOK, response)
}
