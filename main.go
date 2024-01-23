package main

import (
	"merchant/controller_grpc"
	controller "merchant/controllers"
	postgre "merchant/databases/postgresql"
	redis_db "merchant/databases/redis"
	grpc_server "merchant/grpc/server"
	router "merchant/routers"
	"merchant/usecase_grpc"
	usecase "merchant/usecases"
)

func main() {
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis)
	con := controller.InitController(uc)
	ucGrpc:=usecase_grpc.InitUsecaseGrpc(postgre,redis)
	conGrpc:=controller_grpc.InitControllerGrpc(ucGrpc)

	go func(){
		grpc_server.InitGrpcServer(conGrpc)
	}()
	router.MainRouter(con)

}
