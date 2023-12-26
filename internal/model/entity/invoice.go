package model

type Invoice struct {
	ID         int           `json:"id"`
	Subject    string        `json:"subject"`
	IssueDate  string        `json:"issue_date"`
	DueDate    string        `json:"due_date"`
	Address    string        `json:"address"`
	Customer   string        `json:"customer"`
	Items      []InvoiceItem `json:"items"`
	TotalItems int           `json:"total_items"`
	SubTotal   float64       `json:"sub_total"`
	GrandTotal float64       `json:"grand_total"`
	Status     int           `json:"status"`
}

type InvoiceItem struct {
	Qty    float64 `json:"qty"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tax    float64 `json:"tax"`
}
