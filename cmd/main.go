package main

import (
	"log"
	"net/http"

	"github.com/arjunsaxaena/driver_vehicle_profile/controllers"
	"github.com/arjunsaxaena/driver_vehicle_profile/web"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "postgres://postgres:secret@127.0.0.3:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	driverHelperStore := controllers.NewDBDriverHelperStore(db)

	handler := web.NewHandler(driverHelperStore)

	router := gin.Default()

	router.GET("/driver_helpers/:id", handler.GetDriverHelperByID)
	router.GET("/driver_helpers", handler.GetAllDriverHelpers)
	router.POST("/driver_helpers", handler.CreateDriverHelper)
	router.PUT("/driver_helpers/:id", handler.UpdateDriverHelper)
	router.DELETE("/driver_helpers/:id", handler.DeleteDriverHelper)

	router.GET("/driver_helpers/driver", handler.GetDrivers)
	router.GET("/driver_helpers/helpers", handler.GetHelpers)

	router.GET("/driver_helpers/mobile/:mobile", handler.GetDriverHelperByMobileNumber)
	router.GET("/driver_helpers/license/:license", handler.GetDriverHelperByLicenseNumber)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
