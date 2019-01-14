package models

import (
  "fmt"
  "log"
  "github.com/lib/pq"
)

/**
	the struct of events
 */
type Events struct {
  Timestamp string /*int64*/ `form:"timestamp" json:"timestamp"`
  Visitorid string /*int64*/  `form:"visitorid" json:"visitorid"`
  Event_ string /*object*/ `form:"event" json:"event"`
  Itemid string /*int64*/  `form:"itemid" json:"itemid"`
  Transactionid string /*float64*/ `form:"transactionid" json:"transactionid"`
}

func InsertOneEventToDB(event Events) {
  fmt.Println(event)

  row:=DB.QueryRow("INSERT INTO events(timestamp,visitorid,event,itemid,transactionid) VALUES($1,$2,$3,$4,$5)",
    event.Timestamp, event.Visitorid, event.Event_, event.Itemid, event.Transactionid)
  fmt.Println(row)
}

func ImportEventsToDB(events []Events) bool {
  db, err := DB.Begin()
  if err != nil {
    fmt.Println("Input Error: ", err)
    log.Fatal(err)
    return false
  }
  i := 0
  fmt.Println("start of adding data to events db")
  stmt, err := db.Prepare(pq.CopyIn("events","timestamp","visitorid","event","itemid","transactionid"))
  for _, ev := range events {
    _, err = stmt.Exec(ev.Timestamp, ev.Visitorid, ev.Event_, ev.Itemid, ev.Transactionid)
    i++
    if err != nil {
      fmt.Println("Input Error: ", err)
      log.Fatal(err)
    }
  }
  fmt.Println("count of events: ", i ,"\n");
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

func ReadEventsFromDB() []Events{
  rows, err := DB.Query("SELECT timestamp,visitorid,event,itemid,transactionid FROM events")
  if err != nil {
    fmt.Println("Reading Error: ", err)
    log.Fatal(err)
  }
  evs := []Events{}
  for rows.Next() {
    ev := Events{}
    err = rows.Scan(&ev.Timestamp, &ev.Visitorid, &ev.Event_, &ev.Itemid, &ev.Transactionid)
    if err != nil {
      fmt.Println("Scan Error: ", err)
      log.Fatal(err)
    }
    evs = append(evs, ev)
  }
  return evs;
}