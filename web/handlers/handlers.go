package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MainHandler(c echo.Context) error{
	return c.String(http.StatusOK, "Info is here")
}

func ShowAuth(c echo.Context) error{
	return c.Render(http.StatusOK, "auth.html", map[string]interface{}{
		"Title": "Auth",
	})
}
func HandlerMessages(c echo.Context) error{
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title":"Messages",
	})
}