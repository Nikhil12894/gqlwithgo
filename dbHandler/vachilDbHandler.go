package db

import (
	"fmt"
	"math/rand"

	"github.com/Nikhil12894/gqlwithgo/graph/model"
)

// GET /books
// Get all books
func FindVachil() []*model.Vachil {
	var books []*model.Vachil
	DB.Find(&books)
	return books
}

func CreateVachil(input model.NewVachil) *model.Vachil {
	vachil := &model.Vachil{
		Type:      input.Type,
		Brand:     input.Brand,
		RegNo:     input.RegNo,
		Capacity:  input.Capacity,
		UnitPrice: input.UnitPrice,
		ID:        fmt.Sprintf("%d", rand.Int()),
	}
	// Create book/
	DB.Create(&vachil)

	return vachil
}

// GET /books/:id
// Find a book
func FindVachilByType(input string) ([]*model.Vachil, error) { // Get model if exist
	var books []*model.Vachil
	if err := DB.Where("type = ?", input).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// type UpdateBookInput struct {
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// // PATCH /books/:id
// // Update a book
// func UpdateBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Vachil
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input UpdateBookInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Model(&book).Updates(input)

// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// // DELETE /books/:id
// // Delete a book
// func DeleteBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Vachil
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	models.DB.Delete(&book)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
