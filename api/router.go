package api

import (
	"database/sql"

	"github.com/OhVIton/ninsho-kun-backend/controllers"
	"github.com/OhVIton/ninsho-kun-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	ser := services.NewAppService(db)
	aCon := controllers.NewAuthCodeController(ser)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/authcode", aCon.FetchAuthCode)

	return r
}
