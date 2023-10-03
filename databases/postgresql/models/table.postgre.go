package models

import "gorm.io/gorm"

type Items struct{
	gorm.Model
	Name			string		`gorm:"column:name"`
	Type			string		`gorm:"column:type"`
	Price			int			`gorm:"column:price"`
	Quantity		int			`gorm:"column:quantity"`
}

func (t *Items)TableName()string{
	return"items"
}

type Discounts struct{
	gorm.Model
	Name			string		`gorm:"column:name"`
	Type			string		`gorm:"column:type"`
}

func (t *Discounts)TableName()string{
	return"discounts"
}