package main

import (
	"log"
	"net/http"

	"esb/internal/controller/invoicecontroller"
	"esb/internal/database"
	hendler "esb/internal/handler"
	"esb/internal/repository/invoicerepository"
	"esb/internal/usecase/invoiceusecase"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	invoiceRepository := invoicerepository.NewInvoiceRepository()

	// Initialize usecase and controller
	invoiceUseCase := invoiceusecase.NewInvoiceUseCase(invoiceRepository)
	invoiceController := invoicecontroller.NewInvoiceController(invoiceUseCase)

	// Initialize HTTP handler
	invoiceHandler := hendler.NewInvoiceHandler(invoiceController)

	// Setup HTTP routes
	http.HandleFunc("/api/invoices", invoiceHandler.GetAllInvoicesHandler)
	http.HandleFunc("/api/invoice/form", invoiceHandler.FormInvoiceHandler)
	http.HandleFunc("/api/invoice/create", invoiceHandler.CreateInvoiceHandler)

	// Start HTTP server
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
