package controller_grpc

import (
	"context"
	"fmt"
	"log"
	"merchant/constants"
	"merchant/protogen/merchant"
	"merchant/utils"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func (g *ControllerGrpc) TransItems(ctx context.Context,req *merchant.ReqTransItemsModel) (*merchant.ResMerchantTransModel, error){
	fmt.Println("masuk con")
	md, ok := metadata.FromIncomingContext(ctx)
	log.Println("ieu md:",md.Get("Signature"))
	log.Println("ieu md:",md.Get("TimeStamp"))
	log.Println("ieu ok:",ok)
	log.Println("ieu req:",req)
	res:=&merchant.ResMerchantTransModel{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	err:=utils.SignatureValidationGrpc(md,req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		return res,err
	}
	log.Println("signature is valid")
	resp,err:=g.uc.OrderTransItem(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		return res,err
	}
	res.Data=resp
	return res,nil
}
