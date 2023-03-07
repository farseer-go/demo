package model

// OrderPO 订单
type OrderPO struct {
	Id             int64   `gorm:"primaryKey;autoIncrement;comment:订单ID"`
	ProductId      int     `gorm:"type:int;not null;comment:产品ID"`
	ProductCaption string  `gorm:"size:32;not null;comment:产品名称"`
	ProductDesc    string  `gorm:"size:512;not null;comment:产品描述"`
	ProductImgSrc  string  `gorm:"size:256;not null;comment:产品图片"`
	ProductPrice   float32 `gorm:"size:10;not null;index:idx_price;comment:价格"`
	ProductCount   int     `gorm:"type:int;not null;comment:产品ID"`
}
