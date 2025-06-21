package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PedroGuilhermeSilv/api-golang/internal/dto"
	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	"github.com/PedroGuilhermeSilv/api-golang/internal/infra/database"
	entityPkg "github.com/PedroGuilhermeSilv/api-golang/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductRepositoryInterface
}

func NewProductHandler(db database.ProductRepositoryInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.ProductCreateInput true "Product to create"
// @Success 201 {object} dto.ProductCreateOutput
// @Failure 400 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.ProductCreateInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body dto.ProductCreateInput  true "Product to update"
// @Success 200 {object} entity.Product
// @Failure 400 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Update(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204
// @Failure 400 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	err := h.ProductDB.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// @Summary List products
// @Description List products
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Success 200 {object} []entity.Product
// @Failure 400 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
