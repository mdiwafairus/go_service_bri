package models

import "time"

type KartanFarmerTransaction struct {
	ID                uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	SubdistrictCode   string     `gorm:"type:char(10);not null" json:"subdistrict_code"`
	DistrictCode      string     `gorm:"type:char(6);not null" json:"district_code"`
	CityCode          string     `gorm:"type:char(4);not null" json:"city_code"`
	ProvinceCode      string     `gorm:"type:char(2);not null" json:"province_code"`
	NationalID        string     `gorm:"type:char(16);not null" json:"national_id"`
	FarmerName        string     `gorm:"type:varchar(150);not null" json:"farmer_name"`
	UreaKg            float64    `gorm:"type:numeric(8,2);default:0.00" json:"urea_kg"`
	NpkKg             float64    `gorm:"type:numeric(8,2);default:0.00" json:"npk_kg"`
	Sp36Kg            float64    `gorm:"type:numeric(8,2);default:0.00" json:"sp36_kg"`
	ZaKg              float64    `gorm:"type:numeric(8,2);default:0.00" json:"za_kg"`
	NpkFormulaKg      float64    `gorm:"type:numeric(8,2);default:0.00" json:"npk_formula_kg"`
	OrganicKg         float64    `gorm:"type:numeric(8,2);default:0.00" json:"organic_kg"`
	PocKg             float64    `gorm:"type:numeric(8,2);default:0.00" json:"poc_kg"`
	ApprovedBy        *uint      `gorm:"type:int" json:"approved_by"`
	Notes             *string    `gorm:"type:varchar(200)" json:"notes"`
	IsSent            int16      `gorm:"type:smallint;default:3;not null" json:"is_sent"`
	KioskName         *string    `gorm:"type:varchar(200)" json:"kiosk_name"`
	RedeemDay         int16      `gorm:"type:smallint;not null" json:"redeem_day"`
	RedeemMonth       int16      `gorm:"type:smallint;not null" json:"redeem_month"`
	RedeemYear        int16      `gorm:"type:smallint;not null" json:"redeem_year"`
	BankTimestamp     time.Time  `gorm:"type:timestamp;not null" json:"bank_timestamp"`
	UpdatedAt         time.Time  `gorm:"type:timestamp" json:"updated_at"`
	ApprovalTimestamp *time.Time `gorm:"type:timestamp" json:"approval_timestamp"`
	SourceType        int16      `gorm:"type:smallint;default:1;not null" json:"source_type"`
	TransactionCode   string     `gorm:"type:varchar(50);not null" json:"transaction_code"`
	ReferenceCode     string     `gorm:"type:varchar(50);not null" json:"reference_code"`
	ProvinceName      *string    `gorm:"type:varchar(100)" json:"province_name"`
	DistrictName      *string    `gorm:"type:varchar(100)" json:"district_name"`
	SubdistrictName   *string    `gorm:"type:varchar(100)" json:"subdistrict_name"`
	VillageName       *string    `gorm:"type:varchar(100)" json:"village_name"`
	CreatedAt         time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	Commodity         *string    `gorm:"type:varchar(50)" json:"commodity"`
	FarmerGroupID     *uint      `gorm:"type:int" json:"farmer_group_id"`
	FarmerGroupName   *string    `gorm:"type:varchar(150)" json:"farmer_group_name"`
	ReferenceNumber   *string    `gorm:"type:varchar(100)" json:"reference_number"`
	WalletID          *uint      `gorm:"type:int" json:"wallet_id"`
}

// TableName override
func (KartanFarmerTransaction) TableName() string {
	return "kartan_farmer_transaction"
}
