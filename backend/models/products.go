package models

import (
  "fmt"
  "github.com/lib/pq"
  "log"
  "math/rand"
  "strconv"
  "time"
)

/**
	the struct of items
 */
type Items struct{
  Itemid_string string  `json:"itemid_string"`
  Itemid_count float64  `json:"itemid_count"`
}

type ItemsGlobal struct {
  Itemid string         `json:"itemid"`
  Count int64           `json:"count"`
}

type Product struct {
  id string             `json:"product_id"`
  name string           `json:"product_name"`
  cathegory int64       `json:"product_cathegory"`
  price float64         `json:"product_price"`
}


func InitProducts(ids []string) []Product {
  fmt.Println("INIT PRODUCTS")
  rand.Seed(time.Now().UnixNano())
  ps := []Product{}
  for i := 0; i < len(ids); i++ {
    p := Product{}
    p.id = ids[i]
    p.name = "product_name_" + strconv.Itoa(i)
    p.cathegory = rand.Int63n(10)+1
    p.price = float64(rand.Int63n(1000) + 20)
    ps = append(ps, p)
  }
  return ps
}

func ImportProductsToDB(ps []Product) bool {
  fmt.Println("IMPORT PRODUCTS")
  db, err := DB.Begin()
  rows, err := db.Query("SELECT COUNT (id) FROM products")
  if err != nil {
    fmt.Println("Error Count: ", err)
    log.Fatal(err)
    return false
  }
  defer rows.Close();
  var count int64
  for rows.Next() {
    rows.Scan(&count)
  }
  if (count > 0) {
    fmt.Println("db Products is not empty")
    return  false
  }
  if err != nil {
    fmt.Println("Input Error: ", err)
    log.Fatal(err)
    return false
  }
  stmt, err := db.Prepare(pq.CopyIn("products","id","name","cathegory","price"))
  for _, ev := range ps {
    _, err = stmt.Exec(ev.id,ev.name,ev.cathegory,ev.price)
    if err != nil {
      fmt.Println("Input Error: ", err)
      log.Fatal(err)
    }
  }

  _, err = stmt.Exec()
  if err != nil {
    fmt.Println("Error Import ",  err)
    log.Fatal(err)
    return false
  }

  err = stmt.Close()
  if err != nil {
    fmt.Println("Error Import ",  err)
    log.Fatal(err)
    return false
  }

  err = db.Commit()
  if err != nil {
    fmt.Println("Error Import ",  err)
    log.Fatal(err)
    return false
  }
  return true
}

func GetProductsFromDB()  []Product{
  rows, err := DB.Query("SELECT id, name, cathegory, price FROM products")
  recs := []Product{}
  if (rows == nil) {
    fmt.Println("ERROR!!")
    return []Product{}
  }
  for rows.Next() {
    rec := Product{}
    err = rows.Scan(&rec.id, &rec.name, &rec.cathegory, &rec.price)
    if err != nil {
      fmt.Println("Scan Error: ", err)
      log.Fatal(err)
    }
    recs = append(recs, rec)
  }
  if err != nil {
    fmt.Println("Reading Error: ", err)
    log.Fatal(err)
  }
  if (len(recs) > 0) {
    return recs
  } else {
    return []Product{}
  }
}

func GetProductFromDB(pr string)  Product{
  rows, err := DB.Query("SELECT id, name, cathegory, price FROM products WHERE id=$1", pr)
  if (rows == nil) {
    fmt.Println("ERROR!!")
    return Product{}
  }
  rec_ := Product{}
  for rows.Next() {
    rec := Product{}
    err = rows.Scan(&rec.id, &rec.name, &rec.cathegory, &rec.price)
    rec_ = rec
    if err != nil {
      fmt.Println("Scan Error: ", err)
      log.Fatal(err)
    }
  }
  if err != nil {
    fmt.Println("Reading Error: ", err)
    log.Fatal(err)
  }
  return rec_
}