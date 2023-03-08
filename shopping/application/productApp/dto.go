package productApp

type DTO struct {
	Id       int64   // 商品ID
	Caption  string  // 商品名称
	Desc     string  // 商品描述
	ImgSrc   string  // 商品图片
	Price    float32 // 商品价格
	Stock    int     // 库存
	CateName string  // 商品名称
}
