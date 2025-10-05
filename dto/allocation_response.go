package dto

type NikExistsResponse struct {
	FarmerNik       string  `json:"farmer_nik"`
	FarmerName      string  `json:"farmer_name"`
	FarmerGroupName string  `json:"farmer_group_name"`
	RetailerMid     string  `json:"retailer_mid"`
	RetailerName    string  `json:"retailer_name"`
	SubDistrictCode string  `json:"sub_district_code"`
	Urea            float64 `json:"urea"`
	ZA              float64 `json:"za"`
	SP36            float64 `json:"sp36"`
	NPK             float64 `json:"npk"`
	Organic         float64 `json:"organic"`
	NpkFormula      float64 `json:"npk_formula"`
	Poc             float64 `json:"poc"`
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
	NamaPetani      string `json:"nama_petani"`
	NamaKios        string `json:"nama_kios"`
	KelompokTani    string `json:"kelompok_tani"`
	NamaPupuk       string `json:"nama_pupuk"`
	NamaKomoditas   string `json:"nama_komoditas"`
	SubDistrictCode string `json:"sub_district_code"`
	KgBeli          string `json:"kg_beli"`
	Harga           string `json:"harga"`
	KuotaSisa       string `json:"kuota_sisa"`
	KodeDesa        string `json:"kode_desa"`
	TotalBeli       string `json:"total_beli"`
}

func (n *NikExistsResponse) GetUrea() int       { return int(n.Urea) }
func (n *NikExistsResponse) GetZa() int         { return int(n.ZA) }
func (n *NikExistsResponse) GetSp36() int       { return int(n.SP36) }
func (n *NikExistsResponse) GetNpk() int        { return int(n.NPK) }
func (n *NikExistsResponse) GetOrganic() int    { return int(n.Organic) }
func (n *NikExistsResponse) GetNpkFormula() int { return int(n.NpkFormula) }
func (n *NikExistsResponse) GetPoc() int        { return int(n.Poc) }

func (n *NikExistsResponse) ToMap() map[string]int {
	return map[string]int{
		"UREA":        int(n.Urea),
		"ZA":          int(n.ZA),
		"SP36":        int(n.SP36),
		"NPK":         int(n.NPK),
		"ORGANIC":     int(n.Organic),
		"NPK_FORMULA": int(n.NpkFormula),
		"POC":         int(n.Poc),
	}
}
