/////////////////////////////////////////////////////////
// sqlite3 Sample demo using mattn driver at github.com/mattn/go-sqlite3
//
//
// IMPORTAN be sure to run first: go get github.com/mattn/go-sqlite3
//
// For full details visit: https://github.com/mattn/go-sqlite3
// to install this package inthe current directory
///////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"strconv"
)

type Record struct {
	Id	string
	Name	string
	Phone	string
}

func main(){
	dbname := "sqlite3.db"
	var db_handle *sql.DB
	var names = [5]string{"john", "sally", "andy", "Becky", "Donna"}


	fmt.Println("----- SQLite3 example -----")

    fmt.Println("1. Init the DB: ",dbname)
	db_handle= InitDB(dbname)
	fmt.Println("2. Creating Table: ",dbname)
	CreateTable(db_handle)

	fmt.Println("3. Inerting row: ",dbname)
	
	recs := []Record{}  //blank array of struct

	for i:=0; i< 1000; i++ {
	digit:=(rand.Intn(len(names) - 0) + 0)
	digits :=strconv.Itoa(i)
	r1 :=  Record{Id: digits , Name: names[digit], Phone:"123-456-7890"}
	recs = append(recs,r1)
	}

	//fmt.Println("Created row item 3: ",recs )
	//os.Exit(0)
	StoreItem(db_handle, recs)

	fmt.Println("4. Reading Table: ")
	db_row_read := ReadItem(db_handle)

	fmt.Println("5. Row Read from  sQlite3: ",db_row_read)

	fmt.Println("done")
	os.Exit(0)

}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		Phone TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }
}

func StoreItem(db *sql.DB, items []Record) {
	sql_additem := `
	INSERT OR REPLACE INTO items(
		Id,
		Name,
		Phone,
		InsertedDatetime
	) values(?, ?, ?, CURRENT_TIMESTAMP)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.Id, item.Name, item.Phone)
		if err2 != nil { panic(err2) }
	}
}

func SQL(db *sql.DB, sql string) string  {
 var rs string = ""

	rows, err := db.Query(sql)
	if err != nil { panic(err) }
	defer rows.Close()
	
	for rows.Next() {
	 //Loop through all the Rows and display the columns  
	
	}
	return rs
}



func ReadItem(db *sql.DB) []Record {
	sql_readall := `
	SELECT Id, Name, Phone FROM items
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil { panic(err) }
	defer rows.Close()

	var result []Record
	for rows.Next() {
		item := Record{}
		err2 := rows.Scan(&item.Id, &item.Name, &item.Phone)
		if err2 != nil { panic(err2) }
		result = append(result, item)
	}
	return result
}