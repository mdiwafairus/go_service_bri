package models

import "time"

type User struct {
	ID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Password     string `json:"-"`
	ProvinceCode string `json:"province_code"`
	CityCode     string `json:"city_code"`
	DistrictCode string `json:"district_code"`
	RoleID       int    `json:"role_id"`
	RoleType     int    `json:"role_type"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
