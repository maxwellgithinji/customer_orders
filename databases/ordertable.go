package databases

import (
	"fmt"
	"os"
	"time"

	"github.com/maxwellgithinji/customer_orders/models"
)

type OrderTable interface {
	SaveOrder(order models.Order) (*models.Order, error)
	FindAllOrders() ([]models.Order, error)
	FindOrderByCustomerId(customer_id int64) ([]models.Order, error)
}

type ordertable struct{}

var (
	OrderDb Database = NewDatabase()
)

// NewOrdersTable
func NewOrdersTable(db Database) OrderTable {
	OrderDb = db
	return &ordertable{}
}

func (*ordertable) FindAllOrders() ([]models.Order, error) {
	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)
		//
		return nil, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		SELECT * FROM orders
		ORDER BY id
	`
	var orders []models.Order

	rows, err := conn.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying orders table: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order

		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.ItemID,
			&order.TotalPrice,
			&order.OderDate,
			&order.CreatedAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning order rows: %v\n", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (*ordertable) SaveOrder(order models.Order) (*models.Order, error) {
	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)

		return nil, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		INSERT INTO orders 
			(customer_id, item_id, total_price, order_date, created_at)
		VALUES 
			($1, $2, $3, $4, $5)
		RETURNING *;
	`

	var id int64

	order.OderDate = time.Now().Local()
	order.CreatedAt = time.Now().Local()

	err = conn.QueryRow(query,
		order.CustomerID,
		order.ItemID,
		order.TotalPrice,
		order.OderDate,
		order.CreatedAt,
	).Scan(
		&id,
		&order.CustomerID,
		&order.ItemID,
		&order.TotalPrice,
		&order.OderDate,
		&order.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving order in the db: %v\n", err)
		return nil, err
	}
	fmt.Printf("Inserted a single record %v\n", id)
	neworder := order
	neworder.ID = id
	return &neworder, nil
}

func (*ordertable) FindOrderByCustomerId(customer_id int64) ([]models.Order, error) {

	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)

		return nil, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		SELECT * FROM orders 
		WHERE customer_id=$1
	`

	var orders []models.Order

	rows, err := conn.Query(query, customer_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying orders table fetch by customer id: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order

		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.ItemID,
			&order.TotalPrice,
			&order.OderDate,
			&order.CreatedAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning order rows: %v\n", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
