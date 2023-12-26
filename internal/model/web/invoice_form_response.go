package web

import model "esb/internal/model/entity"

type InvoiceFormResponse struct {
	Status   string            `json:"status"`
	Message  string            `json:"message"`
	Customer []*model.Customer `json:"customers"`
	Items    []*model.Items    `json:"items"`
}
