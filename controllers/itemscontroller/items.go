package itemscontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/customer_orders/models"
	"github.com/maxwellgithinji/customer_orders/utils"
)

// GetItems gets all items in the database
// @Summary Get all items in the database
// @Description Get all items in the database
// @Tags  Items
// @Produce  json
// @Success 200 {object} []models.Item{}
// @Router /items [get]
func (*itemcontroller) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := ItemService.FindAllItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	utils.ResponseWithDataHelper(w, "200", "items fetch successful", items)
}

// CreateItem creates a new item
// @Summary Get profile creates a new item
// @Description Get profile creates a new item
// @Tags  Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemPost true "Item"
// @Success 200 {object} models.ItemPost{}
// @Router /auth/item [post]
func (*itemcontroller) CreateItem(w http.ResponseWriter, r *http.Request) {
	var itembody models.Item

	err := json.NewDecoder(r.Body).Decode(&itembody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	err = ItemService.ValidateItems(&itembody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", err.Error())
		return
	}

	item, err := ItemService.CreateItem(itembody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	utils.ResponseWithDataHelper(w, "200", "item details updated successfully", item)
}

// DeleteItem deletes an item
// @Summary enables a user to delete an item
// @Description deletes an item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path int64 true "Item Id"
// @Success 200
// @Router /auth/delete/item/{id} [delete]
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strID := params["id"]

	intId, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", err.Error())
		return
	}
	int64Id := int64(intId)

	item, err := ItemService.DeleteItem(int64Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", err.Error())
		return
	}
	if item == 0 {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", "Item with id '"+strID+"' not found")
		return
	}
	utils.ResponseWithDataHelper(w, "200", "item deleted successfully", item)
}
