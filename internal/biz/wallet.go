package biz

import (
	"context"

	"github.com/haappywater/virtualwallet/internal/biz/types"
	"github.com/shopspring/decimal"
)

type WalletUseCase struct {
	walletRepo      WalletRepo
	transcationRepo TranscationRepo
}

func NewWalletUseCase(walletRepo WalletRepo, transcationRepo TranscationRepo) *WalletUseCase {
	return &WalletUseCase{
		walletRepo:      walletRepo,
		transcationRepo: transcationRepo,
	}
}

type WalletRepo interface {
	// GetVirtualWallet 获取虚拟钱包模型
	GetVirtualWallet(ctx context.Context, userID uint) (*types.VirtualWallet, error)

	// UpdateBalance 更新虚拟钱包余额
	UpdateBalance(ctx context.Context, id uint, amount decimal.Decimal) error
}

type TranscationRepo interface {
	// SaveTranscation 交易持久化
	SaveTranscation(ctx context.Context, ts *types.Transcation) error
}

// GetVirtualWallet 获取用户虚拟钱包
func (w *WalletUseCase) GetVirtualWallet(ctx context.Context, id uint) (*types.VirtualWallet, error) {
	return w.walletRepo.GetVirtualWallet(ctx, id)
}

// Debit 付款
func (w *WalletUseCase) Debit(ctx context.Context, userID uint, amount decimal.Decimal) error {
	// 获取虚拟钱包领域模型
	wallet, err := w.walletRepo.GetVirtualWallet(ctx, userID)
	if err != nil {
		return err
	}

	// 领域模型付款逻辑
	if err := wallet.Debit(amount); err != nil {
		return err
	}

	// 生成流水
	ts := types.NewTranscation(userID, wallet.ID(), amount, types.TranscationTypeDebit)
	if err := w.transcationRepo.SaveTranscation(ctx, ts); err != nil {
		return err
	}

	// 更新钱包余额
	if err := w.walletRepo.UpdateBalance(ctx, wallet.ID(), wallet.Balance()); err != nil {
		return err
	}

	return nil
}
