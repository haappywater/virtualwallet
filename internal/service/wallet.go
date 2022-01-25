package service

import (
	v1 "github.com/haappywater/virtualwallet/api/v1"
	"github.com/haappywater/virtualwallet/internal/biz"
)

type WalletService struct {
	v1.UnimplementedWalletServer

	wallet *biz.WalletUseCase
}

func NewWalletService(wallet *biz.WalletUseCase) *WalletService {
	return &WalletService{wallet: wallet}
}
