package routes

import (
  "fmt"
  "github.com/gin-gonic/gin"
  . "github.com/skelterjohn/go.matrix"
  "net/http"
  "smartstore-analytics/backend/algorithm"
  "smartstore-analytics/backend/models"
  "strconv"
)

var events []models.Events
var visitors []models.Visitor
var removeDublicatesOfVisitors []string
var removeDublicatesOfItems []string
var items []models.ItemsGlobal
var recommendations []algorithm.Recommendation
var matrixOfSales [][]float64
var arrayOfSales []float64
var prefs *DenseMatrix
var products []string
var isImported = false

type EventsList struct {
  Events []models.Events `json:"events"`
  Total int `json:"Total"`
}

func ImportEvents(c *gin.Context)  {
  file, err := c.FormFile("file")
  if err != nil {
    c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
    return
  }
  csvFileName:="api/upload/"+"File.csv"
  if err := c.SaveUploadedFile(file, csvFileName); err != nil {
    c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
    return
  }
  models.ClearDB(models.DB,"events")
  models.ClearDB(models.DB, "visitors")
  models.ClearDB(models.DB,"recommends")
  events = models.ReadingTransactionsFromFile(csvFileName)
  models.ImportEventsToDB(events)
  removeDublicatesOfVisitors = models.MakeUniqArrayOfVisitors(events)
  models.ImportPersonsToDB(models.InitPersons(removeDublicatesOfVisitors))
  removeDublicatesOfItems = models.MakeUniqArrayOfItems(events)
  models.ImportProductsToDB(models.InitProducts(removeDublicatesOfItems))
  isImported = true
  Algorithm(csvFileName)
}

func Algorithm(csvFileName string)  {
  if (isImported == false) {
    events = models.ReadEventsFromDB()
    removeDublicatesOfVisitors = models.MakeUniqArrayOfVisitors(events)
    models.ImportPersonsToDB(models.InitPersons(removeDublicatesOfVisitors))

    removeDublicatesOfItems = models.MakeUniqArrayOfItems(events)
    models.ImportProductsToDB(models.InitProducts(removeDublicatesOfItems))

    fmt.Println("Read from DB!")
  }
  visitors = make([] models.Visitor, len(removeDublicatesOfVisitors))
  fmt.Println("COUNT OF VISITORS:", len(visitors))
  /* make struct of visitors */
  models.InitVisitors(visitors, removeDublicatesOfVisitors)
  /* add items to each visitor */
  models.AddItemsToVisitor(visitors,events)
  models.AddCountToEachProductOfEachVisitor(visitors)
  items = make ([]models.ItemsGlobal, len(events))
  for i := 0; i < len(events); i++ {
    items[i].Itemid = events[i].Itemid
    items[i].Count = 1
  }
  matrixOfSales = models.MakeMatrixOfSales(visitors, removeDublicatesOfVisitors, removeDublicatesOfItems)

  /* init array of sales to get it into CA */
  arrayOfSales = models.MakeArrayOfSales(matrixOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))

  /* CA algorithm*/
  prefs = algorithm.MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
  products = make([]string, 0)
  for i := 0; i < len(removeDublicatesOfItems); i++ {
    products = append(products, strconv.Itoa(i))
  }
/*  need := false
  models.CreateNeuralNetworkPerson(models.GetPersonsFromDB(), models.GetProductsFromDB(), visitors, need)
  nn := gonn.LoadNN("gonnPerson")
  out := nn.Forward([]float64{0.0, 55.0, 12.0, 19.0})
  fmt.Println("nn result:", models.GetResult(out))   */   //  <---- neuralNetwork
  models.ImportVisitorsToDB(visitors)
}

func GetEvents (c *gin.Context) {
  c.JSON(200,events)
}

