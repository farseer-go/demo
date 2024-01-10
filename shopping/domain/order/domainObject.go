package order

import "time"

// DomainObject 商品
type DomainObject struct {
	Id             int64     // 订单ID
	ProductId      int64     // 商品ID
	ProductCaption string    // 商品名称
	ProductImgSrc  string    // 商品图片
	ProductPrice   float64   // 价格
	ProductCount   int       // 商品数量
	CreateAt       time.Time // 下单时间
	CreateId       int64     // 下单人
}
