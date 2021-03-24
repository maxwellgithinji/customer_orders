package databases

import (
	"fmt"
	"os"
	"time"

	"github.com/maxwellgithinji/customer_orders/models"
)

type ItemTable interface {
	SaveItem(item models.Item) (*models.Item, error)
	FindAllItems() ([]models.Item, error)
	DeleteItem(ID int64) (int64, error)
}

type itemtable struct{}

var (
	ItemDb Database = NewDatabase()
)

// NewItemsTable
func NewItemsTable(db Database) ItemTable {
	ItemDb = db
	return &itemtable{}
}

func (*itemtable) FindAllItems() ([]models.Item, error) {
	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)
		//
		return nil, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		SELECT * FROM items
		ORDER BY id
	`
	var items []models.Item

	rows, err := conn.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying items table: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item

		err := rows.Scan(
			&item.ID,
			&item.Item,
			&item.Price,
			&item.CreatedAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning item rows: %v\n", err)
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (*itemtable) DeleteItem(ID int64) (int64, error) {
	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)

		return 0, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		DELETE  FROM items 
		WHERE id=$1
	`

	res, err := conn.Exec(query, ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing delete item: %v\n", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error checking delete item rows affected: %v\n", err)
		return 0, err
	}

	fmt.Println("Rows affected", rowsAffected)

	return rowsAffected, nil
}

func (*itemtable) SaveItem(item models.Item) (*models.Item, error) {
	conn, err := DB.InitializeDbConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to db: %v\n", err)

		return nil, err
	}
	defer conn.Close()
	defer fmt.Printf("Db connection closed")

	const query = `
		INSERT INTO items 
			(item, price, created_at)
		VALUES 
			($1, $2, $3)
		RETURNING *;
	`

	var id int64

	item.CreatedAt = time.Now().Local()

	err = conn.QueryRow(query,
		item.Item,
		item.Price,
		item.CreatedAt,
	).Scan(
		&id,
		&item.Item,
		&item.Price,
		&item.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving item in the db: %v\n", err)
		return nil, err
	}
	fmt.Printf("Inserted a single record %v\n", id)
	newitem := item
	newitem.ID = id
	return &newitem, nil
}
