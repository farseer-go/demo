package stock

// Repository 商品仓储
type Repository interface {
	// Get 获取库存
	Get(productId int64) int
	// GetAll 获取所有库存
	GetAll() map[int64]int
	// Set 设置库存
	Set(productId int64, val int) int
}
