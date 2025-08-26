package dto

type NikExistsResponse struct {
	FarmerNik       string  `json:"farmer_nik"`
	FarmerName      string  `json:"farmer_name"`
	FarmerGroupName string  `json:"farmer_group_name"`
	RetailerMid     string  `json:"retailer_mid"`
	RetailerName    string  `json:"retailer_name"`
	Urea            float64 `json:"urea"`
	ZA              float64 `json:"za"`
	SP36            float64 `json:"sp36"`
	NPK             float64 `json:"npk"`
	Organic         float64 `json:"organic"`
}

type KuotaPupuk struct {
	Pupuk string  `json:"Pupuk"`
	Kuota float64 `json:"Kuota"`
}

type KuotaResponse struct {
	Mid          string       `json:"mid"`
	FarmerName   string       `json:"farmer_name"`
	Namakios     string       `json:"nama_kios"`
	KelompokTani []string     `json:"kelompok_tani"`
	KuotaPupuk   []KuotaPupuk `json:"kuota_pupuk"`
}
