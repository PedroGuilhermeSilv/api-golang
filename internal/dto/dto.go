package dto

type ProductCreateInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductCreateOutput struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UserCreateInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginOutput struct {
	AccessToken string `json:"access_token"`
}

type ProductUpdateInput struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductUpdateOutput struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}