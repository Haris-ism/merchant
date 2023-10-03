package postgre

import (
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
		InquiryItems()(dbs.Items,error)
		InquiryDiscounts()(dbs.Discounts,error)
		AddInquiryItems(items dbs.Items)error
		AddInquiryDiscounts(discounts dbs.Discounts)error
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
	db.AutoMigrate(&dbs.Items{},&dbs.Discounts{})

	return &postgreDB{
		postgre: db,
	}
}
