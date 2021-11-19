package main

import (
	"time"

	"transaction/config"
	"transaction/database"
	_orderDomain "transaction/src/domains/order"
	_handler "transaction/src/interfaces/http"

	"github.com/labstack/echo"
)

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		panic(err)
	}

	dbConnection, _ := database.ConnectDB(conf.Database)

	if err != nil {
		panic(err)
	}

	echo := echo.New()

	timeoutContext := time.Duration(conf.Context.Timeout) * time.Second
	orderDao := _orderDomain.NewOrderDao(dbConnection)
	orderBusiness := _orderDomain.NewOrderBusiness(orderDao, timeoutContext)

	_handler.NewOrderHandler(echo, orderBusiness)
	_handler.NewDefaultHandler(echo)

	echo.Logger.Fatal(echo.Start(conf.Server.Address))
}
