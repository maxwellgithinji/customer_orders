package itemservice

import (
	"errors"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/models"
)

type ItemService interface {
	ValidateItems(Item *models.Item) error
	CreateItem(Item models.Item) (*models.Item, error)
	FindAllItems() ([]models.Item, error)
	DeleteItem(ID int64) (int64, error)
}

type itemservice struct{}

var (
	ItemTable databases.ItemTable = databases.NewItemsTable(databases.DB)
)

func NewItemService(ct databases.ItemTable) ItemService {
	ItemTable = ct
	return &itemservice{}
}

func (*itemservice) ValidateItems(Item *models.Item) error {
	if Item == nil {
		err := errors.New("items are empty")
		return err
	}
	if Item.Item == "" {
		err := errors.New("item is empty")
		return err
	}
	if Item.Price == 0 {
		err := errors.New("items price  should be greater than 0")
		return err
	}
	return nil
}

func (*itemservice) CreateItem(Item models.Item) (*models.Item, error) {
	return ItemTable.SaveItem(Item)
}
func (*itemservice) FindAllItems() ([]models.Item, error) {
	return ItemTable.FindAllItems()
}

func (*itemservice) DeleteItem(ID int64) (int64, error) {
	return ItemTable.DeleteItem(ID)
}
