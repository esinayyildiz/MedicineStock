package main

import (
	"github.com/gin-gonic/gin"
)

var medicines = []medicine{
	{
		Id:             "1",
		Name:           "Medicine 1",
		ExpirationDate: "11/08/2025",
		Total:          22},
	{
		Id:             "2",
		Name:           "Medicine 2",
		ExpirationDate: "11/08/2025",
		Total:          50},
	{
		Id:             "3",
		Name:           "Medicine 3",
		ExpirationDate: "11/08/2025",
		Total:          10},
	{
		Id:             "4",
		Name:           "Medicine 4",
		ExpirationDate: "11/08/2025",
		Total:          7},
}

// main fuction
func main() {
	router := gin.Default()
	router.GET("/books", listMedicines)
	router.POST("/books", postMedicine)
	router.GET("/books/:id", searchMedicine)
	router.PATCH("/checkout", checkMedicine)
	router.PATCH("/return", returnMedicine)
	router.Run("localhost:8080")
}
