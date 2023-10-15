package controller

import (
	usecase "merchant/usecases"

	"github.com/gin-gonic/gin"
)

type (
	controller struct {
		usecase usecase.UsecaseInterface
	}
	ControllerInterface interface {
		Ping(ctx *gin.Context)
		WriteRedis(ctx *gin.Context)
		ReadRedis(ctx *gin.Context)
		InsertPostgre(ctx *gin.Context)
		InquiryItems(ctx *gin.Context)
		InquiryDiscounts(ctx *gin.Context)
		AddInquiryItems(ctx *gin.Context)
		AddInquiryDiscounts(ctx *gin.Context)
		TransItem(ctx *gin.Context)
		GenVoucher(ctx *gin.Context)
	}
)

func InitController(uc usecase.UsecaseInterface) ControllerInterface {
	return &controller{
		usecase: uc,
	}
}
