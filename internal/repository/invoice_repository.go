package repository

import (
	"log"

	"esb/internal/database"

	model "esb/internal/model/entity"
)

func CreateInvoice(invoice *model.Invoice) error {
	_, err := database.DB.Exec("INSERT INTO invoice (subject, issue_date, due_date, address, customer, items, qty, price, amount, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		invoice.Subject, invoice.IssueDate, invoice.DueDate, invoice.Address, invoice.Customer, invoice.Items, invoice.Qty, invoice.Price, invoice.Amount, invoice.Status)

	if err != nil {
		log.Println("Error creating invoice:", err)
		return err
	}

	log.Println("Invoice created successfully")
	return nil
}

func GetAllInvoices() ([]*model.Invoice, error) {
	rows, err := database.DB.Query("SELECT id, subject, issue_date, due_date, address, customer, items, qty, price, amount, status FROM invoice")
	if err != nil {
		log.Println("Error querying invoices:", err)
		return nil, err
	}
	defer rows.Close()

	var invoices []*model.Invoice
	for rows.Next() {
		var invoice model.Invoice
		if err := rows.Scan(&invoice.ID, &invoice.Subject, &invoice.IssueDate, &invoice.DueDate, &invoice.Address, &invoice.Customer, &invoice.Items, &invoice.Qty, &invoice.Price, &invoice.Amount, &invoice.Status); err != nil {
			log.Println("Error scanning invoice row:", err)
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func GetAllCustomers() ([]*model.Customer, error) {
	rows, err := database.DB.Query("SELECT id, company FROM customer")
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

func GetAllItems() ([]*model.Items, error) {
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
