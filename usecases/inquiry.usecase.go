package usecase

import (
	"errors"
	"merchant/constants"
	"merchant/controllers/models"
	dbs "merchant/databases/postgresql/models"
)

func (uc *usecase)InquiryItems()(dbs.Items,error){
	items,err:=uc.postgre.InquiryItems()
	if err!=nil{
		return items,errors.New(constants.ERROR_DB)
	}
	return items,nil
}
func (uc *usecase)InquiryDiscounts()(dbs.Discounts,error){
	discounts,err:=uc.postgre.InquiryDiscounts()
	if err!=nil{
		return discounts,errors.New(constants.ERROR_DB)
	}
	return discounts,nil
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
	items.Quantity=req.Quantity
	err=uc.postgre.AddInquiryItems(items)
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
	return nil
}