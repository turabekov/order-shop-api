package models

type Order struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64 `json:"longtitude"`

	UserId     string `json:"user_id"`
	CustomerId string `json:"customer_id"`
	CourierId  string `json:"courier_id"`
	ProductId  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}


type OrderPrimaryKey struct {
	Id string `json:"id"`
}

// ---- added
type OrderResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64 `json:"longtitude"`

	User      User            `json:"user"`
	Customer  Customer        `json:"customer"`
	Courier   Courier         `json:"courier"`
	Product   ProductCategory `json:"product"`
	Quantity  int             `json:"quantity"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
}

type CreateOrder struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64 `json:"longtitude"`

	UserId     string `json:"user_id"`
	CustomerId string `json:"customer_id"`
	CourierId  string `json:"courier_id"`
	ProductId  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type UpdateOrder struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64 `json:"longtitude"`

	UserId     string `json:"user_id"`
	CustomerId string `json:"customer_id"`
	CourierId  string `json:"courier_id"`
	ProductId  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int              `json:"count"`
	Orders []*OrderResponse `json:"authors"`
}
