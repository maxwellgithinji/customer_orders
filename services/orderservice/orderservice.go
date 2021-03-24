package orderservice

import (
	"errors"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/models"
)

type OrderService interface {
	ValidateOrders(Order *models.Order) error
	CreateOrder(Order models.Order) (*models.Order, error)
	FindAllOrders() ([]models.Order, error)
	FindOrderByCustomerId(customer_id int64) ([]models.Order, error)
}

type orderservice struct{}

var (
	OrderTable databases.OrderTable = databases.NewOrdersTable(databases.DB)
)

func NewOrderService(ct databases.OrderTable) OrderService {
	OrderTable = ct
	return &orderservice{}
}

func (*orderservice) ValidateOrders(Order *models.Order) error {
	if Order == nil {
		err := errors.New("orders are empty")
		return err
	}
	return nil
}

func (*orderservice) CreateOrder(Order models.Order) (*models.Order, error) {
	return OrderTable.SaveOrder(Order)
}
func (*orderservice) FindAllOrders() ([]models.Order, error) {
	return OrderTable.FindAllOrders()
}

func (*orderservice) FindOrderByCustomerId(customer_id int64) ([]models.Order, error) {
	return OrderTable.FindOrderByCustomerId(customer_id)
}
