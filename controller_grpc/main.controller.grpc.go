package controller_grpc

import (
	"context"
	"merchant/protogen/merchant"

	"merchant/usecase_grpc"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	ControllerGrpc struct {
		Config *grpc.Server
		merchant.InquiryServicesServer
		uc usecase_grpc.UsecaseGrpcInterface
	}
	ControllerGrpcInterface interface {
		InquiryItems(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantItemsModel, error)
	}
)

func InitControllerGrpc(uc usecase_grpc.UsecaseGrpcInterface) ControllerGrpc {
	grpcConn:=grpc.NewServer()

	return ControllerGrpc{
		uc: uc,
		Config:grpcConn,
	}
}
