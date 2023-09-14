package repository

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/parse"
	"shopping/domain/stock"
	"shopping/infrastructure/repository/context"
	"strconv"
)

const stockKey = "farseer_go_stock"

func InitStock() {
	container.Register(func() stock.Repository {
		return &StockRepository{}
	})
}

// StockRepository 库存仓储
// @inject stock.Repository
type StockRepository struct {
}

func (receiver *StockRepository) Get(productId int64) int {
	stockVal, _ := context.RedisContext.HashGet(stockKey, strconv.FormatInt(productId, 10))
	return parse.Convert(stockVal, 0)
}

func (receiver *StockRepository) GetAll() map[int64]int {
	all, _ := context.RedisContext.HashGetAll(stockKey)
	result := make(map[int64]int)
	for k, v := range all {
		result[parse.Convert(k, int64(0))] = parse.Convert(v, 0)
	}
	return result
}

func (receiver *StockRepository) Set(productId int64, val int) int {
	stockVal, _ := context.RedisContext.HashIncrInt(stockKey, strconv.FormatInt(productId, 10), val)
	return stockVal
}
