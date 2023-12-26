package invoiceusecase

import (
	model "esb/internal/model/entity"
	"esb/internal/repository"
	"esb/internal/usecase"
)

type invoiceUseCaseImpl struct {
    repository repository.InvoiceRepository
}

func NewInvoiceUseCase(repository repository.InvoiceRepository) usecase.InvoiceUseCase {
    return &invoiceUseCaseImpl{
        repository: repository,
    }
}

func (uc *invoiceUseCaseImpl) CreateInvoice(subject string, issue_date string, due_date string, address string, customer string, items []model.InvoiceItem, total_items int, sub_total float64, grand_total float64, status int) error {
	// Create an Invoice object
	invoice := &model.Invoice{
		Subject:    subject,
		IssueDate:  issue_date,
		DueDate:    due_date,
		Address:    address,
		Customer:   customer,
		Items:      items,
		TotalItems: total_items,
		SubTotal:   sub_total,
		GrandTotal: grand_total,
		Status:     status,
	}

	// Call repository to store the invoice in the database
	err := uc.repository.CreateInvoice(invoice)
	if err != nil {
		return err
	}

	return nil
}

func (uc *invoiceUseCaseImpl) GetAllInvoices() ([]*model.Invoice, error) {
	return uc.repository.GetAllInvoices()
}

func (uc *invoiceUseCaseImpl) GetAllCustomers() ([]*model.Customer, error) {
	return uc.repository.GetAllCustomers()
}

func (uc *invoiceUseCaseImpl) GetAllItems() ([]*model.Items, error) {
	return uc.repository.GetAllItems()
}
