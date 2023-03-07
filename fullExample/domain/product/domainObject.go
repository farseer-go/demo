package product

// DomainObject 产品
type DomainObject struct {
	Id      int64   // 产品ID
	Cate    CateEO  // 商品分类
	Stock   int     // 商品分类ID
	Caption string  // 产品名称
	Desc    string  // 产品描述
	ImgSrc  string  // 产品图片
	Price   float32 // 价格
}
