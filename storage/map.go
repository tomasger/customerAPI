package storage

import (
	"customerAPI/models"
	"errors"
)

type MapStorage struct {
	counter int64
	db      map[int64]*models.Customer
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		db: make(map[int64]*models.Customer),
		counter: 1,
	}
}

func (m *MapStorage) CreateCustomer(data *models.Customer) error {
	m.db[m.counter] = data
	m.counter++
	return nil
}

func (m *MapStorage) GetCustomer(id int64) (*models.Customer, error) {
	if item, exists := m.db[id]; exists {
		return item, nil
	}
	return nil, errors.New("customer not found")
}

func (m *MapStorage) GetCustomers() ([]*models.Customer, error) {
	customerCount := len(m.db)
	if customerCount == 0 {
		return nil, errors.New("database empty")
	}
	customers := make([]*models.Customer, 0, customerCount)
	for _, customer := range m.db {
		customers = append(customers, customer)
	}
	return customers, nil
}
