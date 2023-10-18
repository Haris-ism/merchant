package models


type ReqGenerateVoucher struct{
	Name		string		`json:"name"`
	Type		string		`json:"type"`
	Price		int			`json:"price"`
}

type ReqTransItem struct{
	ID			string			`json:"item_id"`
	Discount	string			`json:"discount"`
	Quantity	string			`json:"quantity"`
	CC			string			`json:"cc_number"`
	Amount		string			`json:"amount"`
}

type ResTransItem struct{
	ID			string		`json:"item_id"`
	Name		string		`json:"item_name"`
	Quantity	string		`json:"quantity"`
	CC			string		`json:"cc_number"`
	Code		string		`json:"code"`
}
type DecTransItem struct{
	ID			string		`json:"item_id"`
	Name		string		`json:"item_name"`
	Quantity	string		`json:"quantity"`
	CC			string		`json:"cc_number"`
	Code		[]string	`json:"code"`
}