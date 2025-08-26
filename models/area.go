package models

type Area struct {
	SubDistrictCode string `gorm:"column:sub_district_code;" json:"sub_district_code"`
}
