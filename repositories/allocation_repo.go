package repositories

import (
	"go-fiber-api/dto"
	"go-fiber-api/models"

	"gorm.io/gorm"
)

type AllocationRepository struct {
	db *gorm.DB
}

func NewAllocationRepository(db *gorm.DB) *AllocationRepository {
	return &AllocationRepository{db: db}
}

func (r *AllocationRepository) GetRetailerByMid(mid string) (models.Retailer, error) {
	var retailer models.Retailer
	result := r.db.Unscoped().Where("retailer_mid = ? AND is_active = ?", mid, 1).First(&retailer)
	return retailer, result.Error
}

func (r *AllocationRepository) GetRetailerByMidInquiry(mid string) ([]models.Retailer, error) {
	var retailers []models.Retailer
	err := r.db.Where("retailer_mid = ? AND is_active = 1", mid).Find(&retailers).Error
	if err != nil {
		return nil, err
	}
	return retailers, nil
}

func (r *AllocationRepository) GetRetailersByNik(nik string) ([]models.Retailer, error) {
	var retailers []models.Retailer
	result := r.db.
		Table("psp_wallets as pw").
		Select("rt.id, rt.retailer_mid, rt.name").
		Joins("JOIN retailers rt ON rt.id = pw.retailer_id").
		Where("pw.farmer_nik = ? AND rt.is_active = ?", nik, 1).
		Group("rt.id, rt.retailer_mid, rt.name").
		Find(&retailers)

	return retailers, result.Error
}

func (r *AllocationRepository) CheckNikExists(nik string) ([]dto.NikExistsResponse, error) {
	var wallets []dto.NikExistsResponse
	result := r.db.Table("psp_wallet").
		Joins("JOIN retailers ON retailers.id = psp_wallet.retailer_id").
		Select(`psp_wallet.farmer_nik AS farmer_nik,psp_wallet.farmer_name AS farmer_name,psp_wallet.farmer_group_name AS farmer_group_name,retailers.retailer_mid AS retailer_mid,retailers.name AS retailer_name,SUM(psp_wallet.urea) as urea,SUM(psp_wallet.za) as za,SUM(psp_wallet.sp36) as sp36,SUM(psp_wallet.npk) as npk,SUM(psp_wallet.organic) as organic`).
		Where("psp_wallet.farmer_nik = ? AND psp_wallet.is_active = ?", nik, 1).
		Group("psp_wallet.farmer_nik, psp_wallet.farmer_name, psp_wallet.farmer_group_name, retailers.retailer_mid, retailers.name").
		Find(&wallets)

	return wallets, result.Error
}

func (r *AllocationRepository) CheckNikExistsWallet(nik, komoditas string) (models.PspWallet, error) {
	var wallets models.PspWallet
	result := r.db.Select(`farmer_nik, retailer_id, farmer_name`).
		Where("farmer_nik = ? AND komoditas = ? AND is_active = ?", nik, komoditas, 1).
		First(&wallets)

	return wallets, result.Error
}

func (r *AllocationRepository) CheckAllocationNotFound(nik string) ([]dto.NikExistsResponse, error) {
	var wallets []dto.NikExistsResponse
	result := r.db.Table("psp_wallet").
		Joins("JOIN retailers ON retailers.id = psp_wallet.retailer_id").
		Select(`psp_wallet.farmer_nik AS farmer_nik,psp_wallet.farmer_name AS farmer_name,psp_wallet.farmer_group_name AS farmer_group_name,retailers.retailer_mid AS retailer_mid,retailers.name AS retailer_name, psp_wallet.sub_district_code,SUM(psp_wallet.urea) as urea,SUM(psp_wallet.za) as za,SUM(psp_wallet.sp36) as sp36,SUM(psp_wallet.npk) as npk,SUM(psp_wallet.organic) as organic`).
		Where("psp_wallet.farmer_nik = ? AND psp_wallet.is_active = ?", nik, 1).
		Group("psp_wallet.farmer_nik, psp_wallet.farmer_name, psp_wallet.farmer_group_name, retailers.retailer_mid, retailers.name, psp_wallet.sub_district_code").
		Find(&wallets)

	return wallets, result.Error
}

func (r *AllocationRepository) CheckAlokasiPetani(nik, komoditas string, retailerID []int) (*dto.NikExistsResponse, error) {
	var wallet dto.NikExistsResponse

	query := r.db.Table("psp_wallet AS w").
		Select(`
			w.farmer_nik,
			w.farmer_name,
			w.sub_district_code AS sub_district_code,
			STRING_AGG(DISTINCT w.farmer_group_name, ', ') AS farmer_group_name,
			STRING_AGG(DISTINCT rt.name, ', ') AS retailer_name,
			SUM(w.urea) AS urea,
			SUM(w.npk) AS npk,
			SUM(w.npk_formula) AS npk_formula,
			SUM(w.sp36) AS sp36,
			SUM(w.za) AS za,
			SUM(w.organic) AS organic,
			SUM(w.poc) AS poc
		`).
		Joins("JOIN retailers AS rt ON rt.id = w.retailer_id").
		Where("w.farmer_nik = ?", nik).
		Where("UPPER(w.komoditas) = ?", komoditas).
		Where("w.retailer_id IN ?", retailerID).
		Where("w.is_active = ?", 1).
		Group("w.farmer_nik, w.farmer_name, w.sub_district_code").
		Limit(1)

	if err := query.Scan(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *AllocationRepository) GetHargaByNama(namaPupuk string) (models.HargaPupuk, error) {
	var harga models.HargaPupuk

	result := r.db.Where("nama_pupuk = ? AND deleted_at IS NULL", namaPupuk).First(&harga)
	return harga, result.Error
}
