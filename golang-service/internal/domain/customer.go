package domain

import (
	"errors"
	"strings"
)

var (
	ErrNotFound   = errors.New("Data not found")
	ErrBadRequest = errors.New("invalid request")
)

type Customer struct {
	ID            int64        `json:"cst_id"`
	NationalityID int64        `json:"nationality_id"`
	Name          string       `json:"cst_name"`
	DOB           string       `json:"cst_dob"`
	PhoneNum      string       `json:"cst_phoneNum"`
	Email         string       `json:"cst_email"`
	Nationality   *Nationality `json:"nationality,omitempty"`
	Families      []Family     `json:"families"`
}

func (c *Customer) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return errors.New("cst_name is required")
	}

	if c.NationalityID == 0 {
		return errors.New("nationality_id is required")
	}

	if strings.TrimSpace(c.DOB) == "" {
		return errors.New("cst_dob is required")
	}

	if strings.TrimSpace(c.Email) == "" {
		return errors.New("cst_email is required")
	}

	for i, f := range c.Families {
		if strings.TrimSpace(f.Name) == "" {
			return errors.New("fl_name keluarga wajib diisi")
		}
		if strings.TrimSpace(f.Relation) == "" {
			return errors.New("fl_relation keluarga wajib diisi")
		}
		_ = i
	}
	return nil
}
