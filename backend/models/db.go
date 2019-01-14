package models

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "log"
)

var DB *sql.DB

func InitDB() bool{
  connection := "host=localhost port=5432 user=postgres password=postgres dbname=ColabFilter sslmode=disable"

 // connection := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

  db, err := sql.Open("postgres", connection)
  if err != nil {
    log.Fatal(err)
    return false
  }
  DB = db
  return true
}

func CreateDB(db *sql.DB) {
  fmt.Println("Start creating db")
  sql := `
  CREATE TABLE IF NOT EXISTS persons(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    surname VARCHAR(255),
    age integer,
    gender boolean,
    properties VARCHAR(255)
  );
  CREATE TABLE IF NOT EXISTS products(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    cathegory integer,
    price real
  );
  CREATE TABLE IF NOT EXISTS events(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    timestamp VARCHAR(255) NOT NULL,
	visitorid VARCHAR(255) NOT NULL, 
	event VARCHAR(255) NOT NULL,
	itemid VARCHAR(255) NOT NULL,
	transactionid VARCHAR(255)
  );
  CREATE TABLE IF NOT EXISTS recommends(
    user_id VARCHAR(255),
    recommend float,
    score float
  );
  CREATE TABLE IF NOT EXISTS visitors(
    visitor_id VARCHAR(255),
    item_id VARCHAR(255)
  );
  CREATE TABLE IF NOT EXISTS inputnn(
    gender float,
    age float,
    category float,
    price float
  );
  CREATE TABLE IF NOT EXISTS targetnn(
    yes float,
    nnn float
  );
  `
  _, err := db.Exec(sql)
  if err != nil {
    panic(err)
  }
  fmt.Println("End creating bd")
}

func ClearDB( db *sql.DB, name string) bool {
  stmt, err := db.Prepare("delete from " + name)
  if (err != nil) {
    panic(err)
  }
  defer stmt.Close()
  res, err := stmt.Exec()
  if (err != nil) {
    panic(err)
  }
  affect, err := res.RowsAffected()
  if (err != nil) {
    panic(err)
  }
  fmt.Println(affect," rows deleted")
  return true
}

