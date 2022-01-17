// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WalletHTTPServer interface {
	GetVirtualWallet(context.Context, *GetVirtualWalletReq) (*GetVirtualWalletResp, error)
}

func RegisterWalletHTTPServer(s *http.Server, srv WalletHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/wallet/virtual_wallet/get", _Wallet_GetVirtualWallet0_HTTP_Handler(srv))
}

func _Wallet_GetVirtualWallet0_HTTP_Handler(srv WalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVirtualWalletReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wallet.v1.Wallet/GetVirtualWallet")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVirtualWallet(ctx, req.(*GetVirtualWalletReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVirtualWalletResp)
		return ctx.Result(200, reply)
	}
}

type WalletHTTPClient interface {
	GetVirtualWallet(ctx context.Context, req *GetVirtualWalletReq, opts ...http.CallOption) (rsp *GetVirtualWalletResp, err error)
}

type WalletHTTPClientImpl struct {
	cc *http.Client
}

func NewWalletHTTPClient(client *http.Client) WalletHTTPClient {
	return &WalletHTTPClientImpl{client}
}

func (c *WalletHTTPClientImpl) GetVirtualWallet(ctx context.Context, in *GetVirtualWalletReq, opts ...http.CallOption) (*GetVirtualWalletResp, error) {
	var out GetVirtualWalletResp
	pattern := "/api/v1/wallet/virtual_wallet/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.wallet.v1.Wallet/GetVirtualWallet"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}