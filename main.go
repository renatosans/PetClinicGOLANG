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

func getPets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func main() {
	pets = append(pets, Pet {Id:1, Name: "Bethoven", Breed: "Saint Bernard", Age: 7, Owner: nil})
	pets = append(pets, Pet {Id:2, Name: "Molly", Breed: "Golden Retriever", Age: 4, Owner: nil})
	pets = append(pets, Pet {Id:3, Name: "Yoshi", Breed: "Shiba Inu", Age: 2, Owner: nil})
	pets = append(pets, Pet {Id:4, Name: "Luigi", Breed: "Beagle", Age: 9, Owner: nil})
	pets = append(pets, Pet {Id:5, Name: "Hulk", Breed: "Pit Bull", Age: 5, Owner: nil})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

    rGroup := r.Group("/api");
	rGroup.GET("/pets", getPets);
	// rGroup.POST("/pets", postPet);
	// rGroup.PATCH("/pets/{pet_id}", patchPet);
	// rGroup.DELETE("/pets/{pet_id}", deletePet);

	r.Run() // listen and serve on 0.0.0.0:8080
}
