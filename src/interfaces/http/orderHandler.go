package http

import (
	"net/http"
	"strconv"
	"transaction/src/domains/order"
	"transaction/src/helpers"

	"github.com/labstack/echo"
)

type orderHandler struct {
	OrderBusiness order.OrderBusiness
}

func NewOrderHandler(echo *echo.Echo, useCase order.OrderBusiness) {
	handler := &orderHandler{
		OrderBusiness: useCase,
	}

	echo.GET("/order", handler.FetchOrder)
}

func (handler *orderHandler) FetchOrder(contex echo.Context) (err error) {
	numS := contex.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := contex.QueryParam("cursor")
	ctx := contex.Request().Context()

	listAr, nextCursor, err := handler.OrderBusiness.FetchOrder(ctx, cursor, int64(num))
	if err != nil {
		return contex.JSON(
			helpers.GetStatusCode(err),
			helpers.ResponseError{Message: err.Error()},
		)
	}

	contex.Response().Header().Set(`X-Cursor`, nextCursor)

	return contex.JSON(http.StatusOK, listAr)
}
