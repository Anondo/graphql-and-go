package repos

import (
	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/models"
	"github.com/jinzhu/gorm"
)

// ProductRepo ...
type ProductRepo struct {
	db *conn.DB
}

// NewProductRepo ...
func NewProductRepo(db *conn.DB) Product {
	return &ProductRepo{
		db: db,
	}
}

// GetProducts ...
func (p *ProductRepo) GetProducts() ([]models.Product, error) {

	products := []models.Product{}

	if err := p.db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

// GetProduct ...
func (p *ProductRepo) GetProduct(productID int) (*models.Product, error) {
	product := models.Product{}

	if err := p.db.Where("id = ?", productID).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil

}

// GetProductsIn ...
func (p *ProductRepo) GetProductsIn(productIDs ...int) ([]models.Product, error) {
	products := []models.Product{}

	if err := p.db.Where("id IN (?)", productIDs).Find(&products).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return products, nil
		}
		return products, err
	}

	return products, nil
}
