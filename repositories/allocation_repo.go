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

func (r *AllocationRepository) GetRetailerByMid(mid int) (models.Retailer, error) {
	var retailer models.Retailer
	result := r.db.Unscoped().Where("retailer_mid = ? AND is_active = ?", mid, 1).First(&retailer)
	return retailer, result.Error
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
