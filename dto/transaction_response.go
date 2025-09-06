package dto

type TransactionResponse struct {
	Mid              string `json:"mid"`
	Nik              string `json:"nik"`
	NamaPupuk        string `json:"nama_pupuk"`
	NamaKomoditas    string `json:"nama_komoditas"`
	KgBeli           int    `json:"kg_beli"`
	TotalRupiah      int    `json:"total_rupiah"`
	RefNum           int    `json:"ref_num"`
	TrxId            int    `json:"trx_id"`
	ClientId         int    `json:"client_id"`
	TanggalTransaksi int    `json:"tanggal_transaksi"`
}
