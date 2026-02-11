package repository

import (
	"mampu-demo/model"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) FindByUserID(userID string) (model.Wallet, error) {
	var wallet model.Wallet
	err := r.db.Where("user_id=?", userID).First(&wallet).Error
	return wallet, err
}

func (r *WalletRepository) UpdateBalance(userID string, newBalance interface{}) error {
	return r.db.Model(&model.Wallet{}).Where("user_id = ?", userID).Update("balance", newBalance).Error
}
