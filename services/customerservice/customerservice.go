package customerservice

import (
	"errors"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/models"
)

type CustomerService interface {
	ValidateCustomer(Customer *models.Customer) error
	CreateCustomer(Customer models.Customer) (*models.Customer, error)
	FindAllCustomers() ([]models.Customer, error)
	FindOneCustomer(ID int64) (*models.Customer, error)
	FindACustomerByEmail(Email string) (models.Customer, error)
}

type customerservice struct{}

var (
	CustomerTable databases.CustomerTable = databases.NewCustomersTable(databases.DB)
)

func NewCustomerService(ct databases.CustomerTable) CustomerService {
	CustomerTable = ct
	return &customerservice{}
}

func (*customerservice) ValidateCustomer(Customer *models.Customer) error {
	if Customer == nil {
		err := errors.New("customers are empty")
		return err
	}
	if Customer.Email == "" {
		err := errors.New("customers email is empty")
		return err
	}
	// TODO: Validate email
	return nil
}
func (*customerservice) CreateCustomer(Customer models.Customer) (*models.Customer, error) {
	return CustomerTable.SaveCustomer(Customer)
}
func (*customerservice) FindAllCustomers() ([]models.Customer, error) {
	return CustomerTable.FindAllCustomers()
}

func (*customerservice) FindOneCustomer(ID int64) (*models.Customer, error) {
	return CustomerTable.FindOneCustomer(ID)
}

func (*customerservice) FindACustomerByEmail(Email string) (models.Customer, error) {
	return CustomerTable.FindCustomerByEmail(Email)
}
