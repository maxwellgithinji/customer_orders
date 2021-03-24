package databases

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database interface {
	InitializeDbConnection() (*sql.DB, error)
}

type postgresql struct{}

func NewDatabase() Database {
	return &postgresql{}
}

//  A replacable connection if you use pgx v4
// func (*postgresql) InitializeDbConnection() (*pgx.Conn, error) {
// 	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Connection to db failed, Reason: %v\n", err)
// 		os.Exit(1)
// 		return nil, err
// 	}
// 	fmt.Println("Successfully connected to database")
// 	return conn, nil
// }

func (*postgresql) InitializeDbConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection to db failed, Reason: %v\n", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection to db failed, Reason: %v\n", err)
		return nil, err
	}
	fmt.Println("Successfully connected to database")
	return db, nil
}
