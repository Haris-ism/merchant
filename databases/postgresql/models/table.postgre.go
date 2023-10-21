package models

import "gorm.io/gorm"

type InquiryItems struct{
	ID				int			`json:"ID"`
	Name			string		`json:"name" gorm:"column:item_name"`
	Type			string		`json:"type" gorm:"column:type"`
	Price			int			`json:"price" gorm:"column:price"`
}
func (t *InquiryItems)TableName()string{
	return"merchant_items"
}
type Items struct{
	gorm.Model
	Name			string		`json:"name" gorm:"column:item_name"`
	Type			string		`json:"type" gorm:"column:type"`
	Price			int			`json:"price" gorm:"column:price"`
	Percentage		int			`json:"percentage,omitempty" gorm:"column:percentage"`
}

func (t *Items)TableName()string{
	return"merchant_items"
}

type Discounts struct{
	gorm.Model
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Percentage		int			`json:"percentage" gorm:"column:percentage"`
}

func (t *Discounts)TableName()string{
	return"merchant_discounts"
}

type InquiryDiscounts struct{
	ID				int			`json:"ID"`
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Percentage		int			`json:"percentage" gorm:"column:percentage"`
}

func (t *InquiryDiscounts)TableName()string{
	return"merchant_discounts"
}

type Order struct{
	gorm.Model
	ItemID			int			`json:"item_id" gorm:"column:item_id"`
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Discount		string		`json:"discount" gorm:"column:discount"`
	Price			int			`json:"price" gorm:"column:price"`
	TotalPrice		int			`json:"total_price" gorm:"column:total_price"`
	Quantity		int			`json:"quantity" gorm:"column:quantity"`
	CC				string		`json:"cc" gorm:"column:cc"`
}

func (t *Order)TableName()string{
	return"merchant_orders"
}

type Vouchers struct{
	gorm.Model
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Price			int			`json:"price" gorm:"column:price"`
	Code			string		`json:"code" gorm:"column:code"`
	Status			string		`json:"status" gorm:"column:status"`
}


func (t *Vouchers)TableName()string{
	return "merchant_vouchers"
}