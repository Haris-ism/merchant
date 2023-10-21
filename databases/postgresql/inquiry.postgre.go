package postgre

import (
	"errors"
	"merchant/constants"
	"merchant/databases/postgresql/models"
)

func (db *postgreDB)InquiryItems()([]models.InquiryItems,error){
	items:=[]models.InquiryItems{}
	err:=db.postgre.Find(&items).Error
	if err!=nil{
		return items,errors.New(constants.ERROR_INQUIRY)
	}
	return items,nil
}
func (db *postgreDB)InquiryDiscounts()([]models.InquiryDiscounts,error){
	discounts:=[]models.InquiryDiscounts{}
	err:=db.postgre.Find(&discounts).Error
	if err!=nil{
		return discounts,errors.New(constants.ERROR_INQUIRY)
	}
	return discounts,nil
}

func (db *postgreDB)QueryInquiryItems(name string) (models.Items,error){
	items:=models.Items{}
	err:=db.postgre.Where("item_name = ?",name).Find(&items).Error
	if err!=nil{
		return items,err
	}
	return items,nil
}
func (db *postgreDB)QueryInquiryDiscounts(name string) (models.Discounts,error){
	discounts:=models.Discounts{}
	err:=db.postgre.Where("name = ?",name).Find(&discounts).Error
	if err!=nil{
		return discounts,err
	}
	return discounts,nil
}

func (db *postgreDB)AddInquiryItems(items models.Items)error{
	err:=db.postgre.Save(&items).Error
	if err!=nil{
		return err
	}
	return nil
}
func (db *postgreDB)AddInquiryDiscounts(discounts models.Discounts)error{
	err:=db.postgre.Save(&discounts).Error
	if err!=nil{
		return err
	}
	return nil
}