package web

import model "esb/internal/model/entity"

type InvoiceRequest struct {
	Subject    string              `json:"subject"`
	IssueDate  string              `json:"issue_date"`
	DueDate    string              `json:"due_date"`
	Address    string              `json:"address"`
	Customer   string              `json:"customer"`
	Items      []model.InvoiceItem `json:"items"`
	TotalItems int                 `json:"total_items"`
	Status     int                 `json:"status"`
}
