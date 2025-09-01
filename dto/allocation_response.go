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

type KiosTidakSesuaiResponse struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Suggest []SuggestKios `json:"suggest,omitempty"`
}

type SuggestKios struct {
	Mid  string `json:"mid"`
	Name string `json:"retailer_name"`
}

type InquiryResponse struct {
	NamaPetani    string `json:"nama_petani"`
	NamaKios      string `json:"nama_kios"`
	KelompokTani  string `json:"kelompok_tani"`
	NamaPupuk     string `json:"nama_pupuk"`
	NamaKomoditas string `json:"nama_komoditas"`
	KgBeli        string `json:"kg_beli"`
	Harga         string `json:"harga"`
	KuotaSisa     string `json:"kuota_sisa"`
	KodeDesa      string `json:"kode_desa"`
	TotalBeli     string `json:"total_beli"`
}
