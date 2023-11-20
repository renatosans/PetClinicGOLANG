package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Pet struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Breed   string  `json:"breed"`
	Age     int     `json:"age"`
	Owner   *Owner   `json:"owner"`
}

type Owner struct {
	name string;
}

var pets []Pet

func main() {
	pets = append(pets, Pet {Id:1, Name: "Bethoven", Breed: "Saint Bernard", Age: 7, Owner: nil})
	pets = append(pets, Pet {Id:2, Name: "Molly", Breed: "Golden Retriever", Age: 4, Owner: nil})
	pets = append(pets, Pet {Id:3, Name: "Yoshi", Breed: "Shiba Inu", Age: 2, Owner: nil})
	pets = append(pets, Pet {Id:4, Name: "Thor", Breed: "Beagle", Age: 9, Owner: nil})

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
