package postgresql

import (
	"app/api/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type courierRepo struct {
	db *sql.DB
}

func NewCourierRepo(db *sql.DB) *courierRepo {
	return &courierRepo{
		db: db,
	}
}

func (r *courierRepo) Create(req *models.CreateCourier) (string, error) {
	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO couriers(
			id, 
			name, 
			phone_number,
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.PhoneNumber,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *courierRepo) GetByID(req *models.CourierPrimaryKey) (*models.Courier, error) {

	var (
		query   string
		courier models.Courier
	)

	query = `
		SELECT
			id,
			name,
			phone_number,
			created_at,
			updated_at
		FROM couriers
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&courier.Id,
		&courier.Name,
		&courier.PhoneNumber,
		&courier.CreatedAt,
		&courier.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &courier, nil
}

func (r *courierRepo) GetList(req *models.GetListCourierRequest) (resp *models.GetListCourierResponse, err error) {

	resp = &models.GetListCourierResponse{}

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
			phone_number,
			created_at,
			updated_at
		FROM couriers
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
		var courier models.Courier
		err = rows.Scan(
			&resp.Count,
			&courier.Id,
			&courier.Name,
			&courier.PhoneNumber,
			&courier.CreatedAt,
			&courier.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Couriers = append(resp.Couriers, &courier)
	}

	return resp, nil
}

func (r *courierRepo) Update(req *models.UpdateCourier) (int64, error) {
	var (
		name  string
		phone string

		filter = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		couriers
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}
	if len(req.PhoneNumber) > 0 {
		phone = " phone_number = '" + req.PhoneNumber + "', "
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

func (r *courierRepo) Delete(req *models.CourierPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM couriers
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
