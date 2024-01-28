package grpc_server

import (
	"fmt"
	"merchant/controller_grpc"
	"merchant/protogen/merchant"
	"merchant/utils"
	"net"

	"google.golang.org/grpc"
)

type (
	GrpcServer struct {
		Config *grpc.Server
		merchant.MerchantServicesServer
		TCP	net.Listener
	}
	ControllerGrpcInterface interface {
		// Run()
	}
)

func InitGrpcServer(grpcCon controller_grpc.ControllerGrpc)  {

	listen,err:=net.Listen("tcp",utils.GetEnv("PORT_GRPC"))
	if err!=nil{
		fmt.Println("failed to listen tcp:",err)
	}

	fmt.Println("register start")
	merchant.RegisterMerchantServicesServer(grpcCon.Config,&grpcCon)
	fmt.Println("register grpc server")

	err=grpcCon.Config.Serve(listen)

	if err!=nil{
		fmt.Println("failed to listen grpc:",err)
	}
}