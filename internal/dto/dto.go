package dto

type ProductCreateInput struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type ProductCreateOutput struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}