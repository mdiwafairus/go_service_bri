package models

import (
	"database/sql"
	"time"
)

type Kartan struct {
	ID              int64          `gorm:"primaryKey;column:id" json:"id"`
	SubDistrictCode string         `gorm:"column:sub_district_code;size:10;not null" json:"sub_district_code"`
	DistrictCode    string         `gorm:"column:district_code;size:6;not null" json:"district_code"`
	CityCode        string         `gorm:"column:city_code;size:4;not null" json:"city_code"`
	ProvinceCode    string         `gorm:"column:province_code;size:2;not null" json:"province_code"`
	NIK             string         `gorm:"column:nik;size:16;not null" json:"nik"`
	FarmerName      string         `gorm:"column:farmer_name;size:150;not null" json:"farmer_name"`
	Urea            float64        `gorm:"column:urea;type:numeric(10,2);default:0.00" json:"urea"`
	Npk             float64        `gorm:"column:npk;type:numeric(10,2);default:0.00" json:"npk"`
	Sp36            float64        `gorm:"column:sp36;type:numeric(10,2);default:0.00" json:"sp36"`
	Za              float64        `gorm:"column:za;type:numeric(10,2);default:0.00" json:"za"`
	NpkFormula      float64        `gorm:"column:npk_formula;type:numeric(10,2);default:0.00" json:"npk_formula"`
	Organic         float64        `gorm:"column:organic;type:numeric(10,2);default:0.00" json:"organic"`
	Poc             float64        `gorm:"column:poc;type:numeric(10,2);default:0.00" json:"poc"`
	CreatedBy       sql.NullInt64  `gorm:"column:created_by" json:"created_by"`
	Notes           sql.NullString `gorm:"column:notes;size:200" json:"notes"`
	Status          int16          `gorm:"column:status;default:1" json:"status"`
	ShopName        sql.NullString `gorm:"column:shop_name;size:255" json:"shop_name"`
	RedeemDate      time.Time      `gorm:"column:redeem_date;type:date;not null" json:"redeem_date"`
	BankDate        time.Time      `gorm:"column:bank_date;type:timestamp;not null" json:"bank_date"`
	ApprovalDate    sql.NullTime   `gorm:"column:approval_date;type:timestamp" json:"approval_date"`
	Source          int16          `gorm:"column:source;default:1" json:"source"`
	TransactionRef  string         `gorm:"column:transaction_ref;size:255;not null" json:"transaction_ref"`
	ReferenceCode   string         `gorm:"column:reference_code;size:255;not null" json:"reference_code"`
	ProvinceName    sql.NullString `gorm:"column:province_name;size:255" json:"province_name"`
	CityName        sql.NullString `gorm:"column:city_name;size:255" json:"city_name"`
	DistrictName    sql.NullString `gorm:"column:district_name;size:255" json:"district_name"`
	VillageName     sql.NullString `gorm:"column:village_name;size:255" json:"village_name"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       sql.NullTime   `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Commodity       sql.NullString `gorm:"column:commodity;size:50" json:"commodity"`
	GroupID         sql.NullInt64  `gorm:"column:group_id" json:"group_id"`
	GroupName       sql.NullString `gorm:"column:group_name;size:255" json:"group_name"`
	ReferenceNumber sql.NullString `gorm:"column:reference_number;size:200" json:"reference_number"`
	WalletID        sql.NullInt64  `gorm:"column:wallet_id" json:"wallet_id"`
}

func (Kartan) TableName() string {
	return "kartans"
}
