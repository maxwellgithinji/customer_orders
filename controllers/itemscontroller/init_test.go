package itemscontroller

import (
	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/services/itemservice"
)

var (
	ItemTableTest      databases.ItemTable     = databases.NewItemsTable(databases.DB)
	ItemServiceTest    itemservice.ItemService = itemservice.NewItemService(ItemTableTest)
	ItemControllerTest ItemController          = NewItemController(ItemServiceTest)
)
