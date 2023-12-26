package main

import (
	"log"
	"net/http"

	"esb/internal/controller"
	"esb/internal/database"
	"esb/internal/handler"
	"esb/internal/usecase"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	// Initialize usecase and controller
	invoiceUseCase := usecase.NewInvoiceUseCase()
	invoiceController := controller.NewInvoiceController(invoiceUseCase)

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
