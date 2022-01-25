package data

import (
	"context"

	"github.com/haappywater/virtualwallet/internal/biz"
	"github.com/haappywater/virtualwallet/internal/biz/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type virtualWalletRepo struct {
	data *Data
}

type VirtualWallet struct {
	gorm.Model
	UserID  uint            `json:"user_id,omitempty"`
	Balance decimal.Decimal `json:"balance,omitempty" gorm:"type:decimal(20,8)"`
}

// GetVirtualWallet 获取虚拟钱包模型
func (r *virtualWalletRepo) GetVirtualWallet(ctx context.Context, userID uint) (*types.VirtualWallet, error) {
	var w VirtualWallet
	if err := r.data.db.WithContext(ctx).Model(&VirtualWallet{}).First(&w, "user_id=?", userID).Error; err != nil {
		return nil, err
	}

	return &types.VirtualWallet{}, nil
}

// UpdateBalance 更新虚拟钱包余额
func (r *virtualWalletRepo) UpdateBalance(ctx context.Context, id uint, amount decimal.Decimal) error {
	return r.data.db.WithContext(ctx).Model(&VirtualWallet{}).Where("id=?", id).Update("balance", amount).Error
}

func NewVirtualWalletRepo(data *Data) biz.WalletRepo {
	return &virtualWalletRepo{data: data}
}
