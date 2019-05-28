package main

import (
	"github.com/gin-gonic/gin"
	 
)
 

func main() {
    r := gin.Default()
    r.GET("/ceshi", func(c *gin.Context) {
    c.JSON(200, gin.H{
    "message": "Hello World",
    "路径":"/src/cce",
    })
    })
    r.Run(":9999") // listen and server on 0.0.0.0:8080
   }
 
 
 