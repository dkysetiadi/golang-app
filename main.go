package main

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )
type Payments struct{
	ID string
	Name string
	Account string
}

func main() {

	payments := []Payments{}
	payments = append(payments, Payments{"P001", "Bank BCA", "123131231"})
	payments = append(payments, Payments{"P002", "Bank BNI", "123131232"})
	payments = append(payments, Payments{"P003", "Bank BRI", "123131233"})


	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"version": "1",
		"data": payments,
	  })
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }