// controller/invoice_controller.go
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

func (c *InvoiceController) CreateInvoice(subject string, issue_date string, due_date string, address string, customer string, items string, qty int, price float64, amount float64, status int) error {
	// Additional processing, input validation, etc.
	// Call use case to perform business logic
	return c.useCase.CreateInvoice(subject, issue_date, due_date, address, customer, items, qty, price, amount, status)
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
