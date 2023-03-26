package models

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type CustomerPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCustomer struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UpdateCustomer struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type GetListCustomerRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCustomerResponse struct {
	Count     int         `json:"count"`
	Customers []*Customer `json:"authors"`
}
