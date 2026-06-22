package usecase

import (
	"context"
	"time"

	"github.com/ChristianTertius/backend_developer_test/internal/domain"
)

type customerUsecase struct {
	repo    domain.CustomerRepository
	timeout time.Duration
}

func NewCustomerUsecase(repo domain.CustomerRepository, timeout time.Duration) domain.CustomerUsecase {
	return &customerUsecase{repo: repo, timeout: timeout}
}

func (u *customerUsecase) Fetch(ctx context.Context) ([]domain.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.Fetch(ctx)
}

func (u *customerUsecase) GetByID(ctx context.Context, id int64) (domain.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.GetByID(ctx, id)
}

func (u *customerUsecase) Store(ctx context.Context, c *domain.Customer) error {
	if err := c.Validate(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.Store(ctx, c)
}

func (u *customerUsecase) Update(ctx context.Context, c *domain.Customer) error {
	if err := c.Validate(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	if _, err := u.repo.GetByID(ctx, c.ID); err != nil {
		return err
	}

	return u.repo.Update(ctx, c)
}

func (u *customerUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	if _, err := u.repo.GetByID(ctx, id); err != nil {
		return err
	}

	return u.repo.Delete(ctx, id)
}
