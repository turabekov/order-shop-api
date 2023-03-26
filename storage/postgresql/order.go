package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(req *models.CreateOrder) (string, error) {
	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO orders(
			id, 
			name, 
			price,
			phone_number,
			latitude,
			longtitude,
    		user_id,
    		customer_id,
    		courier_id,
    		product_id,
    		quantity,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Price,
		req.PhoneNumber,
		req.Latitude,
		req.Longtitude,
		req.UserId,
		req.CustomerId,
		req.CourierId,
		req.ProductId,
		req.Quantity,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *orderRepo) GetByID(req *models.OrderPrimaryKey) (*models.Order, error) {

	var (
		query     string
		order     models.Order
		courierId sql.NullString
	)

	query = `
		SELECT
			id, 
			name, 
			price,
			phone_number,
			latitude,
			longtitude,
			user_id,
			customer_id,
			courier_id,
			product_id,
			quantity,
			created_at,
			updated_at
		FROM orders
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&order.Id,
		&order.Name,
		&order.Price,
		&order.PhoneNumber,
		&order.Latitude,
		&order.Longtitude,
		&order.UserId,
		&order.CustomerId,
		&courierId,
		&order.ProductId,
		&order.Quantity,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	order.CourierId = courierId.String

	return &order, nil
}

func (r *orderRepo) GetList(req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {

	resp = &models.GetListOrderResponse{}

	var (
		query     string
		filter    = " WHERE TRUE "
		offset    = " OFFSET 0"
		limit     = " LIMIT 10"
		courierId sql.NullString
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id, 
			name, 
			price,
			phone_number,
			latitude,
			longtitude,
			user_id,
			customer_id,
			courier_id,
			product_id,
			quantity,
			created_at,
			updated_at
		FROM orders
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&resp.Count,
			&order.Id,
			&order.Name,
			&order.Price,
			&order.PhoneNumber,
			&order.Latitude,
			&order.Longtitude,
			&order.UserId,
			&order.CustomerId,
			&courierId,
			&order.ProductId,
			&order.Quantity,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		order.CourierId = courierId.String

		resp.Orders = append(resp.Orders, &order)
	}

	return resp, nil
}

// =PUT============================================================
func (r *orderRepo) Update(req *models.UpdateOrder) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			name = :name,
			price = :price,
			phone_number = :phone_number,
			latitude = :latitude,
			longtitude = :longtitude, 
			user_id = :user_id,
			customer_id = :customer_id,
			courier_id = :courier_id,
			product_id = :product_id,
			quantity = :quantity,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"price":        req.Price,
		"phone_number": req.PhoneNumber,
		"latitude":     req.Latitude,
		"longtitude":   req.Longtitude,
		"user_id":      req.UserId,
		"customer_id":  req.CustomerId,
		"courier_id":   req.CourierId,
		"product_id":   req.ProductId,
		"quantity":     req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// =PATCH============================================================
func (r *orderRepo) UpdatePatch(req *models.PatchRequest) (int64, error) {
	var (
		query string
		set   string
	)

	if len(req.Fields) == 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE
			orders
		SET
	` + set + `	updated_at = now()
		WHERE id = :id	
	`

	req.Fields["id"] = req.ID

	fmt.Println(req.Fields)
	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *orderRepo) Delete(req *models.OrderPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM orders
		WHERE id = $1
	`

	result, err := r.db.Exec(query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
