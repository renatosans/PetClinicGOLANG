package main

import (
	"context"
	"fmt"
	"net/http"
	"petClinicAPI/prisma/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var client *db.PrismaClient
var pets []db.InnerPet

type Pet struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Owner *Owner `json:"owner"`
}

type Owner struct {
	name  string
	email string
}

func getPets(c *gin.Context) {
	// pets, err := client.Pet.FindMany().Exec(ctx)
	// c.JSON(http.StatusOK, gin.H{"data": pets})
}

func postPet(c *gin.Context) {
	var newPet Pet

	// Bind JSON body to the Pet struct
	if err := c.ShouldBindJSON(&newPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Add logic to save the new pet in your database or perform any other necessary operations
	// insertedPet, err := client.Pet.CreateOne().Exec(ctx)

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

func deletePet(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet " + strconv.Itoa(id) + " deleted successfully"})
}

func main() {
	godotenv.Load(".env")

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pets, err := client.Pet.FindMany().Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query result: %v\n", pets)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := r.Group("/api")
	rGroup.GET("/pets", getPets)
	rGroup.POST("/pets", postPet)
	rGroup.PATCH("/pets/:id", patchPet)
	rGroup.DELETE("/pets/:id", deletePet)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
