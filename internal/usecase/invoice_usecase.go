// usecase/invoice_usecase.go
package usecase

import (
	model "esb/internal/model/entity"
	"esb/internal/repository"
)

type InvoiceUseCase interface {
	CreateInvoice(subject string, issue_date string, due_date string, address string, customer string, items string, qty int, price float64, amount float64, status int) error
}

type invoiceUseCase struct {
}

func NewInvoiceUseCase() InvoiceUseCase {
	return &invoiceUseCase{}
}

func (uc *invoiceUseCase) CreateInvoice(subject string, issue_date string, due_date string, address string, customer string, items string, qty int, price float64, amount float64, status int) error {
	// Business logic, validation, etc.
	// You can add more complex business logic here

	// Create an Invoice object
	invoice := &model.Invoice{
		Subject:   subject,
		IssueDate: issue_date,
		DueDate:   due_date,
		Address:   address,
		Customer:  customer,
		Items:     items,
		Qty:       qty,
		Price:     price,
		Amount:    amount,
		Status:    status,
	}

	// Call repository to store the invoice in the database
	err := repository.CreateInvoice(invoice)
	if err != nil {
		return err
	}

	return nil
}

func GetAllInvoices() ([]*model.Invoice, error) {
	return repository.GetAllInvoices()
}

func GetAllCustomers() ([]*model.Customer, error) {
	return repository.GetAllCustomers()
}

func GetAllItems() ([]*model.Items, error) {
	return repository.GetAllItems()
}
