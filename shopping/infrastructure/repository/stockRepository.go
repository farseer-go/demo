package repository

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/redis"
	"shopping/domain/stock"
	"strconv"
)

const stockKey = "farseer_go_stock"

func InitStock() {
	container.Register(func() stock.Repository {
		return &StockRepository{}
	})
}

type StockRepository struct {
	Redis redis.IClient `inject:"default"` // 使用farseer.yaml的Redis.default配置节点，并自动注入
}

func (receiver *StockRepository) Get(productId int64) int {
	stockVal, _ := receiver.Redis.HashGet(stockKey, strconv.FormatInt(productId, 10))
	return parse.Convert(stockVal, 0)
}

func (receiver *StockRepository) GetAll() map[int64]int {
	all, _ := receiver.Redis.HashGetAll(stockKey)
	result := make(map[int64]int)
	for k, v := range all {
		result[parse.Convert(k, int64(0))] = parse.Convert(v, 0)
	}
	return result
}

func (receiver *StockRepository) Set(productId int64, val int) int {
	stockVal, _ := receiver.Redis.HashIncrInt(stockKey, strconv.FormatInt(productId, 10), val)
	return stockVal
}
