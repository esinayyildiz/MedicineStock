package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// listing medicine
func listMedicines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, medicines)
}

// posting medicine
func postMedicine(c *gin.Context) {
	var newMedicine medicine
	if err := c.BindJSON(&newMedicine); err != nil {
		return
	}
	medicines = append(medicines, newMedicine)
	c.IndentedJSON(http.StatusCreated, newMedicine)
}

// medicine stock control
func searchMedicine(c *gin.Context) {
	id := c.Param("id")
	medicine, err := getMedicineById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Medicine not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, medicine)
}

// get id medicine
func getMedicineById(id string) (*medicine, error) {
	for i, b := range medicines {
		if b.Id == id {
			return &medicines[i], nil

		}
	}
	return nil, errors.New("medicine not found")
}

// check medicine, decrease stock
func checkMedicine(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id parameter"})
	}

	medicine, err := getMedicineById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "medicine not found"})
		return
	}

	if medicine.Total <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "medicine not available"})
		return
	}

	medicine.Total -= 1
	c.IndentedJSON(http.StatusOK, medicine)

}

// increase stock
func returnMedicine(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id parameter"})
	}

	medicine, err := getMedicineById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "medicine not found"})
		return
	}

	medicine.Total += 1
	c.IndentedJSON(http.StatusOK, medicine)

}
