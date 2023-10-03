package models

type GeneralResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Token	string		`json:"token,omitempty"`
}

type ReqInquiry struct {
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Price			int			`json:"price" gorm:"column:price"`
	Quantity		int			`json:"quantity" gorm:"column:quantity"`
	Percentage		int			`json:"percentage"`
}