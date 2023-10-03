package models

import "gorm.io/gorm"

type Items struct{
	gorm.Model
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Price			int			`json:"price" gorm:"column:price"`
	Quantity		int			`json:"quantity" gorm:"column:quantity"`
}

func (t *Items)TableName()string{
	return"items"
}

type Discounts struct{
	gorm.Model
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Percentage		int			`json:"percentage" gorm:"column:percentage"`
}

func (t *Discounts)TableName()string{
	return"discounts"
}