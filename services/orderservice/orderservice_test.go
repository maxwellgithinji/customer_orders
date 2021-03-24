package orderservice

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

func (mock *MockDatabase) SaveOrder(Order models.Order) (*models.Order, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(*models.Order), args.Error(1)
}
func (mock *MockDatabase) FindAllOrders() ([]models.Order, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.([]models.Order), args.Error(1)
}
func (mock *MockDatabase) FindOrderByCustomerId(customer_id int64) ([]models.Order, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.([]models.Order), args.Error(1)
}

func TestValidateEmptyOrder(t *testing.T) {
	testService := NewOrderService(nil)

	err := testService.ValidateOrders(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "orders are empty")
}

func TestValidateSuccess(t *testing.T) {
	order := models.Order{
		ID:         1,
		CustomerID: 1,
		ItemID:     1,
		TotalPrice: 100,
		OderDate:   time.Now().Local(),
		CreatedAt:  time.Now().Local(),
	}
	testService := NewOrderService(nil)

	err := testService.ValidateOrders(&order)
	assert.Nil(t, err)
}

func TestFindAllOrders(t *testing.T) {
	mockDb := new(MockDatabase)
	order := models.Order{
		ID:         1,
		CustomerID: 1,
		ItemID:     1,
		TotalPrice: 0,
		OderDate:   time.Now().Local(),
		CreatedAt:  time.Now().Local(),
	}

	// Setup expectations
	mockDb.On("FindAllOrders").Return([]models.Order{order}, nil)

	testService := NewOrderService(mockDb)
	res, err := testService.FindAllOrders()

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, int64(1), res[0].ID)
	assert.Equal(t, int64(1), res[0].CustomerID)
	assert.Equal(t, int64(1), res[0].ItemID)
	assert.NotNil(t, res[0].OderDate)
	assert.NotNil(t, res[0].CreatedAt)
	assert.Nil(t, err)
}

func TestCreateAnOrder(t *testing.T) {
	mockDb := new(MockDatabase)
	order := models.Order{
		ID:         1,
		CustomerID: 1,
		ItemID:     1,
		TotalPrice: 0,
		OderDate:   time.Now().Local(),
		CreatedAt:  time.Now().Local(),
	}

	// Setup expectations
	mockDb.On("SaveOrder").Return(&order, nil)

	testService := NewOrderService(mockDb)
	res, err := testService.CreateOrder(order)

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, int64(1), res.ID)
	assert.Equal(t, int64(1), res.CustomerID)
	assert.Equal(t, int64(1), res.ItemID)
	assert.NotNil(t, res.OderDate)
	assert.NotNil(t, res.CreatedAt)
	assert.Nil(t, err)
}

func TestFindOrderByCustomerId(t *testing.T) {
	mockDb := new(MockDatabase)
	order := models.Order{
		ID:         1,
		CustomerID: 1,
		ItemID:     1,
		TotalPrice: 0,
		OderDate:   time.Now().Local(),
		CreatedAt:  time.Now().Local(),
	}
	customer := models.Customer{ID: 1, Username: "maxgit", Email: "maxwellgithinji@gmail.com", Code: "123a", PhoneNumber: "0711111111", CreatedAt: time.Now().Local()}

	// Setup expectations
	testService := NewOrderService(mockDb)
	mockDb.On("FindOrderByCustomerId").Return([]models.Order{order}, nil)

	res, err := testService.FindOrderByCustomerId(customer.ID)
	// Mock Assertion : Behavioral
	assert.Equal(t, int64(1), res[0].ID)
	assert.Equal(t, int64(1), res[0].CustomerID)
	assert.Equal(t, int64(1), res[0].ItemID)
	assert.NotNil(t, res[0].OderDate)
	assert.NotNil(t, res[0].CreatedAt)
	assert.Nil(t, err)
}
