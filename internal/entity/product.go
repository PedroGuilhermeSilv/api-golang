package entity

import (
	"time"

	"github.com/PedroGuilhermeSilv/api-golang/pkg/entity"
	errors "github.com/PedroGuilhermeSilv/api-golang/pkg/erros"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errors.ErrorIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errors.ErrorIDIsInvalid
	}
	if p.Name == "" {
		return errors.ErrorNameIsRequired
	}
	if p.Price == 0 {
		return errors.ErrorPriceIsRequired
	}
	if p.Price < 0 {
		return errors.ErrorPriceIsInvalid
	}
	return nil
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}
