package product

// DomainObject 商品
type DomainObject struct {
	Id      int64   // 商品ID
	Cate    CateEO  // 商品分类
	Stock   int     // 库存
	Caption string  // 商品名称
	Desc    string  // 商品描述
	ImgSrc  string  // 商品图片
	Price   float64 // 价格
}
