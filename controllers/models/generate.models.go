package models


type ReqGenerateVoucher struct{
	Name		string		`json:"name"`
	Type		string		`json:"type"`
	Price		int			`json:"price"`
}

type ReqTransItem struct{
	ID			int			`json:"item_id"`
	Discount	string		`json:"discount"`
	Quantity	int			`json:"quantity"`
	CC			string		`json:"cc_number"`
	Amount		int			`json:"amount"`
}

type ResTransItem struct{
	ID			int			`json:"item_id"`
	Name		string		`json:"item_name"`
	Quantity	int			`json:"quantity"`
	CC			string		`json:"cc_number"`
	Code		[]string	`json:"code"`
}