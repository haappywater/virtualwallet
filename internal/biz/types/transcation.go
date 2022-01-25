package types

import (
	"github.com/shopspring/decimal"
)

type Transcation struct {
	userID          uint
	fromWalletID    uint
	amount          decimal.Decimal
	transcationType TranscationType
}

func NewTranscation(userID, fromWalletID uint, amount decimal.Decimal, transcationType TranscationType) *Transcation {
	return &Transcation{
		userID:          userID,
		fromWalletID:    fromWalletID,
		amount:          amount,
		transcationType: transcationType,
	}
}

// TranscationType 交易类型
type TranscationType uint

const (
	// TranscationTypeUnKnown 未知交易类型
	TranscationTypeUnKnown = iota

	// TranscationTypeDebit 付款
	TranscationTypeDebit

	// TranscationTypeDebit 收款
	TranscationTypeCredit
)
