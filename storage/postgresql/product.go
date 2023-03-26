package postgresql

import (
	"app/api/models"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(req *models.CreateProduct) (string, error) {
	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO products(
			id, 
			name, 
			price,
			category_id,
			updated_at
		)
		VALUES ($1, $2, $3, $4, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Price,
		req.CategoryId,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *productRepo) GetByID(req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query   string
		product models.Product
	)

	query = `
		SELECT
			id,
			name,
			price,
			category_id,
			created_at,
			updated_at
		FROM products
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.CategoryId,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) GetList(req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) {

	resp = &models.GetListProductResponse{}

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
			price,
			category_id,
			created_at,
			updated_at
		FROM products
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
		var product models.Product
		err = rows.Scan(
			&resp.Count,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.CategoryId,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product)
	}

	return resp, nil
}

func (r *productRepo) Update(req *models.UpdateProduct) (int64, error) {
	var (
		name     string
		price    string
		category string

		filter = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		products
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}
	if req.Price > 0 {
		price = " price = '" + strconv.Itoa(int(req.Price)) + "', "
	}

	if len(req.CategoryId) > 0 {
		category = " category = '" + req.CategoryId + "', "
	}

	query += name + price + category + " updated_at = now() " + filter

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

func (r *productRepo) Delete(req *models.ProductPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM products
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
