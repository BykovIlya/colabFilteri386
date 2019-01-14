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
	the struct of visitors
 */
type Visitor struct {
  Visitorid_string string `json:"visitorid_string"`
  Items [] Items          `json:"items_array"`
}

type Person struct {
  id string                `json:"person_id"`
  name string             `json:"person_name"`
  surname string          `json:"person_surname"`
  age int64               `json:"person_age"`
  gender bool             `json:"person_gender"`
  properties string       `json:"person_properties"`
}

func ImportVisitorsToDB(visitor []Visitor) bool {
  db, err := DB.Begin()
  if err != nil {
    fmt.Println("Input Error visitors 1: ", err)
    log.Fatal(err)
    return false
  }
  stmt, err := db.Prepare(pq.CopyIn("visitors","visitor_id", "item_id"))
  for i := 0; i  < len(visitor); i++ {
    for j := 0; j < len(visitor[i].Items); j++ {
      _, err = stmt.Exec(visitor[i].Visitorid_string, visitor[i].Items[j].Itemid_string)
      if err != nil {
        fmt.Println("Input Error visitors 2: ", err)
        log.Fatal(err)
        }
      }
    }
  _, err = stmt.Exec()
  if err != nil {
    fmt.Println("Error Import ", err)
    log.Fatal(err)
    return false
  }

  err = stmt.Close()
  if err != nil {
    fmt.Println("Error Import ", err)
    log.Fatal(err)
    return false
  }

  err = db.Commit()
  if err != nil {
    fmt.Println("Error Import ", err)
    log.Fatal(err)
    return false
  }

  return true
}

func InitPersons(ids []string) []Person {
  fmt.Println("INIT PERSONS")
  rand.Seed(time.Now().UnixNano())
  ps := []Person{}
  for i := 0; i < len(ids); i++ {
    p := Person{}
    p.id = ids[i]
    p.name = "user_name_" + strconv.Itoa(i)
    p.surname = "user_surname_" + strconv.Itoa(i)
    p.age = rand.Int63n(56) + 14
    gender := rand.Intn(2)
    if (gender == 1) {
      p.gender = false
    } else {
      p.gender = true;
    }
    p.properties = "nothing"
    ps = append(ps, p)
  }
  return ps
}

func ImportPersonsToDB(ps []Person) bool {
  fmt.Println("IMPORT PERSONS")
  db, err := DB.Begin()
  rows, err := db.Query("SELECT COUNT (id) FROM persons")
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
    fmt.Println("db Persons is not empty")
    return  false
  }
  if err != nil {
    fmt.Println("Input Error: ", err)
    log.Fatal(err)
    return false
  }
  stmt, err := db.Prepare(pq.CopyIn("persons","id","name","surname","age","gender","properties"))
  for _, ev := range ps {
    _, err = stmt.Exec(ev.id,ev.name,ev.surname,ev.age,ev.gender,ev.properties)
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

func GetPersonsFromDB()  []Person{
  rows, err := DB.Query("SELECT id, name, surname, age, gender, properties FROM persons")
  recs := []Person{}
  if (rows == nil) {
    fmt.Println("ERROR!!")
    return []Person{}
  }
  for rows.Next() {
    rec := Person{}
    err = rows.Scan(&rec.id, &rec.name, &rec.surname, &rec.age, &rec.gender, &rec.properties)
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
    return []Person{}
  }
}