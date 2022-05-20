package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting service")
	runRouter()
}

func runRouter() {
	router := gin.Default()
	router.GET("/configurations", getConfigurations)

	router.Run("localhost:8080")	
}

func getConfigurations(c *gin.Context) {

	// get and validate query parameters
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 1 {
		c.IndentedJSON(http.StatusBadRequest, "id required")
		return
	}

	size := c.Query("size")
	if size == "" {
		c.IndentedJSON(http.StatusBadRequest, "size required")
		return
	}

	height, err := strconv.ParseFloat(c.Query("height"), 64)
	if err != nil || height <= 0 {
		c.IndentedJSON(http.StatusBadRequest, "height required")
		return
	}

	weight, err := strconv.ParseFloat(c.Query("weight"), 64)
	if err != nil || weight <= 0 {
		c.IndentedJSON(http.StatusBadRequest, "weight required")
		return
	}

	calculate(id, size, height, weight)

	c.IndentedJSON(http.StatusOK, id)
}

func calculate(id int, size string, height float64, weight float64) {
	fmt.Printf("id: %v \n", id)
	fmt.Printf("size: %v \n", size)
	fmt.Printf("height: %v \n", height)
	fmt.Printf("weight: %v \n", weight)

	// TODO: perform calculation here. Should return a struct of the result
}
