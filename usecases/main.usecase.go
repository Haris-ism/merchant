package usecase

import (
	con "merchant/controllers/models"
	dbs "merchant/databases/postgresql/models"
	"merchant/models"

	postgre "merchant/databases/postgresql"
	redis_db "merchant/databases/redis"
)

type (
	usecase struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
	}
	UsecaseInterface interface {
		WriteRedis(models.RedisReq) error
		ReadRedis(req models.RedisReq) (string, error)
		InsertDB(req models.ItemList) error
		InquiryItems()([]dbs.InquiryItems,error)
		InquiryDiscounts()([]dbs.InquiryDiscounts,error)
		AddInquiryItems(req con.ReqInquiry)error
		AddInquiryDiscounts(req con.ReqInquiry)error
		GenVoucher(req con.ReqGenerateVoucher)(string,error)
		OrderTransItem(req con.DecReqTransItem)(con.DecReqTransItem,error)
	}
)

func InitUsecase(postgre postgre.PostgreInterface, redis redis_db.RedisInterface) UsecaseInterface {
	return &usecase{
		postgre: postgre,
		redis:   redis,
	}
}
