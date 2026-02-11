package service

import (
	"errors"
	"mampu-demo/model"
	"mampu-demo/repository"

	"github.com/shopspring/decimal"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) GetBalance(userID string) (model.WalletResponse, error) {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		return model.WalletResponse{}, errors.New("Data tidak ditemukan")
	}

	response := model.WalletResponse{
		UserID:  wallet.UserID,
		Balance: wallet.Balance,
	}

	return response, nil

}

func (s *WalletService) Withdraw(userID string, amount decimal.Decimal) (model.WithdrawResponse, error) {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		return model.WithdrawResponse{}, errors.New("wallet tidak ditemukan")
	}

	if wallet.Balance.LessThan(amount) {
		return model.WithdrawResponse{}, errors.New("Saldo anda tidak cukup")
	}

	newBalance := wallet.Balance.Sub(amount)
	err = s.repo.UpdateBalance(userID, newBalance)
	if err != nil {
		return model.WithdrawResponse{}, err
	}

	return model.WithdrawResponse{
		Message:    "Tarik tunai berhasil",
		UserID:     userID,
		NewBalance: newBalance,
	}, nil
}
