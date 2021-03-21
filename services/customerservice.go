package service

import (
	"errors"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/models"
)

type CustomerService interface {
	ValidateCustomer(Customer *models.Customer) error
	CreateCustomer(Customer *models.Customer) (*models.Customer, error)
	FindAllCustomers() ([]models.Customer, error)
	FindOneCustomer(ID int64) (*models.Customer, error)
}

type customerservice struct{}

var (
	CustomerTable databases.CustomerTable
)

func NewCustomerService(ct databases.CustomerTable) CustomerService {
	CustomerTable = ct
	return &customerservice{}
}

func (*customerservice) ValidateCustomer(Customer *models.Customer) error {
	if Customer == nil {
		err := errors.New("Customers are empty")
		return err
	}
	if Customer.Name == "" {
		err := errors.New("Customers name is empty")
		return err
	}
	return nil
}
func (*customerservice) CreateCustomer(Customer *models.Customer) (*models.Customer, error) {
	return CustomerTable.SaveCustomer(Customer)
}
func (*customerservice) FindAllCustomers() ([]models.Customer, error) {
	return CustomerTable.FindAllCustomers()
}

func (*customerservice) FindOneCustomer(ID int64) (*models.Customer, error) {
	return CustomerTable.FindOneCustomer(ID)
}
