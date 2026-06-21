package usecase

import (
	"context"
	"time"

	"github.com/ChristianTertius/backend_developer_test/internal/domain"
)

type nationalityUsecase struct {
	repo    domain.NationalityRepository
	timeout time.Duration
}

func NewNationalityUseCase(repo domain.NationalityRepository, timeout time.Duration) domain.NationalityUsecase {
	return &nationalityUsecase{repo: repo, timeout: timeout}
}

func (u *nationalityUsecase) Fetch(ctx context.Context) ([]domain.Nationality, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.Fetch(ctx)
}
