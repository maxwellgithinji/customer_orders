package orderscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/maxwellgithinji/customer_orders/models"
	"github.com/maxwellgithinji/customer_orders/utils"
)

// GetOrders gets all orders in the database
// @Summary Get all orders in the database
// @Description Get all orders in the database
// @Tags  Orders
// @Produce  json
// @Success 200 {object} []models.Order{}
// @Router /auth/orders [get]
func (*ordercontroller) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := OrderService.FindAllOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	utils.ResponseWithDataHelper(w, "200", "orders fetch successful", orders)
}

// CreateOrder creates a new order
// @Summary creates a new order
// @Description creates a new order
// @Tags  Orders
// @Accept  json
// @Produce  json
// @Param order body models.OrderPost true "Order"
// @Success 200 {object} models.OrderPost{}
// @Router /auth/orders [post]
func (*ordercontroller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderbody models.Order

	err := json.NewDecoder(r.Body).Decode(&orderbody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	strItemId := strconv.Itoa(int(orderbody.ItemID))

	err = OrderService.ValidateOrders(&orderbody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", err.Error())
		return
	}

	err = OpenIDAuthService.InitSession()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	session, err := OpenIDAuthService.NewStore().Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	profile := session.Values["profile"]
	email := fmt.Sprintf("%v", profile.(map[string]interface{})["email"])
	defaultcustomerstate := models.Customer{}

	customer, err := CustomerService.FindACustomerByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	if customer == defaultcustomerstate {
		if err == nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseHelper(w, "400", "Customer with email "+email+" does not exist in db")
			return
		}
	}

	if customer.Status != "active" {
		if err == nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseHelper(w, "400", "Please complete the onboarding process to make an order")
			return
		}
	}

	if customer.PhoneNumber == "" {
		if err == nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseHelper(w, "400", "A phone number is required to make an order")
			return
		}
	}

	item, err := ItemService.FindOneItem(orderbody.ItemID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	if item == nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", "Item with Id "+strItemId+" does not exist in db")
		return
	}

	orderbody.CustomerID = customer.ID
	orderbody.ItemID = item.ID
	orderbody.TotalPrice = float64(item.Price)
	orderbody.OderDate = time.Now().Local()

	order, err := OrderService.CreateOrder(orderbody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	messsage := "Hi " + customer.Username + "\n Your order for " + item.Item + " was completed successfully \n Please pick your item within the next 5 working days"

	// Use wait groups to wait for the sms sending to finish
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Send message
		status, err := SMS.SendMessage(customer.PhoneNumber, messsage)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseHelper(w, "500", err.Error())
			return
		}
		if status != "201 Created" {
			w.WriteHeader(http.StatusCreated)
			utils.ResponseWithDataHelper(w, "201", "order created successfully but there was an error sending message", order)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			utils.ResponseWithDataHelper(w, "201", "order created successfully and sms sent to "+customer.PhoneNumber, order)
		}
	}()
	wg.Wait()
}

// FindCurrentUserOrders gets orders of currently logged in user
// @Summary enables the current user to get their orders
// @Description gets orders of currently logged in user
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Order{}
// @Router /auth/currentuser/orders [get]
func (*ordercontroller) FindCurrentUserOrders(w http.ResponseWriter, r *http.Request) {
	err := OpenIDAuthService.InitSession()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	session, err := OpenIDAuthService.NewStore().Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	profile := session.Values["profile"]
	email := fmt.Sprintf("%v", profile.(map[string]interface{})["email"])

	customer, err := CustomerService.FindACustomerByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", "User with email "+email+" does not exist in db")
		return
	}

	orders, err := OrderService.FindOrderByCustomerId(customer.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", err.Error())
		return
	}

	utils.ResponseWithDataHelper(w, "200", "orders for "+email+" fetched successfully", orders)
}
