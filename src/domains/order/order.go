package order

import (
	"context"
	"transaction/src/models"
)

type OrderBusiness interface {
	FetchOrder(ctx context.Context, cursor string, num int64) (res []models.Order, nextCursor string, err error)
}

type OrderDao interface {
	Fetch(ctx context.Context, cursor string) (res []models.Order, err error)
}
