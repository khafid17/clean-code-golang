package hendler

import (
	"encoding/json"
	"esb/internal/controller"
	"esb/internal/model/web"
	"net/http"
)

type InvoiceHandler struct {
	controller *controller.InvoiceController
}

func NewInvoiceHandler(controller *controller.InvoiceController) *InvoiceHandler {
	return &InvoiceHandler{controller: controller}
}

func (h *InvoiceHandler) CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var request web.InvoiceRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Convert items to a JSON string
	itemsJSON, err := json.Marshal(request.Items)
	if err != nil {
		http.Error(w, "Error convert", http.StatusInternalServerError)
		return
	}
	itemsString := string(itemsJSON)

	err = h.controller.CreateInvoice(request.Subject, request.IssueDate, request.DueDate, request.Address, request.Customer, itemsString, request.Qty, request.Price, request.Amount, request.Status)

	response := web.InvoiceResponse{
		Status:  "success",
		Message: "Invoice created successfully",
	}

	// Mengubah respons ke format JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *InvoiceHandler) GetAllInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	invoices, err := h.controller.GetAllInvoices()
	if err != nil {
		http.Error(w, "Error getting invoices", http.StatusInternalServerError)
		return
	}

	// Buat respons JSON untuk data invoices
	response := web.InvoiceGetAllResponse{
		Status:  "success",
		Message: "Invoice get all successfully",
		Data:    invoices,
	}

	// Mengubah respons ke format JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *InvoiceHandler) FormInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	customer, err := h.controller.GetAllCustomers()
	items, err := h.controller.GetAllItems()
	if err != nil {
		http.Error(w, "Error getting items and customer", http.StatusInternalServerError)
		return
	}

	// Buat respons JSON untuk data invoices
	response := web.InvoiceFormResponse{
		Status:   "success",
		Message:  "Invoice get all successfully",
		Customer: customer,
		Items:    items,
	}

	// Mengubah respons ke format JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
