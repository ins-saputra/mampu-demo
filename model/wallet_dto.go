package model

import "github.com/shopspring/decimal"

type WithdrawRequest struct {
	UserID string          `json:"user_id"`
	Amount decimal.Decimal `json:"amount"`
}

type WithdrawResponse struct {
	Message    string          `json:"message"`
	UserID     string          `json:"user_id"`
	NewBalance decimal.Decimal `json:"new_balance"`
}

type WalletResponse struct {
	UserID  string          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}
