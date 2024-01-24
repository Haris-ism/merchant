package usecase

import (
	"errors"
	"fmt"
	"merchant/controllers/models"
	dbModels "merchant/databases/postgresql/models"
	"merchant/utils"
	"strconv"

	"github.com/sirupsen/logrus"
)


func (uc *usecase)OrderTransItem(req models.DecReqTransItem)(models.DecReqTransItem,error){
	res:=models.DecReqTransItem{}
	codes:=[]string{}
	decryptedReq,err:=utils.DecryptTransItem(req)
	if err!=nil{
		logrus.Error(err)
		return res, err
	}
	// fmt.Println("decryptedReq",decryptedReq)
	amount,_:=strconv.Atoi(decryptedReq.Amount)
	qtys,_:=strconv.Atoi(decryptedReq.Quantity)
	items,err:=uc.postgre.CheckItems(decryptedReq)
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
	reqDB.CC=decryptedReq.CC
	reqDB.Discount=decryptedReq.Discount

	err=uc.postgre.OrderTransItem(reqDB)
	if err!=nil{
		return res,err
	}
	resp:=models.DecTransItem{}

	resp.ID=strconv.Itoa(int(items.ID))
	resp.Name=items.Name
	resp.CC=decryptedReq.CC
	resp.Quantity=decryptedReq.Quantity
	resp.Code=codes
	fmt.Println("res real:",res)
	encryptedRes,err:=utils.EncryptTransItemRes(resp)
	if err!=nil{
		logrus.Error(err)
		return res, err
	}
	fmt.Println("encryptedRes",encryptedRes)

	return encryptedRes,nil
}