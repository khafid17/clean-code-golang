package repository

import model "esb/internal/model/entity"

type InvoiceRepository interface {
	CreateInvoice(invoice *model.Invoice) error
	GetAllInvoices() ([]*model.Invoice, error)
	GetAllCustomers() ([]*model.Customer, error)
	GetAllItems() ([]*model.Items, error)
}
