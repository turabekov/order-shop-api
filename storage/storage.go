package storage

import "app/api/models"

type StorageI interface {
	CloseDB()
	Customer() CustomerRepoI
	User() UserRepoI
	Courier() CourierRepoI
	Product() ProductRepoI
	Category() CategoryRepoI
	Order() OrderRepoI
}

type CustomerRepoI interface {
	Create(*models.CreateCustomer) (string, error)
	GetByID(*models.CustomerPrimaryKey) (*models.Customer, error)
	GetList(*models.GetListCustomerRequest) (*models.GetListCustomerResponse, error)
	Update(req *models.UpdateCustomer) (int64, error)
	Delete(req *models.CustomerPrimaryKey) (int64, error)
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	GetByID(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(req *models.UpdateUser) (int64, error)
	Delete(req *models.UserPrimaryKey) (int64, error)
}

type CourierRepoI interface {
	Create(*models.CreateCourier) (string, error)
	GetByID(*models.CourierPrimaryKey) (*models.Courier, error)
	GetList(*models.GetListCourierRequest) (*models.GetListCourierResponse, error)
	Update(req *models.UpdateCourier) (int64, error)
	Delete(req *models.CourierPrimaryKey) (int64, error)
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetByID(*models.ProductPrimaryKey) (*models.Product, error)
	GetList(*models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(req *models.UpdateProduct) (int64, error)
	Delete(req *models.ProductPrimaryKey) (int64, error)
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (*models.Category, error)
	GetList(*models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(req *models.UpdateCategory) (int64, error)
	Delete(req *models.CategoryPrimaryKey) (int64, error)
}

type OrderRepoI interface {
	Create(*models.CreateOrder) (string, error)
	GetByID(*models.OrderPrimaryKey) (*models.OrderResponse, error)
	GetList(*models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(req *models.UpdateOrder) (int64, error)
	UpdatePatch(req *models.PatchRequest) (int64, error)
	Delete(req *models.OrderPrimaryKey) (int64, error)
}
