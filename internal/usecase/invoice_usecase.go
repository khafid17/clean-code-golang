package usecase

import model "esb/internal/model/entity"

type InvoiceUseCase interface {
	CreateInvoice(subject string, issue_date string, due_date string, address string, customer string, items []model.InvoiceItem, total_items int, sub_total float64, grand_total float64, status int) error
	GetAllInvoices() ([]*model.Invoice, error)
	GetAllCustomers() ([]*model.Customer, error)
	GetAllItems() ([]*model.Items, error)
}
