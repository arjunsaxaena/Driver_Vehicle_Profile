package web

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/arjunsaxaena/driver_vehicle_profile/controllers"
	"github.com/arjunsaxaena/driver_vehicle_profile/model"
)

type Handler struct {
	Store *controllers.DBDriverHelperStore
}

func NewHandler(store *controllers.DBDriverHelperStore) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) GetDriverHelperByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	dh, err := h.Store.DriverHelperByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver/Helper not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"driver_helper": dh})
}

func (h *Handler) GetAllDriverHelpers(c *gin.Context) {
	dhs, err := h.Store.DriverHelpers()
	if err != nil {
		log.Printf("Error retrieving driver/helpers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve driver or helpers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"driver_helpers": dhs})
}

func (h *Handler) GetDrivers(c *gin.Context) {
	drivers, err := h.Store.Drivers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve drivers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"drivers": drivers})
}

func (h *Handler) GetHelpers(c *gin.Context) {
	helpers, err := h.Store.Helpers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve helpers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"helpers": helpers})
}

func (h *Handler) CreateDriverHelper(c *gin.Context) {
	var dh model.DriverHelper
	if err := c.ShouldBindJSON(&dh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	dh.ID = uuid.New()
	dh.CreatedAt = time.Now()
	dh.UpdatedAt = time.Now()

	if err := h.Store.CreateDriverHelper(&dh); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create driver/helper", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Driver/Helper created successfully", "driver_helper": dh})
}

func (h *Handler) UpdateDriverHelper(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dh model.DriverHelper
	if err := c.ShouldBindJSON(&dh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	dh.ID = id
	dh.UpdatedAt = time.Now()

	if err := h.Store.UpdateDriverHelper(&dh); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update driver/helper", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Driver/Helper updated successfully", "driver_helper": dh})
}

func (h *Handler) DeleteDriverHelper(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.Store.DeleteDriverHelper(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete driver/helper", "details": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Driver/Helper deleted successfully"})
}

func (h *Handler) GetDriverHelperByMobileNumber(c *gin.Context) {
	mobile := c.Param("mobile")
	if mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number is required"})
		return
	}

	dh, err := h.Store.DriverHelperByMobileNumber(mobile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver/Helper not found with the given mobile number"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"driver_helper": dh})
}
