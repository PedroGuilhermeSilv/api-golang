package database

import (
	"testing"

	stdErrors "errors"

	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	customErrors "github.com/PedroGuilhermeSilv/api-golang/pkg/erros"
	"github.com/PedroGuilhermeSilv/api-golang/pkg/testutils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProductDBCreateSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)

	product, _ := entity.NewProduct("Product 1", 100)

	err := productDb.Create(product)
	assert.Nil(t, err)

	productFound := &entity.Product{}
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProductDBFindByIDSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)

	product, _ := entity.NewProduct("Product 1", 100)

	err := productDb.Create(product)
	assert.Nil(t, err)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProductDBFindByIDNotFound(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)

	productFound, err := productDb.FindByID("123")
	assert.Error(t, err)
	assert.Nil(t, productFound)
	assert.True(t, stdErrors.Is(err, customErrors.ErrorProductNotFound))
}

func TestProductDBFindAllSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)

	product1, _ := entity.NewProduct("Product 1", 100)
	product2, _ := entity.NewProduct("Product 2", 200)
	product3, _ := entity.NewProduct("Product 3", 300)

	productDb.Create(product1)
	productDb.Create(product2)
	productDb.Create(product3)

	products, err := productDb.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, product1.ID, products[0].ID)
	assert.Equal(t, product2.ID, products[1].ID)
	assert.Equal(t, product3.ID, products[2].ID)
}

func TestProductDBFindAllByPageAndLimitAndSortSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)

	product1, _ := entity.NewProduct("Product 1", 100)
	product2, _ := entity.NewProduct("Product 2", 200)
	product3, _ := entity.NewProduct("Product 3", 300)
	product4, _ := entity.NewProduct("Product 4", 400)
	product5, _ := entity.NewProduct("Product 5", 500)
	product6, _ := entity.NewProduct("Product 6", 600)

	productDb.Create(product1)
	productDb.Create(product2)
	productDb.Create(product3)
	productDb.Create(product4)
	productDb.Create(product5)
	productDb.Create(product6)

	products, err := productDb.FindAll(2, 2, "desc")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, product4.ID, products[0].ID)
	assert.Equal(t, product3.ID, products[1].ID)
}

func TestProductDBDeleteSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)

	product, _ := entity.NewProduct("Product 1", 100)

	err := productDb.Create(product)
	assert.Nil(t, err)

	err = productDb.Delete(product.ID.String())
	assert.Nil(t, err)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.True(t, stdErrors.Is(err, customErrors.ErrorProductNotFound))
	assert.Nil(t, productFound)

}

func TestProductDBDeleteNotFound(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)
	id := uuid.New()
	err := productDb.Delete(id.String())
	assert.True(t, stdErrors.Is(err, customErrors.ErrorProductNotFound))
}

func TestUpdateProductSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.Product{})
	productDb := NewProductDB(db)
	product, _ := entity.NewProduct("Product 1", 100)
	err := productDb.Create(product)
	assert.Nil(t, err)

	product.Name = "Product 2"
	product.Price = 200

	err = productDb.Update(product)
	assert.Nil(t, err)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

}
