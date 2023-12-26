package invoicecontroller

import (
	"esb/internal/controller"
	model "esb/internal/model/entity"
	"esb/internal/usecase"
)

type InvoiceControllerImpl struct {
	useCase usecase.InvoiceUseCase
}

func NewInvoiceController(useCase usecase.InvoiceUseCase) controller.InvoiceController {
	return &InvoiceControllerImpl{
		useCase: useCase,
	}
}

func (c *InvoiceControllerImpl) CreateInvoice(subject string, issue_date string, due_date string, address string, customer string, items []model.InvoiceItem, total_items int, sub_total float64, grand_total float64, status int) error {
	return c.useCase.CreateInvoice(subject, issue_date, due_date, address, customer, items, total_items, sub_total, grand_total, status)
}

func (c *InvoiceControllerImpl) GetAllInvoices() ([]*model.Invoice, error) {
	return c.useCase.GetAllInvoices()
}

func (c *InvoiceControllerImpl) GetAllCustomers() ([]*model.Customer, error) {
	return c.useCase.GetAllCustomers()
}

func (c *InvoiceControllerImpl) GetAllItems() ([]*model.Items, error) {
	return c.useCase.GetAllItems()
}