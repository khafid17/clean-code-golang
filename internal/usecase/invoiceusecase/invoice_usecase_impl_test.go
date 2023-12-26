package invoiceusecase

import (
	model "esb/internal/model/entity"
	"esb/internal/model/web"
	"esb/internal/usecase/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoice(t *testing.T) {
	mockRepo := new(mocks.MockInvoiceRepository)
	mockRepo.On("CreateInvoice", mock.Anything).Return(nil)
	useCase := NewInvoiceUseCase(mockRepo)

	// invoice request
	request := &web.InvoiceRequest{
		Subject:    "Test Invoice",
		IssueDate:  "2023-01-01",
		DueDate:    "2023-02-01",
		Address:    "Test Address",
		Customer:   "Test Customer",
		Items:      []model.InvoiceItem{{Qty: 2, Price: 10, Tax: 5}},
		TotalItems: 1,
		Status:     1,
	}

	var subTotal, grandTotal float64
	for _, item := range request.Items {
		subTotal += item.Qty * item.Price
	}
	grandTotal = subTotal + (request.Items[0].Tax/100)*subTotal

	// usecase
	err := useCase.CreateInvoice(
		request.Subject,
		request.IssueDate,
		request.DueDate,
		request.Address,
		request.Customer,
		request.Items,
		request.TotalItems,
		subTotal,   
		grandTotal, 
		request.Status,
	)

	// Assertions
	assert.NoError(t, err, "Expected no error")
	mockRepo.AssertExpectations(t)
}

func TestGetAllInvoices(t *testing.T) {
	mockRepo := new(mocks.MockInvoiceRepository)
	mockUseCase := NewInvoiceUseCase(mockRepo)

	// expected data from the repository
	expectedInvoices := []*model.Invoice{
		{
			ID:         19,
			Subject:    "Invoice Subject",
			IssueDate:  "2023-01-01 00:00:00",
			DueDate:    "2023-01-15 00:00:00",
			Address:    "Invoice Address",
			Customer:   "Customer Company",
			Items:      []model.InvoiceItem{{Qty: 2, Price: 10, Tax: 5}},
			TotalItems: 2,
			SubTotal:   190,
			GrandTotal: 209,
			Status:     0,
		},
		{
			ID:         20,
			Subject:    "Invoice Subject",
			IssueDate:  "2023-01-01 00:00:00",
			DueDate:    "2023-01-15 00:00:00",
			Address:    "Invoice Address",
			Customer:   "Customer Company",
			Items:      []model.InvoiceItem{{Qty: 2, Price: 10, Tax: 5}},
			TotalItems: 2,
			SubTotal:   190,
			GrandTotal: 209,
			Status:     0,
		},
	}

	// Mock the GetAllInvoices method in the repository
	mockRepo.On("GetAllInvoices").Return(expectedInvoices, nil)
	actualInvoices, err := mockUseCase.GetAllInvoices()

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedInvoices, actualInvoices, "Invoices do not match")
	mockRepo.AssertExpectations(t)
}

func TestGetAllCustomers(t *testing.T) {
	mockRepo := new(mocks.MockInvoiceRepository)
	mockUseCase := NewInvoiceUseCase(mockRepo)

	// expected data from the repository
	expectedCustomers := []*model.Customer{
		{ID: 1, Company: "Customer 1",},
		{ID: 2, Company: "Customer 2",},
	}

	// Mock the GetAllCustomers method in the repository
	mockRepo.On("GetAllCustomers").Return(expectedCustomers, nil)
	actualCustomers, err := mockUseCase.GetAllCustomers()

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedCustomers, actualCustomers, "Customers do not match")
	mockRepo.AssertExpectations(t)
}

func TestGetAllItems(t *testing.T) {
	mockRepo := new(mocks.MockInvoiceRepository)
	mockUseCase := NewInvoiceUseCase(mockRepo)

	// expected data from the repository
	expectedItems := []*model.Items{
		{ID: 1, ItemName: "Item 1", Type: "Service"},
		{ID: 2, ItemName: "Item 2", Type: "Hardware"},
	}

	// Mock the GetAllItems method in the repository
	mockRepo.On("GetAllItems").Return(expectedItems, nil)
	actualItems, err := mockUseCase.GetAllItems()

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedItems, actualItems, "Items do not match")
	mockRepo.AssertExpectations(t)
}