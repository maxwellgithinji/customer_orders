package itemscontroller

import (
	"net/http"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/services/itemservice"
)

type ItemController interface {
	GetItems(w http.ResponseWriter, r *http.Request)
	CreateItem(w http.ResponseWriter, r *http.Request)
	// DeleteItem(w http.ResponseWriter, r *http.Request)
}

type itemcontroller struct{}

var (
	ItemTable   databases.ItemTable     = databases.NewItemsTable(databases.DB)
	ItemService itemservice.ItemService = itemservice.NewItemService(ItemTable)
)

func NewItemController(service itemservice.ItemService) ItemController {
	ItemService = service
	return &itemcontroller{}
}
