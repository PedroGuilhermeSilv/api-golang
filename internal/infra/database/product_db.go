package database

import (
	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	errors "github.com/PedroGuilhermeSilv/api-golang/pkg/erros"
	"gorm.io/gorm"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Create(product *entity.Product) error {
	return p.db.Create(product).Error
}

func (p *ProductDB) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page > 0 && limit > 0 {
		err = p.db.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.db.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *ProductDB) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.db.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, errors.ErrorProductNotFound
	}
	return &product, nil
}

func (p *ProductDB) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return errors.ErrorProductNotFound
	}
	return p.db.Save(product).Error
}

func (p *ProductDB) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return errors.ErrorProductNotFound
	}
	err = p.db.Delete(&entity.Product{}, "id = ?", id).Error
	if err != nil {
		return errors.ErrorProductDelete
	}
	return nil
}
