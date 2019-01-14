package routes

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "smartstore-analytics/backend/algorithm"
  "smartstore-analytics/backend/utils"
  "smartstore-analytics/backend/models"
)

var myVisitor string

func GetUsers (c *gin.Context) {
  c.JSON(200,visitors)
}

func GetPerson(c *gin.Context)  {
  myVisitor = c.Param("id")
  recommendations = models.GetRecommendsFromBD(myVisitor)
  if (len(recommendations) == 0) {
    indexOfVisitor := models.GetIndVisitor(visitors, myVisitor)
    if (indexOfVisitor == -1) {
      fmt.Println("Error: visitor doesn't found!")
      c.JSON(400, utils.ApiMessage{"User doesn't found"})
    }
    var err error
    recommendations, err = algorithm.GetRecommendations(prefs, /*getIndVisitor(visitors, myVisitor)*/ indexOfVisitor, products)
    if err != nil {
      fmt.Println("WHAT!?")
    }
    if (len(recommendations) > 0) {
      models.ImportRecommendsToDB(myVisitor, recommendations)
      c.JSON(200, recommendations)
    } else {
      	fmt.Println("No recommendations!")
    }
  } else {
    c.JSON(200, recommendations)
  }
}
