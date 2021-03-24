package customerservice

import (
	"testing"
	"time"

	"github.com/maxwellgithinji/customer_orders/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Databases
type MockDatabase struct {
	mock.Mock
}

func (mock *MockDatabase) SaveCustomer(Customer models.Customer) (*models.Customer, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(*models.Customer), args.Error(1)
}
func (mock *MockDatabase) FindAllCustomers() ([]models.Customer, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.([]models.Customer), args.Error(1)
}
func (mock *MockDatabase) FindOneCustomer(ID int64) (*models.Customer, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(*models.Customer), args.Error(1)
}
func (mock *MockDatabase) FindCustomerByEmail(Email string) (models.Customer, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(models.Customer), args.Error(1)
}

func (mock *MockDatabase) OnboardCustomer(Email string, customer models.Customer) (*models.Customer, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(*models.Customer), args.Error(1)
}

func TestValidateEmptyCustomer(t *testing.T) {
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers are empty")

	err2 := testService.ValidateCustomerOnboarding(nil)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers are empty")
}

func TestValidateEmail(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "user", Code: "123a", PhoneNumber: "0711111111", Status: "active"}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers email is invalid")

	err2 := testService.ValidateCustomerOnboardingParams(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers email is invalid")
}
func TestValidateEmailLong(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "q", Code: "123a", PhoneNumber: "0711111111", Status: "active"}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers email is invalid")

	err2 := testService.ValidateCustomerOnboardingParams(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers email is invalid")
}

func TestValidateEmailEmpty(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "", Code: "123a", PhoneNumber: "0711111111", Status: "active"}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers email is empty")

	err2 := testService.ValidateCustomerOnboardingParams(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers email is empty")
}

func TestValidateStatus(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "user@email.com", Code: "123a", PhoneNumber: "0711111111", Status: "ac"}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers status can only be active or inactive is empty")

	err2 := testService.ValidateCustomerOnboarding(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers status can only be active or inactive is empty")
}

func TestValidateStatusEmpty(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "user@email.com", Code: "123a", PhoneNumber: "0711111111", Status: ""}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers status is empty")

	err2 := testService.ValidateCustomerOnboarding(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers status is empty")
}

func TestValidateUsername(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "", Email: "user@email.com", Code: "123a", PhoneNumber: "0711111111", Status: ""}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers username is empty")

	err2 := testService.ValidateCustomerOnboarding(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers username is empty")
}
func TestValidateCode(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "uu", Email: "user@email.com", Code: "", PhoneNumber: "0711111111", Status: ""}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers code is empty")

	err2 := testService.ValidateCustomerOnboarding(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "customers code is empty")
}

func TestValidatePhoneNumber(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "uu", Email: "user@email.com", Code: "ds", PhoneNumber: "071111ui11", Status: "active"}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "phone number can only be in digits")

	err2 := testService.ValidateCustomerOnboarding(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "phone number can only be in digits")
}

func TestValidatePhoneNumberLen(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "uu", Email: "user@email.com", Code: "ds", PhoneNumber: "07111111", Status: "active"}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "phone number length should be equal to 10")

	err2 := testService.ValidateCustomerOnboarding(&customer)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "phone number length should be equal to 10")
}

func TestValidateSuccess(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a", PhoneNumber: "0711111111", Status: "active"}
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(&customer)
	assert.Nil(t, err)

	err2 := testService.ValidateCustomerOnboardingParams(&customer)
	assert.Nil(t, err2)

	err3 := testService.ValidateCustomerOnboarding(&customer)
	assert.Nil(t, err3)
}

func TestFindAllCustomers(t *testing.T) {
	mockDb := new(MockDatabase)
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a"}

	// Setup expectations
	mockDb.On("FindAllCustomers").Return([]models.Customer{customer}, nil)

	testService := NewCustomerService(mockDb)

	res, _ := testService.FindAllCustomers()

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, int64(1), res[0].ID)
	assert.Equal(t, "maxgit", res[0].Username)
	assert.Equal(t, "maxwellgithinji@gmail.com", res[0].Email)
	assert.Equal(t, "123a", res[0].Code)

}

func TestCreateACustomer(t *testing.T) {
	mockDb := new(MockDatabase)
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a", PhoneNumber: "0711111111", CreatedAt: time.Now().Local()}

	// Setup expectations
	mockDb.On("SaveCustomer").Return(&customer, nil)

	testService := NewCustomerService(mockDb)
	res, err := testService.CreateCustomer(customer)

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, res.ID)
	assert.Equal(t, "maxgit", res.Username)
	assert.Equal(t, "maxwellgithinji@gmail.com", res.Email)
	assert.Equal(t, "123a", res.Code)
	assert.Nil(t, err)
}

func TestFindACustomerByEmail(t *testing.T) {
	mockDb := new(MockDatabase)
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a", PhoneNumber: "0711111111", CreatedAt: time.Now().Local()}

	// Setup expectations
	testService := NewCustomerService(mockDb)
	mockDb.On("FindCustomerByEmail").Return(customer, nil)

	res2, err := testService.FindACustomerByEmail(customer.Email)
	// Mock Assertion : Behavioral
	assert.NotNil(t, res2.ID)
	assert.Equal(t, "maxgit", res2.Username)
	assert.Equal(t, "maxwellgithinji@gmail.com", res2.Email)
	assert.Equal(t, "123a", res2.Code)
	assert.Nil(t, err)
}

func TestFindACustomerByID(t *testing.T) {
	mockDb := new(MockDatabase)
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a", PhoneNumber: "0711111111", CreatedAt: time.Now().Local()}

	// Setup expectations
	testService := NewCustomerService(mockDb)
	mockDb.On("FindOneCustomer").Return(&customer, nil)

	res2, err := testService.FindOneCustomer(customer.ID)
	// Mock Assertion : Behavioral
	assert.NotNil(t, res2.ID)
	assert.Equal(t, "maxgit", res2.Username)
	assert.Equal(t, "maxwellgithinji@gmail.com", res2.Email)
	assert.Equal(t, "123a", res2.Code)
	assert.Nil(t, err)
}

func TestOnboardACustomer(t *testing.T) {
	mockDb := new(MockDatabase)
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com"}
	updatecustomer := models.Customer{Username: "maxgit2", Code: "123a", PhoneNumber: "0711111111", Status: "active"}
	// Setup expectations
	mockDb.On("SaveCustomer").Return(&customer, nil)

	testService := NewCustomerService(mockDb)
	res, err := testService.CreateCustomer(customer)

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, res.ID)
	assert.Equal(t, "maxgit", res.Username)
	assert.Equal(t, "maxwellgithinji@gmail.com", res.Email)
	assert.Nil(t, err)

	// Update the customer
	mockDb.On("OnboardCustomer").Return(&updatecustomer, nil)
	res2, err := testService.OnboardCustomer(customer.Email, updatecustomer)
	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)
	assert.NotNil(t, res2.ID)
	assert.Equal(t, "maxgit2", res2.Username)
	assert.Equal(t, "123a", res2.Code)
	assert.Equal(t, "0711111111", res2.PhoneNumber)
	assert.Equal(t, "active", res2.Status)
	assert.Nil(t, err)
}
