package service

import v1 "github.com/haappywater/virtualwallet/api/v1"

type WalletService struct {
	v1.UnimplementedWalletServer
}

func NewWalletService() *WalletService {
	return &WalletService{}
}
