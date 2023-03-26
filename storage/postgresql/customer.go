package postgresql

import (
	"database/sql"
	"fmt"

	"app/api/models"

	"github.com/google/uuid"
)

type customerRepo struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) *customerRepo {
	return &customerRepo{
		db: db,
	}
}

func (r *customerRepo) Create(req *models.CreateCustomer) (string, error) {
	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO customers(
			id, 
			name, 
			phone,
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Phone,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *customerRepo) GetByID(req *models.CustomerPrimaryKey) (*models.Customer, error) {

	var (
		query    string
		customer models.Customer
	)

	query = `
		SELECT
			id,
			name,
			phone,
			created_at,
			updated_at
		FROM customers
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepo) GetList(req *models.GetListCustomerRequest) (resp *models.GetListCustomerResponse, err error) {

	resp = &models.GetListCustomerResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			phone,
			created_at,
			updated_at
		FROM customers
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
		var customer models.Customer
		err = rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.Name,
			&customer.Phone,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Customers = append(resp.Customers, &customer)
	}

	return resp, nil
}

func (r *customerRepo) Update(req *models.UpdateCustomer) (int64, error) {
	var (
		name  string
		phone string

		filter = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		customers
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}
	if len(req.Phone) > 0 {
		phone = " phone = '" + req.Phone + "', "
	}

	query += name + phone + " updated_at = now() " + filter

	result, err := r.db.Exec(query)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *customerRepo) Delete(req *models.CustomerPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM customers
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
