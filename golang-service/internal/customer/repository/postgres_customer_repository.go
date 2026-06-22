package repository

import (
	"context"
	"database/sql"

	"github.com/ChristianTertius/backend_developer_test/internal/domain"
)

type postgresCustomerRepository struct {
	db *sql.DB
}

// NewPostgresCustomerRepository membuat implementasi CustomerRepository.
func NewPostgresCustomerRepository(db *sql.DB) domain.CustomerRepository {
	return &postgresCustomerRepository{db: db}
}

// Fetch mengambil seluruh customer beserta nationality dan daftar keluarganya.
func (r *postgresCustomerRepository) Fetch(ctx context.Context) ([]domain.Customer, error) {
	query := `
		SELECT c.cst_id, c.nationality_id, TRIM(c.cst_name),
		       TO_CHAR(c.cst_dob, 'YYYY-MM-DD'), c."cst_phoneNum", c.cst_email,
		       n.nationality_id, n.nationality_name, n.nationality_code
		FROM customer c
		JOIN nationality n ON n.nationality_id = c.nationality_id
		ORDER BY c.cst_id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]domain.Customer, 0)
	indexByID := make(map[int64]int)
	ids := make([]int64, 0)

	for rows.Next() {
		var c domain.Customer
		var n domain.Nationality
		if err := rows.Scan(
			&c.ID, &c.NationalityID, &c.Name, &c.DOB, &c.PhoneNum, &c.Email,
			&n.ID, &n.Name, &n.Code,
		); err != nil {
			return nil, err
		}
		c.Nationality = &n
		c.Families = make([]domain.Family, 0)
		indexByID[c.ID] = len(customers)
		customers = append(customers, c)
		ids = append(ids, c.ID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return customers, nil
	}

	// Ambil seluruh keluarga sekaligus (hindari N+1 query).
	families, err := r.fetchFamiliesByCustomerIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, f := range families {
		if idx, ok := indexByID[f.CustomerID]; ok {
			customers[idx].Families = append(customers[idx].Families, f)
		}
	}
	return customers, nil
}

func (r *postgresCustomerRepository) fetchFamiliesByCustomerIDs(ctx context.Context, ids []int64) ([]domain.Family, error) {
	query := `SELECT fl_id, cst_id, fl_relation, fl_name, fl_dob
	          FROM family_list
	          WHERE cst_id = ANY($1)
	          ORDER BY fl_id`
	rows, err := r.db.QueryContext(ctx, query, pqInt64Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.Family, 0)
	for rows.Next() {
		var f domain.Family
		if err := rows.Scan(&f.ID, &f.CustomerID, &f.Relation, &f.Name, &f.DOB); err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, rows.Err()
}

// GetByID mengambil satu customer beserta keluarganya.
func (r *postgresCustomerRepository) GetByID(ctx context.Context, id int64) (domain.Customer, error) {
	query := `
		SELECT c.cst_id, c.nationality_id, TRIM(c.cst_name),
		       TO_CHAR(c.cst_dob, 'YYYY-MM-DD'), c."cst_phoneNum", c.cst_email,
		       n.nationality_id, n.nationality_name, n.nationality_code
		FROM customer c
		JOIN nationality n ON n.nationality_id = c.nationality_id
		WHERE c.cst_id = $1`

	var c domain.Customer
	var n domain.Nationality
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.NationalityID, &c.Name, &c.DOB, &c.PhoneNum, &c.Email,
		&n.ID, &n.Name, &n.Code,
	)
	if err == sql.ErrNoRows {
		return c, domain.ErrNotFound
	}
	if err != nil {
		return c, err
	}
	c.Nationality = &n

	families, err := r.fetchFamiliesByCustomerIDs(ctx, []int64{id})
	if err != nil {
		return c, err
	}
	c.Families = families
	return c, nil
}

// Store menyimpan customer + keluarga dalam satu transaksi.
func (r *postgresCustomerRepository) Store(ctx context.Context, c *domain.Customer) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	insertCustomer := `
		INSERT INTO customer (nationality_id, cst_name, cst_dob, "cst_phoneNum", cst_email)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING cst_id`
	if err := tx.QueryRowContext(ctx, insertCustomer,
		c.NationalityID, c.Name, c.DOB, c.PhoneNum, c.Email,
	).Scan(&c.ID); err != nil {
		return err
	}

	if err := insertFamilies(ctx, tx, c.ID, c.Families); err != nil {
		return err
	}
	return tx.Commit()
}

// Update memperbarui customer dan menyetel ulang daftar keluarganya (delete + insert).
func (r *postgresCustomerRepository) Update(ctx context.Context, c *domain.Customer) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	updateCustomer := `
		UPDATE customer
		SET nationality_id = $1, cst_name = $2, cst_dob = $3,
		    "cst_phoneNum" = $4, cst_email = $5
		WHERE cst_id = $6`
	res, err := tx.ExecContext(ctx, updateCustomer,
		c.NationalityID, c.Name, c.DOB, c.PhoneNum, c.Email, c.ID,
	)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return domain.ErrNotFound
	}

	if _, err := tx.ExecContext(ctx, `DELETE FROM family_list WHERE cst_id = $1`, c.ID); err != nil {
		return err
	}
	if err := insertFamilies(ctx, tx, c.ID, c.Families); err != nil {
		return err
	}
	return tx.Commit()
}

// Delete menghapus customer (keluarga ikut terhapus via ON DELETE CASCADE).
func (r *postgresCustomerRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM customer WHERE cst_id = $1`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func insertFamilies(ctx context.Context, tx *sql.Tx, customerID int64, families []domain.Family) error {
	insertFamily := `
		INSERT INTO family_list (cst_id, fl_relation, fl_name, fl_dob)
		VALUES ($1, $2, $3, $4)
		RETURNING fl_id`
	for i := range families {
		families[i].CustomerID = customerID
		if err := tx.QueryRowContext(ctx, insertFamily,
			customerID, families[i].Relation, families[i].Name, families[i].DOB,
		).Scan(&families[i].ID); err != nil {
			return err
		}
	}
	return nil
}
