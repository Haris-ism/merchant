package controller_grpc

import (
	"context"
	"fmt"
	"merchant/constants"
	"merchant/protogen/merchant"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *ControllerGrpc) InquiryItems(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantItemsModel, error){
	fmt.Println("masuk con")
	res:=merchant.InquiryMerchantItemsModel{
		Message: constants.SUCCESS,
		Code: http.StatusOK,
	}
	result,err:=g.uc.InquiryItems()
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		return &res,err
	}
	fmt.Println("grps res:",result)
	res.Data=result
	return &res,nil
}

func (g *ControllerGrpc) InquiryDiscounts(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantDiscountsModel, error){
	fmt.Println("masuk con discount")
	res:=merchant.InquiryMerchantDiscountsModel{
		Message: constants.SUCCESS,
		Code: http.StatusOK,
	}
	result,err:=g.uc.InquiryDiscounts()
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		return &res,err
	}
	fmt.Println("grps res:",result)
	res.Data=result
	return &res,nil
}