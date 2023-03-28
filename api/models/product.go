package models

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type ProductCategory struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Price     float64  `json:"price"`
	Category  Category `json:"category"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

type UpdateProduct struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

type GetListProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"authors"`
}
