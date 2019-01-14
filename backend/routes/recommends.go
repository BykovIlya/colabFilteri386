package routes

import "github.com/gin-gonic/gin"

func GetRecommends (c *gin.Context) {
  if myVisitor != "" {
    c.JSON(200, recommendations)
  }
}

