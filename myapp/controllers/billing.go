package controllers

import (
	"fmt"
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BillsController struct to bind methods to.
type BillsController struct {
	billsrepo *repositories.BillsRepository
}

// NewBillsController creates a new BillsController with the given billsrepository.
func NewBillsController(r *repositories.BillsRepository) *BillsController {
	return &BillsController{billsrepo: r}
}

// GetbillByID handles the request to get a bill by their ID.
func (uc *BillsController) GetBillsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	bill, err := uc.billsrepo.GetBillsByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bill)
}

// GetbillByID handles the request to get a bill by their ID.
func (uc *BillsController) GetBillsList(c *gin.Context) {

	billlist, err := uc.billsrepo.ListBills()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, billlist)
}

// Createbill handles creating a new bill.
func (uc *BillsController) CreateBills(c *gin.Context) {
	fmt.Println("Entered CreateBills =============================")
	var newbill models.Billing
	if err := c.ShouldBindJSON(&newbill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("bill: %+v", newbill)
	bill, err := uc.billsrepo.CreateBills(&newbill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bill)
}

// Updatebill handles updating an existing bill.
func (uc *BillsController) UpdateBills(c *gin.Context) {
	var updatedbill models.Billing
	if err := c.ShouldBindJSON(&updatedbill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.billsrepo.UpdateBills(&updatedbill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedbill)
}

// Deletebill handles the deletion of a bill.
func (uc *BillsController) DeleteBills(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	if err := uc.billsrepo.DeleteBills(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bill deleted successfully"})
}

// RegisterbillRoutes registers the bill-related routes to the router.
func RegisterBillingRoutes(router *gin.Engine, uc *BillsController) {
	router.GET("/billing/:id", uc.GetBillsByID)
	router.POST("/billing", uc.CreateBills)
	router.PUT("/billing", uc.UpdateBills)
	router.DELETE("/billing/:id", uc.DeleteBills)
	router.GET("/billing/list", uc.GetBillsList)
}
