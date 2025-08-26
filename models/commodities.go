package models

import "time"

type Commodities struct {
	Name        string    `json:"name"`
	SubSector   string    `json:"sub_sector"`
	SubSectorID string    `json:"subsector_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Commodities) TableName() string {
	return "commodities"
}
