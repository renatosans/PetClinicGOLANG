package main

import (
	"strconv"
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

func postPet(c *gin.Context) {
	var newPet Pet

	// Bind JSON body to the Pet struct
	if err := c.ShouldBindJSON(&newPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Add logic to save the new pet in your database or perform any other necessary operations

	// Assuming successful creation, return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Pet created successfully", "pet": newPet})
}

func patchPet(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the Pet struct for updates
	var updatedPet Pet
	if err := c.ShouldBindJSON(&updatedPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet " + strconv.Itoa(id) + " updated successfully", "pet": updatedPet})
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
	rGroup.POST("/pets", postPet);
	rGroup.PATCH("/pets/:id", patchPet);
	// rGroup.DELETE("/pets/{pet_id}", deletePet);

	r.Run() // listen and serve on 0.0.0.0:8080
}
