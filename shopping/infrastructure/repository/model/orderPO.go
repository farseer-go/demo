package model

import "time"

// OrderPO 订单
type OrderPO struct {
	Id             int64     `gorm:"primaryKey;comment:订单ID"`
	ProductId      int64     `gorm:"type:bigint;not null;comment:商品ID"`
	ProductCaption string    `gorm:"size:32;not null;comment:商品名称"`
	ProductImgSrc  string    `gorm:"size:256;not null;comment:商品图片"`
	ProductPrice   float64   `gorm:"size:10;not null;comment:订单价格"`
	ProductCount   int       `gorm:"type:int;not null;comment:商品数量"`
	CreateAt       time.Time `gorm:"type:timestamp;size:6;not null;comment:下单时间"`
	CreateId       int64     `gorm:"type:bigint;not null;comment:下单人"`
}
