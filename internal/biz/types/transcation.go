package types

import (
	"github.com/shopspring/decimal"
)

type Transcation struct {
	userID          uint
	fromWalletID    uint
	toWalletID      uint
	amount          decimal.Decimal
	transcationType TranscationType
}

func (ts *Transcation) UserID() uint {
	return ts.userID
}

func (ts *Transcation) FromWalletID() uint {
	return ts.fromWalletID
}

func (ts *Transcation) ToWalletID() uint {
	return ts.toWalletID
}

func (ts *Transcation) Amount() decimal.Decimal {
	return ts.amount
}

func (ts *Transcation) TranscationType() TranscationType {
	return ts.transcationType
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
