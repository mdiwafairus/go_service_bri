package repositories

import (
	"go-fiber-api/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
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
