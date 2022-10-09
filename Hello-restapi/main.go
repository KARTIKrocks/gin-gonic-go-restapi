package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

func main() {
	r := gin.Default() // this is our instance for gin

	// HTTP, GET, POST, PUT, DELETE
	r.GET("/", getHello)
	r.POST("/", postHello)
	r.PUT("/", putHello)
	r.DELETE("/", deleteHello)

	// group in
	r1 := r.Group("/api")
	{
		r1.GET("/hello", getHello) // http://<server>/api/hello
		r1.POST("/hello", postHello)
		r1.PUT("/hello", putHello)
		r1.DELETE("/hello", deleteHello)
	}

	// handling path params
	r.GET("/product/:id", getProductById)
	r.GET("/profile/:username", showProfile)
	r.GET("/compute/:num1/add/:num2", compute)

	// handling query string params
	// /emplyee?firstname=zzzz&lastname=cccc&id=2
	r.GET("/employee", showEmployee)

	// binding post data
	r.POST("/product", performProduct)
	r.POST("/products", performProducts)

	// reading .enc file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
		//panic(err)
	}
	vala := os.Getenv("VALA")
	valb := os.Getenv("VALB")
	val_a, _ := strconv.ParseInt(vala, 10, 0)
	val_b, _ := strconv.ParseInt(valb, 10, 0)
	result := val_a + val_b
	fmt.Printf("\n====================\n  %d + %d = %d \n", val_a, val_b, result)

	r.Run() // run 8080
	fmt.Println("server is running")
}
func performProduct(c *gin.Context) {
	var product Product
	err := c.BindJSON(&product)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}
func performProducts(c *gin.Context) {
	var products []Product
	err := c.BindJSON(&products)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}
func showEmployee(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "")
	lastname := c.DefaultQuery("lastname", "")
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 0)
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":        id,
		"firstname": firstname,
		"lastname":  lastname,
	})
}
func getProductById(c *gin.Context) {
	id := c.Param("id") // string
	idn, _ := strconv.ParseInt(id, 10, 0)

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":   idn,
		"name": "Product A",
	})
}
func showProfile(c *gin.Context) {
	username := c.Param("username") // string

	c.IndentedJSON(http.StatusOK, gin.H{
		"username": username,
	})
}
func compute(c *gin.Context) {
	num1, _ := strconv.ParseInt(c.Param("num1"), 10, 0)
	num2, _ := strconv.ParseInt(c.Param("num2"), 10, 0)

	result := num1 + num2

	c.IndentedJSON(http.StatusOK, gin.H{
		"num1":   num1,
		"num2":   num2,
		"result": result,
	})
}

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "hello  REST API - HTTP GET",
	})
}
func postHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "hello  REST API - HTTP POST",
	})
}
func putHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "hello  REST API - HTTP PUT",
	})
}
func deleteHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "hello  REST API - HTTP DELETE",
	})
}
