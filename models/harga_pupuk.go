package models

import (
	"database/sql"
	"time"
)

type HargaPupuk struct {
	NamaPupuk  string
	Harga      int
	ClientType string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}

func (HargaPupuk) TableName() string {
	return "harga_pupuk"
}
