package postgre

import (
	cModels "merchant/controllers/models"
	"merchant/databases/postgresql/models"
)



func (db *postgreDB)CheckItems(req cModels.ReqTransItem)(models.Items,error){
	reqDB:=models.Items{}
	err:=db.postgre.Select("merchant_items.ID, merchant_items.item_name , merchant_items.type, merchant_items.price, merchant_discounts.percentage, merchant_discounts.name as discount_name").
		Joins("left join merchant_discounts on merchant_discounts.type = merchant_items.type").
		Where("merchant_items.id = ? and merchant_discounts.name = ?",req.ID,req.Discount).Find(&reqDB).Error
	if err!=nil{
		return reqDB,err
	}
	return reqDB,nil
}

func (db *postgreDB)OrderTransItem(req models.Order)error{
	
	err:=db.postgre.Create(&req).Error
	if err!=nil{
		return err
	}

	return nil
}