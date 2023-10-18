package usecase

import (
	"errors"
	"merchant/controllers/models"
	dbModels "merchant/databases/postgresql/models"
	"merchant/utils"
	"strconv"

	"github.com/sirupsen/logrus"
)


func (uc *usecase)OrderTransItem(req models.ReqTransItem)(models.ResTransItem,error){
	res:=models.ResTransItem{}
	codes:=[]string{}
	req,err:=utils.DecryptTransItem(req)
	if err!=nil{
		logrus.Error(err)
		return res, err
	}
	
	// id,_:=strconv.Atoi(req.ID)
	amount,_:=strconv.Atoi(req.Amount)
	qtys,_:=strconv.Atoi(req.Quantity)
	items,err:=uc.postgre.CheckItems(req)
	if err!=nil{
		logrus.Error(err)
		return res,err
	}
	if items.ID==0{
		logrus.Error(errors.New("Invalid Item ID"))
		return res,errors.New("Invalid Item ID")
	}


	discount:=float64(items.Percentage)/100
	price:=float64(items.Price)
	qty:=float64(qtys)
	reqTotalPrice:=int((price-(price*discount))*qty)

	if reqTotalPrice!=amount{
		logrus.Error(errors.New("Invalid Amount"))
		return res,errors.New("Invalid Amount")
	}

	for i:=0;i<qtys;i++{
		code,err:=utils.GenerateRandom(25)
		if err!=nil{
			logrus.Error(errors.New("Failed to Generate Voucher"))
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
	reqDB.Quantity=qtys
	reqDB.CC=req.CC
	reqDB.Discount=req.Discount

	err=uc.postgre.OrderTransItem(reqDB)
	if err!=nil{
		return res,err
	}

	res.ID=strconv.Itoa(int(items.ID))
	res.Name=items.Name
	res.CC=req.CC
	res.Quantity=req.Quantity

	res,err=utils.EncryptTransItemRes(res,codes)
	if err!=nil{
		logrus.Error(err)
		return res, err
	}

	return res,nil
}