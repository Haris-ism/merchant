package usecase_grpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"merchant/constants"
	"merchant/databases/postgresql/models"
	"merchant/protogen/merchant"
)

func (uc *usecaseGrpc)InquiryItems()([]*merchant.InquiryItemsModel, error){
	fmt.Println("masuk uc")
	res:=[]*merchant.InquiryItemsModel{}
	result,err:=uc.redisInquiryItems()
	if err!=nil || len(result)<1 {
		result,err=uc.postgre.InquiryItems()
		if err!=nil{
			return res,errors.New(constants.ERROR_DB)
		}
		bytes,err:=json.Marshal(result)
		if err!=nil{
			return res,errors.New(constants.ERROR_DB)
		}
		err=json.Unmarshal(bytes,&res)

		err=uc.redis.WriteRedis(constants.INQUIRY_ITEMS,string(bytes),300)
		if err!=nil{
			return res,errors.New(constants.ERROR_DB)
		}
		return res,nil
	}

	bytes,err:=json.Marshal(result)
		if err!=nil{
			return res,errors.New(constants.ERROR_DB)
		}
		err=json.Unmarshal(bytes,&res)
		
	return res,nil
}
func (uc *usecaseGrpc)InquiryDiscounts()([]*merchant.InquiryDiscountsModel, error){
	fmt.Println("masuk uc")
	res:=[]*merchant.InquiryDiscountsModel{}
	result,err:=uc.redisInquiryDiscounts()
	if err!=nil || len(result)<1{
		result,err=uc.postgre.InquiryDiscounts()
		if err!=nil{
			return res,errors.New(constants.ERROR_DB)
		}
		bytes,err:=json.Marshal(result)
		if err!=nil{
			return res,errors.New(constants.ERROR_DB)
		}
		
		err=json.Unmarshal(bytes,&res)
		err=uc.redis.WriteRedis(constants.INQUIRY_DISCOUNTS,string(bytes),300)
		if err!=nil{
			return res,errors.New(constants.ERROR_DB)
		}
		return res,nil
	}

	bytes,err:=json.Marshal(result)
	if err!=nil{
		return res,errors.New(constants.ERROR_DB)
	}
	err=json.Unmarshal(bytes,&res)
		
	return res,nil
}

func (uc *usecaseGrpc)redisInquiryItems()([]models.InquiryItems,error){
	result:=[]models.InquiryItems{}
	redisItems,err:=uc.redis.ReadRedis(constants.INQUIRY_ITEMS)
	if err!=nil{
		return result,errors.New(constants.ERROR_DB)
	}
	err=json.Unmarshal([]byte(redisItems),&result)
	if err!=nil{
		return result,errors.New(constants.ERROR_DB)
	}
	return result,nil
}

func (uc *usecaseGrpc)redisInquiryDiscounts()([]models.InquiryDiscounts,error){
	result:=[]models.InquiryDiscounts{}
	redisItems,err:=uc.redis.ReadRedis(constants.INQUIRY_DISCOUNTS)
	if err!=nil{
		return result,errors.New(constants.ERROR_DB)
	}
	err=json.Unmarshal([]byte(redisItems),&result)
	if err!=nil{
		return result,errors.New(constants.ERROR_DB)
	}
	return result,nil
}