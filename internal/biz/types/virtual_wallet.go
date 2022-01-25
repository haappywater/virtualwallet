package types

import (
	"errors"

	"github.com/shopspring/decimal"
)

// VirtualWallet Domain领域模型（充血模型）
type VirtualWallet struct {
	id      uint
	balance decimal.Decimal
}

// ID 钱包ID
func (vw *VirtualWallet) ID() uint {
	return vw.id
}

// Balance 钱包金额
func (vw *VirtualWallet) Balance() decimal.Decimal {
	return vw.balance
}

// Debit 付款
func (vw *VirtualWallet) Debit(amount decimal.Decimal) error {
	if vw.balance.LessThan(amount) {
		return errors.New("insufficient balance")
	}
	vw.balance = vw.balance.Sub(amount)
	return nil
}

// Credit 收款
func (vw *VirtualWallet) Credit(amount decimal.Decimal) error {
	if amount.Sign() < 1 {
		return errors.New("invalid amount")
	}
	vw.balance = vw.balance.Add(amount)
	return nil
}
