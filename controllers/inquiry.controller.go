package controller

import (
	"merchant/constants"
	"merchant/controllers/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller)InquiryItems(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code: http.StatusOK,
	}
	result,err:=c.usecase.InquiryItems()
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	res.Data=result
	ctx.JSON(http.StatusOK,res)
}
func (c *controller)InquiryDiscounts(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code: http.StatusOK,
	}
	result,err:=c.usecase.InquiryDiscounts()
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	res.Data=result
	ctx.JSON(http.StatusOK,res)
}

func (c *controller)AddInquiryItems(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code: http.StatusOK,
	}
	req:=models.ReqInquiry{}
	err:=ctx.BindJSON(&req)
	if err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	err=c.usecase.AddInquiryItems(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	ctx.JSON(http.StatusOK,res)
}
func (c *controller)AddInquiryDiscounts(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code: http.StatusOK,
	}
	req:=models.ReqInquiry{}
	err:=ctx.BindJSON(&req)
	if err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	err=c.usecase.AddInquiryDiscounts(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	ctx.JSON(http.StatusOK,res)
}