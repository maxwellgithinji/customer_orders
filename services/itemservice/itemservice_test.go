package itemservice

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

func (mock *MockDatabase) SaveItem(Item models.Item) (*models.Item, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(*models.Item), args.Error(1)
}
func (mock *MockDatabase) FindAllItems() ([]models.Item, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.([]models.Item), args.Error(1)
}
func (mock *MockDatabase) DeleteItem(ID int64) (int64, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(int64), args.Error(1)
}

func (mock *MockDatabase) FindOneItem(ID int64) (*models.Item, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(*models.Item), args.Error(1)
}

func TestValidateEmptyItem(t *testing.T) {
	testService := NewItemService(nil)

	err := testService.ValidateItems(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "items are empty")
}

func TestValidateItemItem(t *testing.T) {
	item := models.Item{ID: 1, Item: "", Price: 100}
	testService := NewItemService(nil)

	err := testService.ValidateItems(&item)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "item is empty")

}

func TestValidateprice(t *testing.T) {
	item := models.Item{ID: 1, Item: "itemm 2", Price: 0}
	testService := NewItemService(nil)

	err := testService.ValidateItems(&item)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "items price  should be greater than 0")
}

func TestValidateSuccess(t *testing.T) {
	item := models.Item{ID: 1, Item: "item 2", Price: 100}
	testService := NewItemService(nil)

	err := testService.ValidateItems(&item)
	assert.Nil(t, err)
}

func TestFindAllItems(t *testing.T) {
	mockDb := new(MockDatabase)
	item := models.Item{ID: 1, Item: "item 2", Price: 100, CreatedAt: time.Now().Local()}

	// Setup expectations
	mockDb.On("FindAllItems").Return([]models.Item{item}, nil)

	testService := NewItemService(mockDb)

	res, _ := testService.FindAllItems()

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, int64(1), res[0].ID)
	assert.Equal(t, "item 2", res[0].Item)
	assert.Equal(t, int64(100), res[0].Price)
}

func TestCreateAnItem(t *testing.T) {
	mockDb := new(MockDatabase)
	item := models.Item{ID: 1, Item: "item 2", Price: 100, CreatedAt: time.Now().Local()}

	// Setup expectations
	mockDb.On("SaveItem").Return(&item, nil)

	testService := NewItemService(mockDb)
	res, err := testService.CreateItem(item)

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, res.ID)
	assert.Equal(t, int64(1), res.ID)
	assert.Equal(t, "item 2", res.Item)
	assert.Equal(t, int64(100), res.Price)
	assert.Nil(t, err)
}

func TestDeleteAnItem(t *testing.T) {
	mockDb := new(MockDatabase)
	item := models.Item{ID: 1, Item: "item 2", Price: 100, CreatedAt: time.Now().Local()}
	ret := int64(1)
	// Setup expectations
	testService := NewItemService(mockDb)
	mockDb.On("DeleteItem").Return(ret, nil)

	res2, err := testService.DeleteItem(item.ID)
	// Mock Assertion : Behavioral

	assert.Equal(t, int64(1), res2)
	assert.Nil(t, err)
}

func TestFindOneItem(t *testing.T) {
	mockDb := new(MockDatabase)
	item := models.Item{ID: 1, Item: "item 2", Price: 100, CreatedAt: time.Now().Local()}

	// Setup expectations
	mockDb.On("FindOneItem").Return(&item, nil)

	testService := NewItemService(mockDb)
	res, err := testService.FindOneItem(item.ID)

	// Mock Assertion : Behavioral
	mockDb.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, res.ID)
	assert.Equal(t, int64(1), res.ID)
	assert.Equal(t, "item 2", res.Item)
	assert.Equal(t, int64(100), res.Price)
	assert.Nil(t, err)
}
