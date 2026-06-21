package repository

import (
	"context"
	"database/sql"
	"github.com/ChristianTertius/backend_developer_test/internal/domain"
)

type postgresNationalityRepository struct {
	db *sql.DB
}

func NewPostgresNationalityRepository(db *sql.DB) domain.NationalityRepository {
	return &postgresNationalityRepository{db: db}
}

func (r *postgresNationalityRepository) Fetch(ctx context.Context) ([]domain.Nationality, error) {
	query := `SELECT nationality_id, nationality_name, nationality_code
	          FROM nationality
	          ORDER BY nationality_name`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]domain.Nationality, 0)
	for rows.Next() {
		var n domain.Nationality
		if err := rows.Scan(&n.ID, &n.Name, &n.Code); err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, rows.Err()
}

func (r *postgresNationalityRepository) GetByID(ctx context.Context, id int64) (domain.Nationality, error) {
	query := `SELECT nationality_id, nationality_name, nationality_code
	          FROM nationality WHERE nationality_id = $1`
	var n domain.Nationality
	err := r.db.QueryRowContext(ctx, query, id).Scan(&n.ID, &n.Name, &n.Code)
	if err == sql.ErrNoRows {
		return n, domain.ErrNotFound
	}
	return n, err
}
