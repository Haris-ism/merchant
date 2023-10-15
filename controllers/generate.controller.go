package controller

import (
	"merchant/constants"
	"merchant/controllers/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller)GenVoucher(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:http.StatusOK,
	}
	req:=models.ReqGenerateVoucher{}
	if err:=ctx.BindJSON(&req);err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		return
	}
	data,err:=c.usecase.GenVoucher(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		return
	}
	res.Data=data
	ctx.JSON(http.StatusOK,res)
}