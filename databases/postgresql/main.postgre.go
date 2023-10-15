package postgre

import (
	cModels "merchant/controllers/models"
	con "merchant/controllers/models"
	dbs "merchant/databases/postgresql/models"
	"merchant/models"
	"merchant/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	postgreDB struct {
		postgre *gorm.DB
	}
	PostgreInterface interface {
		Insert(req models.ItemList) error
		QueryInquiryItems(name string) (dbs.Items,error)
		QueryInquiryDiscounts(name string) (dbs.Discounts,error)
		InquiryItems()([]dbs.InquiryItems,error)
		InquiryDiscounts()([]dbs.InquiryDiscounts,error)
		AddInquiryItems(items dbs.Items)error
		AddInquiryDiscounts(discounts dbs.Discounts)error
		GenVoucher(req con.ReqGenerateVoucher,voucher string)error
		CheckItems(req cModels.ReqTransItem)(dbs.Items,error)
		OrderTransItem(req dbs.Order)error
	}
)

func InitPostgre() PostgreInterface {
	host := utils.GetEnv("POSTGRE")
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Errorf("Failed to Init Postgre, Err:", err)
	} else {
		logrus.Printf("Init Postgre Success")
	}
	db.AutoMigrate(&dbs.Items{},&dbs.Discounts{},&dbs.Vouchers{},dbs.Order{})

	return &postgreDB{
		postgre: db,
	}
}
