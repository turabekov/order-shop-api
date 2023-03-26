package postgresql

import (
	"app/config"
	"app/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	db       *sql.DB
	customer storage.CustomerRepoI
	user     storage.UserRepoI
	courier  storage.CourierRepoI
	product  storage.ProductRepoI
	category storage.CategoryRepoI
	order    storage.OrderRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {
	connection := fmt.Sprintf(
		"host=%s user=%s database=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		db:       db,
		customer: NewCustomerRepo(db),
		user:     NewUserRepo(db),
		courier:  NewCourierRepo(db),
		product:  NewProductRepo(db),
		category: NewCategoryRepo(db),
		order:    NewOrderRepo(db),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Customer() storage.CustomerRepoI {
	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}

	return s.customer
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *Store) Courier() storage.CourierRepoI {
	if s.courier == nil {
		s.courier = NewCourierRepo(s.db)
	}

	return s.courier
}
func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}
