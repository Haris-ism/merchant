package usecase

import (
	"errors"
	"merchant/controllers/models"
	dbModels "merchant/databases/postgresql/models"
	"merchant/utils"
)


func (uc *usecase)OrderTransItem(req models.ReqTransItem)(models.ResTransItem,error){
	res:=models.ResTransItem{}
	codes:=[]string{}
	
	items,err:=uc.postgre.CheckItems(req)
	if err!=nil{
		return res,err
	}
	if items.ID==0{
		return res,errors.New("Invalid Item ID")
	}

	discount:=float64(items.Percentage)/100
	price:=float64(items.Price)
	qty:=float64(req.Quantity)
	reqTotalPrice:=int((price-(price*discount))*qty)

	if reqTotalPrice!=req.Amount{
		return res,errors.New("Invalid Amount")
	}

	for i:=0;i<req.Quantity;i++{
		code,err:=utils.GenerateRandom(25)
		if err!=nil{
			return res,errors.New("Failed to Generate Voucher")
		}
		codes=append(codes,code)
	}

	reqDB:=dbModels.Order{}
	reqDB.ItemID=int(items.ID)
	reqDB.Name=items.Name
	reqDB.Type=items.Type
	reqDB.Price=items.Price
	reqDB.TotalPrice=reqTotalPrice
	reqDB.Quantity=req.Quantity
	reqDB.CC=req.CC
	reqDB.Discount=req.Discount

	err=uc.postgre.OrderTransItem(reqDB)
	if err!=nil{
		return res,err
	}
	res.ID=int(items.ID)
	res.Name=items.Name
	res.CC=req.CC
	res.Quantity=req.Quantity
	res.Code=codes
	return res,nil
}