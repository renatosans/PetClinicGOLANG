package utils;

import (
	"net/http"
	"petclinic/prisma/db"
	"github.com/gin-gonic/gin"
)

// TODO:  verficar se o prisma trabalha com Pool de Conexões
func GetPrisma(c *gin.Context) *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	return client
}

/*
func GetPool(databaseURL string) *pgxpool.Pool {
    config, err := pgxpool.ParseConfig(databaseURL)
    if err != nil {
		panic(err)
    }
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
		panic(err)
    }

	return pool;
}

func AppMiddleware(pool *pgxpool.Pool) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("pool", pool)
        c.Next()
    }
}
*/
