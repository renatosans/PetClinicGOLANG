package main

import (
	"net/http"
	"petclinic/prisma/db"
	"petclinic/src/handlers"
	"petclinic/src/utils"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getPets(c *gin.Context) {
	// var pets []db.InnerPet

	client := utils.GetPrisma(c)

	pets, err := client.Pet.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func postPet(c *gin.Context) {
	var payload db.InnerPet

	// Bind JSON body to the Pet struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	insertedPet, err := client.Pet.CreateOne(
		db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.FlagRemoved.Set(payload.FlagRemoved),
		db.Pet.Age.SetOptional(payload.Age),
		db.Pet.Owner.SetOptional(payload.Owner),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet created successfully", "pet": insertedPet})
}

func patchPet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the Pet struct
	var payload db.InnerPet
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	updatedPet, err := client.Pet.FindUnique(
		db.Pet.ID.Equals(id),
	).Update(
		db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.FlagRemoved.Set(payload.FlagRemoved),
		db.Pet.Age.SetOptional(payload.Age),
		db.Pet.Owner.SetOptional(payload.Owner),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet patched", "pet": updatedPet})
}

func deletePet(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := utils.GetPrisma(c)
	deletedPet, err := client.Pet.FindUnique(
		db.Pet.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully", "pet id": deletedPet.ID})
}

func main() {
	godotenv.Load(".env")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	rGroup.GET("/pets", getPets)
	rGroup.POST("/pets", postPet)
	rGroup.PATCH("/pets/:id", patchPet)
	rGroup.DELETE("/pets/:id", deletePet)
	rGroup.POST("/veterinarians", handlers.PostVeterinarian)

	router.Use(cors.Default())
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
