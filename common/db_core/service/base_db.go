package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Connection struct {
	DBType   string `mapstructure:"dbType"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

func GetConnection(dbType string,
	host string,
	port int,
	user string,
	password string,
	dbName string) Connection {
	con := Connection{}
	con.DBType = dbType
	con.Host = host
	con.Port = port
	con.User = user
	con.Password = password
	con.DBName = dbName
	return con
}

type User struct {
	name       string
	occupation string
	age        int
}

func OpenDB(connection Connection) (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connection.Host, connection.Port, connection.User, connection.Password, connection.DBName)

	db, err := sql.Open(connection.DBName, psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	return db, err

}
func CloseDB(db *sql.DB) {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateTransaction(db *sql.DB) (context.Context, *sql.Tx, error) {
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return ctx, tx, err
}

func Commit(tx *sql.Tx) {
	tx.Commit()
}

func Rollback(tx *sql.Tx) {
	tx.Rollback()
}

/*
			// `tx` is an instance of `*sql.Tx` through which we can execute our queries

	// Here, the query is executed on the transaction instance, and not applied to the database yet
	statement := `insert into "Student"("id","name") values($1, $2)`
	_, err = tx.ExecContext(ctx, statement, 1, "Adi")
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		fmt.Println("Error insert into student - " + err.Error())
		return
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err == nil {
		fmt.Println("Insert into student ended successfully")
	} else if err != nil {
		log.Fatal(err)
	}
*/
