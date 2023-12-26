package hendler

import (
	"encoding/json"
	"esb/internal/controller"
	model "esb/internal/model/entity"
	"esb/internal/model/web"
	"net/http"
)

type InvoiceHandler struct {
	controller controller.InvoiceController
}

func NewInvoiceHandler(controller controller.InvoiceController) *InvoiceHandler {
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

	var subTotal, grandTotal float64
	var items []model.InvoiceItem
	for _, itemReq := range request.Items {
		amount := itemReq.Qty * itemReq.Price
		item := model.InvoiceItem{
			Qty:    itemReq.Qty,
			Price:  itemReq.Price,
			Amount: amount,
		}
		items = append(items, item)
		subTotal += amount
	}

	tax := (request.Items[0].Tax / 100) * subTotal
	grandTotal = tax + subTotal

	err := h.controller.CreateInvoice(request.Subject, request.IssueDate, request.DueDate, request.Address, request.Customer, items, len(items), subTotal, grandTotal, request.Status)
	if err != nil {
		http.Error(w, "Error create data", http.StatusInternalServerError)
		return
	}

	response := web.InvoiceResponse{
		Status:  "success",
		Message: "Invoice created successfully",
		InvoiceSummary: web.InvoiceSummary{
			TotalItems: len(items),
			SubTotal:   subTotal,
			Tax:        request.Items[0].Tax,
			GrandTotal: grandTotal,
		},
	}

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

	response := web.InvoiceGetAllResponse{
		Status:  "success",
		Message: "Invoice get all successfully",
		Data:    invoices,
	}

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

	response := web.InvoiceFormResponse{
		Status:   "success",
		Message:  "Invoice get all successfully",
		Customer: customer,
		Items:    items,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
