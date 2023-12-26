package invoicerepository

import (
	"log"

	"esb/internal/database"
	"esb/internal/repository"

	model "esb/internal/model/entity"
)


type invoiceRepositoryImpl struct {}

func NewInvoiceRepository() repository.InvoiceRepository {
    return &invoiceRepositoryImpl{}
}


func (r *invoiceRepositoryImpl) CreateInvoice(invoice *model.Invoice) error {
    _, err := database.DB.Exec("INSERT INTO invoices (subject, issue_date, due_date, address, customer, total_items, sub_total, grand_total, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		invoice.Subject, invoice.IssueDate, invoice.DueDate, invoice.Address, invoice.Customer, invoice.TotalItems, invoice.SubTotal, invoice.GrandTotal, invoice.Status)

	if err != nil {
		log.Println("Error creating invoice:", err)
		return err
	}

	// Get the last inserted ID
	var lastID int
	err = database.DB.QueryRow("SELECT LAST_INSERT_ID()").Scan(&lastID)
	if err != nil {
		log.Println("Error getting last inserted ID:", err)
		return err
	}

	// Insert items related to the invoice
	for _, item := range invoice.Items {
		_, err := database.DB.Exec("INSERT INTO invoice_items (invoice_id, qty, price, amount) VALUES (?, ?, ?, ?)",
			lastID, item.Qty, item.Price, item.Amount)

		if err != nil {
			log.Println("Error creating invoice item:", err)
			return err
		}
	}

	log.Println("Invoice created successfully")
	return nil
}

func (r *invoiceRepositoryImpl) GetAllInvoices() ([]*model.Invoice, error) {
    rows, err := database.DB.Query("SELECT id, subject, issue_date, due_date, address, customer, total_items, sub_total, grand_total, status FROM invoices")
	if err != nil {
		log.Println("Error querying invoices:", err)
		return nil, err
	}
	defer rows.Close()

	var invoices []*model.Invoice
	for rows.Next() {
		var invoice model.Invoice
		if err := rows.Scan(&invoice.ID, &invoice.Subject, &invoice.IssueDate, &invoice.DueDate, &invoice.Address, &invoice.Customer, &invoice.TotalItems, &invoice.SubTotal, &invoice.GrandTotal, &invoice.Status); err != nil {
			log.Println("Error scanning invoice row:", err)
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func (r *invoiceRepositoryImpl) GetAllCustomers() ([]*model.Customer, error) {
    rows, err := database.DB.Query("SELECT id, company FROM customers")
	if err != nil {
		log.Println("Error querying customers:", err)
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.ID, &customer.Company); err != nil {
			log.Println("Error scanning invoice row:", err)
			return nil, err
		}
		customers = append(customers, &customer)
	}

	return customers, nil
}

func (r *invoiceRepositoryImpl) GetAllItems() ([]*model.Items, error) {
    rows, err := database.DB.Query("SELECT id, item_name, type FROM items")
	if err != nil {
		log.Println("Error querying items:", err)
		return nil, err
	}
	defer rows.Close()

	var items []*model.Items
	for rows.Next() {
		var item model.Items
		if err := rows.Scan(&item.ID, &item.ItemName, &item.Type); err != nil {
			log.Println("Error scanning invoice row:", err)
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
