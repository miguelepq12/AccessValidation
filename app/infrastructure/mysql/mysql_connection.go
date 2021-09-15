package oracle

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnection struct {
	db *sql.DB
}

const(
	USERNAME = "root"
	PASSWORD = "admin"
	HOST = "127.0.0.1:3306"
	DATABASE = "my_dbs"
)

func NewMysqlConnection() (*MysqlConnection, error){
	fmt.Println("... Setting up Database Connection")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",USERNAME,PASSWORD,HOST,DATABASE))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("... Opening Database Connection")
	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n\n", err)
		return nil, err
	}
	fmt.Println("... Connected to Database")

	return &MysqlConnection{db: db}, nil
}

func (c MysqlConnection) Close()  {
	err := c.db.Close()
	if err != nil {
		return
	}
}
