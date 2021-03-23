package customerservice

import (
	"testing"

	"github.com/maxwellgithinji/customer_orders/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Databases
type MockDatabase struct {
	mock.Mock
}

func (mock *MockDatabase) SaveCustomer(Customer models.Customer) (models.Customer, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(models.Customer), args.Error(1)
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

func TestValidateEmptyCustomer(t *testing.T) {
	testService := NewCustomerService(nil)

	err := testService.ValidateCustomer(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers are empty")
}

func TestValidateEmptyEmail(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "", Code: "123a", PhoneNumber: "0711111111"}
	testService := NewCustomerService(nil)
	err := testService.ValidateCustomer(&customer)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "customers email is empty")
}

func TestValidateSuccess(t *testing.T) {
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a", PhoneNumber: "0711111111"}
	testService := NewCustomerService(nil)
	err := testService.ValidateCustomer(&customer)

	assert.Nil(t, err)
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
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a", PhoneNumber: "0711111111"}

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
