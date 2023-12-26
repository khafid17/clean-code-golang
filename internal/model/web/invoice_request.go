package web

type InvoiceRequest struct {
	Subject   string   `json:"subject"`
	IssueDate string   `json:"issue_date"`
	DueDate   string   `json:"due_date"`
	Address   string   `json:"address"`
	Customer  string   `json:"customer"`
	Items     []string `json:"items"`
	Qty       int      `json:"qty"`
	Price     float64  `json:"price"`
	Amount    float64  `json:"amount"`
	Status    int      `json:"status"`
}
