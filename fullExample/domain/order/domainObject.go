package order

import "time"

// DomainObject 产品
type DomainObject struct {
	Id             int64     // 订单ID
	ProductId      int64     // 产品ID
	ProductCaption string    // 产品名称
	ProductImgSrc  string    // 产品图片
	ProductPrice   float32   // 价格
	ProductCount   int       // 产品数量
	CreateAt       time.Time // 下单时间
	CreateId       int64     // 下单人
}
