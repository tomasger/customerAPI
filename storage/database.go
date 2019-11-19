package storage

import "customerAPI/models"

type Database interface {
	CreateCustomer(data *models.Customer) error
	GetCustomer(id int64) (*models.Customer, error)
	GetCustomers() ([]*models.Customer, error)
}
