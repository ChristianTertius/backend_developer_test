package repository

import "database/sql"

type postgresCustomerRepository struct {
	db *sql.DB
}

func NewPostgresCustomerRepository(db *sql.DB) domain.CustomerRepository {

}
