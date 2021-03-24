package databases

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/maxwellgithinji/customer_orders/models"
)

type CustomerTable interface {
	SaveCustomer(customer models.Customer) (*models.Customer, error)
	FindAllCustomers() ([]models.Customer, error)
	FindOneCustomer(ID int64) (*models.Customer, error)
	FindCustomerByEmail(Email string) (models.Customer, error)
}

type customertable struct{}

var (
	DB Database = NewDatabase()
)

// NewCustomersTable
func NewCustomersTable(db Database) CustomerTable {
	DB = db
	return &customertable{}
}

func (*customertable) FindAllCustomers() ([]models.Customer, error) {
	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)
		//
		return nil, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		SELECT * FROM customers
		ORDER BY id
	`
	var customers []models.Customer

	rows, err := conn.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying customers table: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Username,
			&customer.Email,
			&customer.PhoneNumber,
			&customer.Code,
			&customer.Status,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning customer rows: %v\n", err)
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
func (*customertable) FindOneCustomer(ID int64) (*models.Customer, error) {
	return nil, nil
}

func (*customertable) SaveCustomer(customer models.Customer) (*models.Customer, error) {
	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)

		return nil, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		INSERT INTO customers 
			(username, email, phone_number, code, status, created_at)
		VALUES 
			($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`

	var id int64

	customer.CreatedAt = time.Now().Local()

	err = conn.QueryRow(query,
		customer.Username,
		customer.Email,
		customer.PhoneNumber,
		customer.Code,
		customer.Status,
		customer.CreatedAt,
	).Scan(
		&id,
		&customer.Username,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.Code,
		&customer.Status,
		&customer.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving customer in the db: %v\n", err)
		return nil, err
	}
	fmt.Printf("Inserted a single record %v\n", id)
	newcustomer := customer
	newcustomer.ID = id
	return &newcustomer, nil
}

func (*customertable) FindCustomerByEmail(Email string) (models.Customer, error) {
	var customer models.Customer

	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)

		return customer, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		SELECT * FROM customers 
		WHERE email=$1
		LIMIT $2
	`
	limit := 1

	row := conn.QueryRow(query, Email, limit)

	err = row.Scan(
		&customer.ID,
		&customer.Username,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.Code,
		&customer.Status,
		&customer.CreatedAt,
	)
	switch err {
	case sql.ErrNoRows:
		fmt.Printf("No rows were returned")
		return customer, nil
	case nil:
		return customer, nil
	default:
		fmt.Fprintf(os.Stderr, "Unable to scan customer rows: %v\n", err)
	}
	return customer, err
}
