package databases

import "github.com/maxwellgithinji/customer_orders/models"

type CustomerTable interface {
	SaveCustomer(customer *models.Customer) (*models.Customer, error)
	FindAllCustomers() ([]models.Customer, error)
	FindOneCustomer(ID int64) (*models.Customer, error)
	FindCustomerByEmail(Email string) (*models.Customer, error)
}

type customertable struct{}

// NewCustomersTable
func NewCustomersTable() CustomerTable {
	return &customertable{}
}

func (*customertable) FindAllCustomers() ([]models.Customer, error) {
	return nil, nil
}
func (*customertable) FindOneCustomer(ID int64) (*models.Customer, error) {
	return nil, nil
}

func (*customertable) SaveCustomer(customer *models.Customer) (*models.Customer, error) {
	return nil, nil
}

func (*customertable) FindCustomerByEmail(Email string) (*models.Customer, error) {
	return nil, nil
}
