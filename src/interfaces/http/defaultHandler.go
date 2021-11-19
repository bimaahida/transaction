package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type defaultHandler struct{}

func NewDefaultHandler(echo *echo.Echo) {
	handler := &defaultHandler{}

	echo.GET("/", handler.DefaultUrl)
}

func (handler *defaultHandler) DefaultUrl(ctx echo.Context) error {
	ctx.Logger().Info("Hit Base Url")
	return ctx.String(http.StatusOK, "Service listing in [::]:8080")
}
