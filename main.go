package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db := CreateConnection()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	//insert(db)
	// Execute the query
	results, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}

}

func CreateConnection() *sql.DB {
	dataSourceName := GetMysqlConnectionURI()
	db, err := sql.Open("mysql", dataSourceName)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

type TagDAO struct {
	conn *sql.DB
}


func (dao TagDAO) PrepareInsert(tagName string) (err error) {
	stmtIns, err := dao.conn.Prepare("INSERT INTO tags (name) VALUES (?)")
	if err != nil {
		log.Println(err)
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(tagName)
	return err
}

func (dao TagDAO) PrepareSelect() ([]Tag, error) {
	rows, err := dao.conn.Query("SELECT id, name FROM tags")
	if err != nil {
		return nil, err
	}

	fmt.Println(rows.Columns())
	data := make([]Tag, 0)
	for rows.Next() {
		var tag Tag
		err = rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			log.Println(err)
		} else {
			data = append(data, tag)
		}

	}
	return data, nil
}

func insert(db *sql.DB) {
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO tags VALUES (2, 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}