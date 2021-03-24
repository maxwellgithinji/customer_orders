package customerservice

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/models"
)

type CustomerService interface {
	ValidateCustomer(Customer *models.Customer) error
	CreateCustomer(Customer models.Customer) (*models.Customer, error)
	FindAllCustomers() ([]models.Customer, error)
	FindOneCustomer(ID int64) (*models.Customer, error)
	FindACustomerByEmail(Email string) (models.Customer, error)
	OnboardCustomer(email string, customer models.Customer) (*models.Customer, error)
	ValidateCustomerOnboarding(Customer *models.Customer) error
	ValidateCustomerOnboardingParams(Customer *models.Customer) error
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
	if Customer.Username == "" {
		err := errors.New("customers username is empty")
		return err
	}
	if Customer.Code == "" {
		err := errors.New("customers code is empty")
		return err
	}
	if Customer.Status == "" {
		err := errors.New("customers status is empty")
		return err
	}
	if Customer.Status != "" {
		if Customer.Status != "active" && Customer.Status != "inactive" {
			err := errors.New("customers status can only be active or inactive is empty")
			return err
		}
	}
	if !isEmailValid((Customer.Email)) {
		err := errors.New("customers email is invalid")
		return err
	}
	_, err := strconv.Atoi(Customer.PhoneNumber)
	if err != nil {
		err := errors.New("phone number can only be in digits")
		return err
	}
	if len(Customer.PhoneNumber) != 10 {
		err := errors.New("phone number length should be equal to 10")
		return err
	}
	return nil
}

func (*customerservice) ValidateCustomerOnboarding(Customer *models.Customer) error {
	if Customer == nil {
		err := errors.New("customers are empty")
		return err
	}
	if Customer.Username == "" {
		err := errors.New("customers username is empty")
		return err
	}
	if Customer.Code == "" {
		err := errors.New("customers code is empty")
		return err
	}
	if Customer.Status == "" {
		err := errors.New("customers status is empty")
		return err
	}
	if Customer.Status != "" {
		if Customer.Status != "active" && Customer.Status != "inactive" {
			err := errors.New("customers status can only be active or inactive is empty")
			return err
		}
	}
	_, err := strconv.Atoi(Customer.PhoneNumber)
	if err != nil {
		err := errors.New("phone number can only be in digits")
		return err
	}
	if len(Customer.PhoneNumber) != 10 {
		err := errors.New("phone number length should be equal to 10")
		return err
	}
	return nil
}

func (*customerservice) ValidateCustomerOnboardingParams(Customer *models.Customer) error {
	if Customer.Email == "" {
		err := errors.New("customers email is empty")
		return err
	}
	if !isEmailValid((Customer.Email)) {
		err := errors.New("customers email is invalid")
		return err
	}
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

func (*customerservice) OnboardCustomer(email string, customer models.Customer) (*models.Customer, error) {
	return CustomerTable.OnboardCustomer(email, customer)
}

// Helper function
func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
