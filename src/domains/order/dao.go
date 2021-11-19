package order

import (
	"context"
	"fmt"
	"log"
	"transaction/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderDao struct {
	Conn *mongo.Collection
}

func NewOrderDao(Conn *mongo.Database) OrderDao {
	return &orderDao{
		Conn: Conn.Collection("transaction"),
	}
}

func (od *orderDao) Fetch(ctx context.Context, cur string) (res []models.Order, err error) {
	cursor, err := od.Conn.Find(ctx, bson.M{})

	res = make([]models.Order, 0)

	if err != nil {
		log.Fatal("Finding all documents ERROR:", err)

		return res, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order bson.M
		err := cursor.Decode(&order)

		fmt.Println(order)

		if err != nil {
			log.Fatal("Error Mapping Order : ", err)
		}

		// res = append(res, order)
	}

	return
}
