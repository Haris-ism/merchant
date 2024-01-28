package usecase_grpc

import (
	"fmt"
	"merchant/protogen/merchant"

	postgre "merchant/databases/postgresql"
	"merchant/databases/postgresql/models"
	redis_db "merchant/databases/redis"
)

type (
	usecaseGrpc struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
		// host	host.HostInterface
	}
	UsecaseGrpcInterface interface {
		InquiryItems()([]*merchant.InquiryItemsModel, error)
		InquiryDiscounts()([]*merchant.InquiryDiscountsModel, error)
		redisInquiryItems()([]models.InquiryItems,error)
		redisInquiryDiscounts()([]models.InquiryDiscounts,error)
		OrderTransItem(req *merchant.ReqTransItemsModel)(string,error)
	}
)

func InitUsecaseGrpc(postgre postgre.PostgreInterface, redis redis_db.RedisInterface) UsecaseGrpcInterface {
	fmt.Println("init uc grpc")
	return &usecaseGrpc{
		postgre: postgre,
		redis:   redis,
		// host: host,
	}
}
