package usecase

import (
	"encoding/json"
	"errors"
	"merchant/constants"
	"merchant/controllers/models"
	dbs "merchant/databases/postgresql/models"
)

func (uc *usecase)InquiryItems()([]dbs.InquiryItems,error){
	result,err:=uc.redisInquiryItems()
	if err!=nil || len(result)<1 {
		result,err=uc.postgre.InquiryItems()
		if err!=nil{
			return result,errors.New(constants.ERROR_DB)
		}
		bytes,err:=json.Marshal(result)
		if err!=nil{
			return result,errors.New(constants.ERROR_DB)
		}
		err=uc.redis.WriteRedis(constants.INQUIRY_ITEMS,string(bytes),300)
		if err!=nil{
			return result,errors.New(constants.ERROR_DB)
		}
	}
	return result,nil
}

func (uc *usecase)redisInquiryItems()([]dbs.InquiryItems,error){
	result:=[]dbs.InquiryItems{}
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

func (uc *usecase)InquiryDiscounts()([]dbs.InquiryDiscounts,error){
	result,err:=uc.redisInquiryDiscounts()
	if err!=nil || len(result)<1{
		result,err=uc.postgre.InquiryDiscounts()
		if err!=nil{
			return result,errors.New(constants.ERROR_DB)
		}
		bytes,err:=json.Marshal(result)
		if err!=nil{
			return result,errors.New(constants.ERROR_DB)
		}
		err=uc.redis.WriteRedis(constants.INQUIRY_DISCOUNTS,string(bytes),300)
		if err!=nil{
			return result,errors.New(constants.ERROR_DB)
		}
	}
	return result,nil
}

func (uc *usecase)redisInquiryDiscounts()([]dbs.InquiryDiscounts,error){
	result:=[]dbs.InquiryDiscounts{}
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

func (uc *usecase)AddInquiryItems(req models.ReqInquiry)error{
	items:=dbs.Items{}
	res,err:=uc.postgre.QueryInquiryItems(req.Name)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}
	if res.ID!=0{
		return errors.New(constants.ERROR_ITEMS)
	}
	items.Name=req.Name
	items.Price=req.Price
	items.Type=req.Type
	err=uc.postgre.AddInquiryItems(items)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}
	err=uc.redis.WriteRedis(constants.INQUIRY_ITEMS,"",0)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}
	return nil
}
func (uc *usecase)AddInquiryDiscounts(req models.ReqInquiry)error{
	discounts:=dbs.Discounts{}
	res,err:=uc.postgre.QueryInquiryDiscounts(req.Name)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}
	if res.ID!=0{
		return errors.New(constants.ERROR_DISCOUNTS)
	}
	discounts.Name=req.Name
	discounts.Type=req.Type
	discounts.Percentage=req.Percentage
	err=uc.postgre.AddInquiryDiscounts(discounts)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}
	err=uc.redis.WriteRedis(constants.INQUIRY_DISCOUNTS,"",0)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}
	return nil
}