package models

import (
	"math"

	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"sort"
  "fmt"
)

/**
	reading data from .csv
 */
func ReadingTransactionsFromFile(csvFileName string) []Events {
	csvFile, err := os.Open(csvFileName)
  if err != nil {
    fmt.Println("Error with reading from file:", err)
    log.Fatal(err)
  }

  fmt.Println("success open file!")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var events []Events
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[2] == "transaction" {
      var event= Events{}
      event.Timestamp = line[0]
      event.Visitorid = line[1]
      event.Event_ = line[2]
      event.Itemid = line[3]
      event.Transactionid = line[4]

      events = append(events, event)
    	}
	}

	return events
}

func MakeUniqArrayOfVisitors(events []Events) []string {
	bufOfVisitors := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfVisitors[i] = events[i].Visitorid
	}
	sort.Strings(bufOfVisitors)
	removeDublicatesOfVisitors := removeDuplicates(bufOfVisitors)
	return removeDublicatesOfVisitors
}

func MakeUniqArrayOfItems(events []Events) [] string {
	bufOfItems := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfItems[i] = events[i].Itemid
	}
	sort.Strings(bufOfItems)
	removeDublicatesOfItems := removeDuplicates(bufOfItems)
	return removeDublicatesOfItems
}

func MakeMatrixOfSales (visitors [] Visitor, removeDublicatesOfVisitors [] string, removeDublicatesOfItems [] string) [][] float64{
	/*
		init matrix
	 */
	matrixOfSales := make([][] float64, len(removeDublicatesOfVisitors))
	for i := 0; i < len(removeDublicatesOfVisitors); i++  {
		matrixOfSales[i] = make([] float64, len(removeDublicatesOfItems))
	}
	/*
	make matrix
	 */
	for i := 0; i < len(removeDublicatesOfVisitors); i++ {
		for j := 0; j < len(visitors[i].Items); j++ {
			matrixOfSales[i][getIndItem(removeDublicatesOfItems,visitors[i].Items[j].Itemid_string)] = visitors[i].Items[j].Itemid_count;
		}
	}
	return matrixOfSales
}

func MakeArrayOfSales (matrixOfSales [][] float64, n int, m int) [] float64 {
	arrayOfSales := make ([]float64, 0)
	arrayOfSales = toArray(matrixOfSales, n, m, arrayOfSales)
	return arrayOfSales
}
func AddCountToEachProductOfEachVisitor (visitors [] Visitor) {
	for i := 0; i < len(visitors); i++  {
		sort.Slice(visitors[i].Items, func(j, k int) bool { return visitors[i].Items[j].Itemid_string < visitors[i].Items[k].Itemid_string })
	}
	for i := 0; i < len(visitors); i++ {
		visitors[i].Items = findCount(visitors[i].Items)
	}
}
/**
	get index of visitor
 */
func GetIndVisitor (visitor [] Visitor, finder string) int {
	for i := 0; i < len(visitor); i++ {
		if visitor[i].Visitorid_string == finder {
			return i
		}
	}
	return -1
}

/**
	get index of item
 */
func getIndItem (items [] string, finder string) int {
	for i := 0; i < len(items); i++ {
		if items[i] == finder {
			return i
		}
	}
	return -1
}

/**
	set the field visitorid_strnig of the structure Visitor to the value of unique visitors from the array buffer
 */
func InitVisitors (visitor [] Visitor, buffer [] string) {
	for i := 0; i < len(buffer); i++ {
		visitor[i].Visitorid_string =  buffer[i]
	}
}

/**
	set each visitor an array of items
 */
func AddItemsToVisitor (visitor [] Visitor, events []Events){
	for i := 0; i < len(visitor); i++ {
		for j := 0; j < len(events); j++ {
			if visitor[i].Visitorid_string == events[j].Visitorid {
				visitor[i].Items = append(visitor[i].Items, Items{
					Itemid_string: events[j].Itemid,
					Itemid_count: 1,
				})
			}
		}
	}
}

/**
	remove dublicates from visitors and itmes for make uniq arrays
 */
func removeDuplicates(array [] string) [] string{
	if len(array) == 1 || len(array) == 0 {
		return array
	}
	unique := 1
	for i := 1; i < len(array); i++{
		if array[i] != array[i - 1] {
			unique++;
		}
	}
	result := make([] string, unique)
	k := 0;
	if len(result) > 0 {
		result[k] = array[0]
		k++
	}
	for i := 1; i < len(array); i++ {
		if array[i] != array[i - 1] {
			result[k] = array[i];
			k++
		}
	}
	return result;
}

/**
	convert matrix to array
 */

func toArray (matrix [][] float64, n int, m int, array [] float64) []float64 {
	fmt.Println("start writing to array")
	fmt.Println("length:", n)
	LowerLength := 2000
	for i := 0; i < /*n -replace for my comp*/ LowerLength; i++  {
		for j := 0; j < m; j++ {
			array = append(array, matrix[i][j])
		}
		//fmt.Println(i)
	}
	fmt.Println("end writing to array")
	return array
}

/**
	find count of each items in array of items for each visitor
 */
func findCount (item []Items) [] Items{
	buffer := make( [] Items, 0);
	var prev string
	for i := 0; i < len(item); i++ {
		if (item[i].Itemid_string != prev) {
			buffer = append(buffer, Items {
				item[i].Itemid_string,
				1,
			})
		} else {
			buffer[len(buffer) - 1].Itemid_count++
		}
		prev = item[i].Itemid_string
	}
	return buffer
}

/*
	remove unnecessary elements from score array
 */
func optimizeScores(scores [] float64, good [] float64) []float64{
	for i := 0; i < len(scores); i++ {
		if !math.IsNaN(scores[i]) {
			good = append(good, scores[i])
		}
	}
	return good
}

func FindProductInPerson (id int, id_product string, visitors []Visitor) bool {
	for j := 0; j < len(visitors[id].Items) ; j++ {
		if (id_product == visitors[id].Items[j].Itemid_string) {
			return true
		}
	}
	return false
}