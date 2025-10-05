package repositories

import (
	"go-fiber-api/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func (r *TransactionRepository) GetDB() *gorm.DB {
	return r.db
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetRetailerByMidInquiry(mid string) ([]models.Retailer, error) {
	var retailers []models.Retailer
	err := r.db.Where("retailer_mid = ? AND is_active = 1", mid).Find(&retailers).Error
	if err != nil {
		return nil, err
	}
	return retailers, nil
}

func (r *TransactionRepository) CheckAlokasiPetaniTransaction(nik, komoditas string, retailerID []int) ([]models.PspWallet, error) {
	var wallet []models.PspWallet
	err := r.db.Where("farmer_nik = ? AND komoditas = ? AND retailer_id IN ? AND is_active = 1", nik, komoditas, retailerID).
		Order("urea + npk + npk_formula + sp36 + za + organic + poc DESC").
		Find(&wallet).Error

	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *TransactionRepository) CheckDuplicateTransaction(refnum int) (*models.KartanFarmerTransaction, error) {
	var kartanfarmer models.KartanFarmerTransaction
	err := r.db.Where("refnum = ?", refnum).First(&kartanfarmer).Error
	if err != nil {
		return nil, err
	}
	return &kartanfarmer, nil
}

func (r *TransactionRepository) InsertTransaction(gormTx *gorm.DB, tx *models.KartanFarmerTransaction) error {
	return gormTx.Create(tx).Error
}
