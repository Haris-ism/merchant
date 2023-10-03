package main

import (
	controller "merchant/controllers"
	postgre "merchant/databases/postgresql"
	redis_db "merchant/databases/redis"
	router "merchant/routers"
	usecase "merchant/usecases"
)

func main() {
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis)
	con := controller.InitController(uc)

	router.MainRouter(con)

}
