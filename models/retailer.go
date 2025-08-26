package models

import "time"

type Retailer struct {
	Name            string    `json:"name"`
	PihcCode        string    `json:"pihc_code"`
	SubDistrictCode string    `json:"sub_district_code"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsActive        int       `json:"is_active"`
	PihcCodeAgent   string    `json:"pihc_code_agent"`
	RetailerMid     string    `json:"retailer_mid"`
	KodeAgen        int       `json:"kode_agen"`
	Year            int       `json:"year"`

	Area Area `gorm:"foreignKey:SubDistrictCode;references:SubDistrictCode" json:"-"`
}
