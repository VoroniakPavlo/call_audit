package model

import (
	"time"
)

type CallQuestionnaireRule struct {
	Id            int       `db:"id"`
	DomainId      int       `db:"domain_id"`
	CreatedAt     time.Time `db:"created_at"`
	CreatedBy     int64     `db:"created_by"`
	UpdatedAt     time.Time `db:"updated_at"`
	UpdatedBy     int64     `db:"updated_by"`
	Last          time.Time `db:"last"`
	Enabled       bool      `db:"enabled"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	CallDirection string    `db:"call_direction"`
	Active        int64     `db:"active"`
	Limit         int64     `db:"limit"`
}
