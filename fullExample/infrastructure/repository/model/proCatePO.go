package model

// ProCatePO 产品分类
type ProCatePO struct {
	Id   int    `gorm:"primaryKey;autoIncrement;comment:商品分类ID"`
	Name string `gorm:"size:32;not null;comment:商品分类名称"`
}
