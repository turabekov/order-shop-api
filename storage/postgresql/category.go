package postgresql

import (
	"app/api/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(req *models.CreateCategory) (string, error) {
	var (
		query    string
		id       = uuid.New()
		parentId string
	)

	if len(req.ParentId) > 0 {
		parentId = "parent_id" + req.ParentId + ", "
	}

	query = `
		INSERT INTO categories(
			id, 
			name,
	` + parentId + ` 
		updated_at
	)
	VALUES ($1, $2, now())`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *categoryRepo) GetByID(req *models.CategoryPrimaryKey) (*models.Category, error) {

	var (
		query    string
		category models.Category
		parentId sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			parent_id,
			created_at,
			updated_at
		FROM categories
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&parentId,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	category.ParentId = parentId.String

	return &category, nil
}

func (r *categoryRepo) GetList(req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error) {

	resp = &models.GetListCategoryResponse{}

	var (
		query    string
		filter   = " WHERE TRUE "
		offset   = " OFFSET 0"
		limit    = " LIMIT 10"
		parentId sql.NullString
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			parent_id,
			created_at,
			updated_at
		FROM categories
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
		var category models.Category
		err = rows.Scan(
			&resp.Count,
			&category.Id,
			&category.Name,
			&parentId,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		category.ParentId = parentId.String

		resp.Categories = append(resp.Categories, &category)
	}

	return resp, nil
}

func (r *categoryRepo) Update(req *models.UpdateCategory) (int64, error) {
	var (
		name     string
		parentId string
		filter   = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		categories
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}
	if len(req.ParentId) > 0 {
		parentId = " parent_id = '" + req.ParentId + "', "
	}

	query += name + parentId + " updated_at = now() " + filter

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

func (r *categoryRepo) Delete(req *models.CategoryPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM categories
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
