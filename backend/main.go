package main

import (
	"path/filepath"
	"os"
  "fmt"
  "github.com/gin-gonic/gin"
  "smartstore-analytics/backend/routes"
  "log"
  "smartstore-analytics/backend/models"
)

func main() {
	if models.InitDB() {

		models.CreateDB(models.DB)

		defer models.DB.Close()

	} else {
		log.Panic("Error:db not init")
	}

	CreateDirsForFiles()
	csvFileName := "api/upload/" + "File.csv"
	if _,err := os.Stat(csvFileName); err != nil {
	if os.IsNotExist(err) {
	  	fmt.Println("file does not exist\n")
	}
	} else {
		routes.Algorithm(csvFileName)
	}

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Static("/api/c/tmp", "./api/c/tmp")
	router.POST("/import", routes.ImportEvents)
	router.GET("/events",routes.GetEvents)
	router.GET("/users", routes.GetUsers)
	router.GET("/products", routes.GetProducts)
	router.GET("/recommendations", routes.GetRecommends)
	router.GET("/recommendations/:id", routes.GetPerson)
	router.GET("/users/:id", routes.GetPerson)
	//router.GET("/neuralnetwork/:age/:gender/:category/:price", routes.GetResult)
	router.Run(":5001")
}

func CreateDirsForFiles() {
	tmp := filepath.Join(".", "api/c/tmp")
	os.MkdirAll(tmp, os.ModePerm)

	upload := filepath.Join(".", "api/upload")
	os.MkdirAll(upload, os.ModePerm)
}

func CORSMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Max-Age", "86400")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length,X-Requested-With, Accept-Encoding, X-CSRF-Token, Authorization")
    c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

    if c.Request.Method == "OPTIONS" {
      fmt.Println("OPTIONS")
      c.AbortWithStatus(200)
    } else {
      c.Next()
    }
  }
}
