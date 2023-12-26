package controller

import (
	model "esb/internal/model/entity"
	"esb/internal/usecase"
)

type InvoiceController struct {
	useCase usecase.InvoiceUseCase
}

func NewInvoiceController(useCase usecase.InvoiceUseCase) *InvoiceController {
	return &InvoiceController{
		useCase: useCase,
	}
}

func (c *InvoiceController) CreateInvoice(subject string, issue_date string, due_date string, address string, customer string, items []model.InvoiceItem, total_items int, sub_total float64, grand_total float64, status int) error {
	return c.useCase.CreateInvoice(subject, issue_date, due_date, address, customer, items, total_items, sub_total, grand_total, status)
}

func (c *InvoiceController) GetAllInvoices() ([]*model.Invoice, error) {
	return usecase.GetAllInvoices()
}

func (c *InvoiceController) GetAllCustomers() ([]*model.Customer, error) {
	return usecase.GetAllCustomers()
}

func (c *InvoiceController) GetAllItems() ([]*model.Items, error) {
	return usecase.GetAllItems()
}
