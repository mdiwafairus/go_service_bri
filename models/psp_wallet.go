package models

import "time"

type PspWallet struct {
	Province        string    `json:"province"`
	ProvinceCode    string    `json:"province_code"`
	City            string    `json:"city"`
	CityCode        string    `json:"city_code"`
	District        string    `json:"district"`
	DistrictCode    string    `json:"district_code"`
	SubDistrict     string    `json:"sub_district"`
	SubDistrictCode string    `json:"sub_district_code"`
	RetailerID      int       `json:"retailer_id"`
	PIHCCode        string    `json:"pihc_code"`
	RetailerName    string    `json:"retailer_name"`
	FarmerGroupId   int       `json:"farmer_group_id"`
	FarmerGroupName string    `json:"farmer_group_name"`
	FarmerNIK       string    `json:"farmer_nik"`
	FarmerName      string    `json:"farmer_name"`
	Komoditas       string    `json:"komoditas"`
	PlantingArea    float64   `json:"planting_area"`
	UreaAlo         float64   `json:"urea_alo"`
	Urea            float64   `json:"urea"`
	NpkAlo          float64   `json:"npk_alo"`
	Npk             float64   `json:"npk"`
	NpkFormulaAlo   float64   `json:"npk_formula_alo"`
	NpkFormula      float64   `json:"npk_formula"`
	Sp36Alo         float64   `json:"sp36_alo"`
	Sp36            float64   `json:"sp36"`
	ZaAlo           float64   `json:"za_alo"`
	Za              float64   `json:"za"`
	OrganicAlo      float64   `json:"organic_alo"`
	Organic         float64   `json:"organic"`
	PocAlo          float64   `json:"poc_alo"`
	Poc             float64   `json:"poc"`
	Year            int       `json:"year"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsActive        bool      `json:"is_active"`
	IsBlock         bool      `json:"is_block"`
	IdWallet        int       `json:"id_wallet"`
}

func (PspWallet) TableName() string {
	return "psp_wallet"
}
