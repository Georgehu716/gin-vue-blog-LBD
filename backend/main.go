package main

import (
	"gin-vue-blog-LBD/models"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProducts(c *gin.Context) {
	products := models.GetProducts()

	if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

func getProduct(c *gin.Context) {
	code := c.Param("code")

	product := models.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func addProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:code", getProduct)
	router.POST("/products", addProduct)
	router.Run("localhost:9999")
}