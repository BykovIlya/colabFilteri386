package models

import (
	"fmt"
	"log"
	"github.com/lib/pq"
)

func IsEmptyTargetNN() bool {
	db, err := DB.Begin()
	rows, err := db.Query("SELECT COUNT (yes) FROM targetNN")
	if err != nil {
		fmt.Println("Error Count: ", err)
		log.Fatal(err)
		return false
	}
	defer rows.Close();
	var count int64
	for rows.Next() {
		rows.Scan(&count)
		if (count > 0) {
			fmt.Println("db TARGETNN is not empty")
			return  false
		}
	}
	return true
}
func ImportTargetNNToDB(arr [][]float64) bool {
	db, err := DB.Begin()
	/*rows, err := db.Query("SELECT COUNT (yes) FROM targetNN")
	if err != nil {
		fmt.Println("Error Count: ", err)
		log.Fatal(err)
		return false
	}
	defer rows.Close();
	var count int64
	for rows.Next() {
		rows.Scan(&count)
		if (count > 0) {
			fmt.Println("db TARGETNN is not empty")
			return  false
		}
	}*/

	if err != nil {
		fmt.Println("Input Error 1: ", err)
		log.Fatal(err)
		return false
	}
	stmt, err := db.Prepare(pq.CopyIn("targetnn","yes", "nnn"))
	for i := 0; i  < len(arr); i++ {
			_, err = stmt.Exec(arr[i][0], arr[i][1])
			if err != nil {
				fmt.Println("Input Error 2: ", err)
				log.Fatal(err)
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

func GetTargetNNFromDB() [][]float64 {
	rows, err := DB.Query("SELECT yes,nnn FROM targetnn")
	recs := make([][]float64, 0)
	if (rows == nil) {
		fmt.Println("ERROR!!")
		return make([][]float64, 0)
	}
	for rows.Next() {
		rec := make([]float64,2)
		err = rows.Scan(&rec[0], &rec[1])
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
		return make([][]float64, 0)
	}
}