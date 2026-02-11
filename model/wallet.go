package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID       uint            `gorm:"primaryKey" json:"id"`
	UserID   string          `gorm:"uniqueIndex; not null; size:191" json:"user_id"`
	Balance  decimal.Decimal `gorm:"decimal(20,0);default:0" json:"balance"`
	CreateAt time.Time       `json:"create_at"`
	UpdateAt time.Time       `json:"update_at"`
}
