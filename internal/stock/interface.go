package stock

type Storage interface {
	UpdateStock(productId, delta int) error
}
