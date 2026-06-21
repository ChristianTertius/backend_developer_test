package domain

import "context"

type Nationality struct {
	ID   int64  `json:"nationality_id"`
	Name string `json:"nationality_name"`
	Code string `json:"nationality_code"`
}

type NationalityRepository interface {
	Fetch(ctx context.Context) ([]Nationality, error)
	GetByID(ctx context.Context, id int64) (Nationality, error)
}

type NationalityUsecase interface {
	Fetch(ctx context.Context) ([]Nationality, error)
}
