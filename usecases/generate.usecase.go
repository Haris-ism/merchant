package usecase

import (
	"errors"
	"merchant/constants"
	"merchant/controllers/models"
	"merchant/utils"
)

func (uc *usecase)GenVoucher(req models.ReqGenerateVoucher)(string,error){
	voucher,_:=utils.GenerateRandom(25)

	err:=uc.postgre.GenVoucher(req,voucher)
	
	if err!=nil{
		return voucher,errors.New(constants.ERROR_DB)
	}

	return voucher,nil
}