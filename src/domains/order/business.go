package order

import (
	"context"
	"time"
	"transaction/src/models"
)

type orderBusiness struct {
	orderDao       OrderDao
	contextTimeout time.Duration
}

func NewOrderBusiness(orderDao OrderDao, timeout time.Duration) OrderBusiness {
	return &orderBusiness{
		orderDao:       orderDao,
		contextTimeout: timeout,
	}
}

func (ob *orderBusiness) FetchOrder(ctx context.Context, cursor string, num int64) (res []models.Order, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, ob.contextTimeout)
	defer cancel()

	res, err = ob.orderDao.Fetch(ctx, cursor)
	if err != nil {
		return nil, "", err
	}

	return
}
