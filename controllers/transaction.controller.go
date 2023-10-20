package controller

import (
	"merchant/constants"
	"merchant/controllers/models"
	"merchant/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *controller)TransItem(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:http.StatusOK,
	}
	req:=models.ReqTransItem{}
	if err:=ctx.BindJSON(&req);err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		logrus.Error(res.Message)
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	reqHeader:=models.ReqHeader{}
	if err:=ctx.BindHeader(&reqHeader);err!=nil{
		logrus.Error(err)
		res.Message=err.Error()
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	
	err:=utils.SignatureValidation(reqHeader,req)
	if err!=nil{
		logrus.Error(err)
		res.Message=err.Error()
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	data,err:=c.usecase.OrderTransItem(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		logrus.Error(res.Message)
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	
	res.Data=data
	ctx.JSON(res.Code,res)
}