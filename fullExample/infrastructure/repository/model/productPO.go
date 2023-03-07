package model

type ProductPO struct {
	Id       int64   `gorm:"primaryKey;autoIncrement;comment:产品ID"`
	CateId   int     `gorm:"type:int;not null;index:idx_cateId;comment:商品分类ID"`
	CateName string  `gorm:"size:32;not null;comment:商品分类名称"`
	Caption  string  `gorm:"size:32;not null;comment:产品名称"`
	Desc     string  `gorm:"size:512;not null;comment:产品描述"`
	ImgSrc   string  `gorm:"size:256;not null;comment:产品图片"`
	Price    float32 `gorm:"size:10;not null;index:idx_price;comment:价格"`
}
