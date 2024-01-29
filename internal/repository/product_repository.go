package repository

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (pr *ProductRepository) CreateProduct() {

}

func (pr *ProductRepository) GetProducts() {

}
