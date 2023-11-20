package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Pet struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Species Species `json:"species"`
	Diet    string  `json:"diet"`
}

{"id":1,"name":"Bethoven","breed":"Saint Bernard","age":7,"owner":null},
{"id":2,"name":"Molly","breed":"Golden Retriever","age":4,"owner":null},
{"id":3,"name":"Yoshi","breed":"Shiba Inu","age":2,"owner":null},
{"id":4,"name":"Thor","breed":"Beagle","age":9,"owner":null}


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hell yeah",
		})
	})

    rGroup := r.Group("/api");
	rGroup.GET(getAnimals);
	rGroup.POST(postAnimals);
	rGroup.DELETE(deleteAnimal);

	r.Run() // listen and serve on 0.0.0.0:8080
}
