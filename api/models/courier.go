package models

type Courier struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type CourierPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCourier struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateCourier struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type GetListCourierRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCourierResponse struct {
	Count    int        `json:"count"`
	Couriers []*Courier `json:"authors"`
}
