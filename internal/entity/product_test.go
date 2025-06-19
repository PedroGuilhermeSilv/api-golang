package entity

import (
	"testing"

	errors "github.com/PedroGuilhermeSilv/api-golang/pkg/erros"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100, product.Price)
}

func TestProductValidateSuccess(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100, product.Price)
}

func TestProductValidateError(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "name is required", err.Error())
	assert.IsType(t, errors.ErrorNameIsRequired, err)
}

func TestProductValidateErrorPrice(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "price is required", err.Error())
	assert.IsType(t, errors.ErrorPriceIsRequired, err)
}

func TestProductValidateErrorPriceInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "price is invalid", err.Error())
	assert.IsType(t, errors.ErrorPriceIsInvalid, err)
}
