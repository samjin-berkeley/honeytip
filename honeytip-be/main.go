package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var todos = []todo{
	{ID: 1, Name: "Sam"},
	{ID: 2, Name: "DY"},
	{ID: 3, Name: "Nam"},
}

// getTodos responds with the list of all Todos with JSON
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)

	router.Run("localhost:8080")
}
