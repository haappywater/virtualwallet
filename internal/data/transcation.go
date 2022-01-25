package data

import (
	"context"

	"github.com/haappywater/virtualwallet/internal/biz"
	"github.com/haappywater/virtualwallet/internal/biz/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transcation struct {
	gorm.Model
	UserID          uint            `json:"user_id,omitempty"`
	FromWalletID    uint            `json:"from_wallet_id,omitempty"`
	ToWalletID      uint            `json:"to_wallet_id,omitempty"`
	Amount          decimal.Decimal `json:"amount,omitempty"`
	TranscationType uint            `gorm:"type:tinyint(4)"`
}

type transcationRepo struct {
	data *Data
}

func (r *transcationRepo) SaveTranscation(ctx context.Context, ts *types.Transcation) error {
	transcation := &Transcation{
		UserID:          ts.UserID(),
		FromWalletID:    ts.FromWalletID(),
		ToWalletID:      ts.ToWalletID(),
		Amount:          ts.Amount(),
		TranscationType: uint(ts.TranscationType()),
	}
	return r.data.db.WithContext(ctx).Create(transcation).Error
}

func NewTranscationRepo(data *Data) biz.TranscationRepo {
	return &transcationRepo{data: data}
}
