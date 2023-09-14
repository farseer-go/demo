package repository

import (
	"github.com/farseer-go/fs/container"
	"shopping/domain/order"
)

func InitRepository() {
	container.Register(func() order.Repository {
		return &OrderRepository{}
	})
}
