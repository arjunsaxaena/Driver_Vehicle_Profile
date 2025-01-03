package web

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/arjunsaxaena/driver_vehicle_profile/controllers"
	"github.com/arjunsaxaena/driver_vehicle_profile/model"
)

type VehicleHandler struct {
	Store *controllers.DBVehicleStore
}

func NewVehicleHandler(store *controllers.DBVehicleStore) *VehicleHandler {
	return &VehicleHandler{Store: store}
}

func (h *VehicleHandler) CreateVehicle(c *gin.Context) {
	var v model.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	v.ID = uuid.New()
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()

	if err := h.Store.CreateVehicle(&v); err != nil {
		if strings.Contains(err.Error(), "unique_violation") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Vehicle with number already exists",
				"details": "A vehicle with the same vehicle number already exists.",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vehicle", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Vehicle created successfully", "vehicle": v})
}

func (h *VehicleHandler) UpdateVehicle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var v model.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	v.ID = id
	v.UpdatedAt = time.Now()

	if err := h.Store.UpdateVehicle(&v); err != nil {
		if strings.Contains(err.Error(), "constraint violation") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed to update vehicle due to constraint violation",
				"details": "Please check the input values (e.g., seats_available should not exceed total_students_capacity).",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vehicle", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vehicle updated successfully", "vehicle": v})
}

func (h *VehicleHandler) DeleteVehicle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.Store.DeleteVehicle(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vehicle", "details": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Vehicle deleted successfully"})
}

func (h *VehicleHandler) GetAllVehicles(c *gin.Context) {
	vehicles, err := h.Store.Vehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vehicles", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}

func (h *VehicleHandler) GetVehicleByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	v, err := h.Store.VehicleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicle": v})
}

func (h *VehicleHandler) GetVehiclesByDriverHelperID(c *gin.Context) {
	driverHelperIDParam := c.Param("driver_helper_id")
	driverHelperID, err := uuid.Parse(driverHelperIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Driver Helper ID format"})
		return
	}

	vehicles, err := h.Store.VehiclesByDriverHelperID(driverHelperID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vehicles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}

func (h *VehicleHandler) GetVehiclesByRouteNumber(c *gin.Context) {
	routeNumber := c.Param("route _number")
	if routeNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Route number is required"})
		return
	}

	vehicles, err := h.Store.VehiclesByRouteNumber(routeNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vehicles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}

func (h *VehicleHandler) GetExpiredCertificatesVehicles(c *gin.Context) {
	vehicles, err := h.Store.ExpiredCertificatesVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vehicles with expired certificates", "details": err.Error()})
		return
	}

	if len(vehicles) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No vehicles with expired certificates found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}
