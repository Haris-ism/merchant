package postgre

import (
	"merchant/controllers/models"
	dModels "merchant/databases/postgresql/models"
)

func (db *postgreDB)GenVoucher(req models.ReqGenerateVoucher,voucher string)error{
	table:=dModels.Vouchers{
		Name:req.Name,
		Type:req.Type,
		Status:"Not Used",
		Code:voucher,
		Price:req.Price,
	}

	err:=db.postgre.Save(&table).Error
	if err!=nil{
		return err
	}
	return nil
}