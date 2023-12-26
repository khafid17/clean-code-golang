package web

import model "esb/internal/model/entity"

type InvoiceGetAllResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    []*model.Invoice `json:"data"`
}
