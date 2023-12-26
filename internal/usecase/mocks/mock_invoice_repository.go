// internal/usecase/mocks/mock_invoice_repository.go

package mocks

import (
	model "esb/internal/model/entity"

	"github.com/stretchr/testify/mock"
)

type MockInvoiceRepository struct {
	mock.Mock
}

func (m *MockInvoiceRepository) CreateInvoice(invoice *model.Invoice) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func (m *MockInvoiceRepository) GetAllInvoices() ([]*model.Invoice, error) {
	args := m.Called()
	return args.Get(0).([]*model.Invoice), args.Error(1)
}

func (m *MockInvoiceRepository) GetAllCustomers() ([]*model.Customer, error) {
	args := m.Called()
	return args.Get(0).([]*model.Customer), args.Error(1)
}

func (m *MockInvoiceRepository) GetAllItems() ([]*model.Items, error) {
	args := m.Called()
	return args.Get(0).([]*model.Items), args.Error(1)
}
