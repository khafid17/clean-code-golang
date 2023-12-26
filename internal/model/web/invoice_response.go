package web

type InvoiceResponse struct {
	Status         string         `json:"status"`
	Message        string         `json:"message"`
	InvoiceSummary InvoiceSummary `json:"invoice_summary"`
}

type InvoiceSummary struct {
	TotalItems int     `json:"total_items"`
	SubTotal   float64 `json:"sub_total"`
	Tax        float64 `json:"tax"`
	GrandTotal float64 `json:"grand_total"`
}
