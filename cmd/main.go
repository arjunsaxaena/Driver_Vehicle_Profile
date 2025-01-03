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
	db, err := sqlx.Connect("postgres", "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	driverHelperStore := controllers.NewDBDriverHelperStore(db)
	vehicleStore := controllers.NewDBVehicleStore(db)
	driverHelperHandler := web.NewHandler(driverHelperStore)
	vehicleHandler := web.NewVehicleHandler(vehicleStore)

	router := gin.Default()

	// Driver Helper Routes
	router.GET("/driver_helpers/:id", driverHelperHandler.GetDriverHelperByID)
	router.GET("/driver_helpers", driverHelperHandler.GetAllDriverHelpers)
	router.POST("/driver_helpers", driverHelperHandler.CreateDriverHelper)
	router.PUT("/driver_helpers/:id", driverHelperHandler.UpdateDriverHelper)
	router.DELETE("/driver_helpers/:id", driverHelperHandler.DeleteDriverHelper)

	router.GET("/driver_helpers/driver", driverHelperHandler.GetDrivers)
	router.GET("/driver_helpers/helpers", driverHelperHandler.GetHelpers)
	router.GET("/driver_helpers/mobile/:mobile", driverHelperHandler.GetDriverHelperByMobileNumber)

	// Vehicle Routes
	router.GET("/vehicles", vehicleHandler.GetAllVehicles)
	router.POST("/vehicles", vehicleHandler.CreateVehicle)
	router.GET("/vehicles/:id", vehicleHandler.GetVehicleByID)
	router.PUT("/vehicles/:id", vehicleHandler.UpdateVehicle)
	router.DELETE("/vehicles/:id", vehicleHandler.DeleteVehicle)
	router.GET("/vehicles/driver_helper/:driver_helper_id", vehicleHandler.GetVehiclesByDriverHelperID)
	router.GET("/vehicles/route/:route_number", vehicleHandler.GetVehiclesByRouteNumber)
	router.GET("/vehicles/expired_certificates", vehicleHandler.GetExpiredCertificatesVehicles)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
