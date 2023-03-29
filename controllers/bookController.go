package controllers

import (
	"challenge-06/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

// Create Book
func CreateBook(ctx *gin.Context) {
	var newBook = models.Book{}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `INSERT INTO books (title, author, description) VALUES ($1, $2, $3)`

	err := DB.QueryRowContext(ctx, sqlStatement, newBook.Title, newBook.Author, newBook.Description).Err()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusCreated, "New Book Created")

}

// Get All Books
func GetBooks(ctx *gin.Context) {
	var results = []models.Book{}

	sqlStatement := `SELECT * from books`

	rows, err := DB.QueryContext(ctx, sqlStatement)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		results = append(results, book)
	}

	ctx.JSON(http.StatusOK, results)
}

// Get Book by ID
func GetBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `SELECT id, title, author, description FROM books WHERE id = $1`

	row := DB.QueryRowContext(ctx, sqlStatement, id)
	book := models.Book{}

	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "NOT FOUND",
			"message": fmt.Sprintf("Book with ID %d not found", id),
		})
	} else if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, book)
	}
}

// UpdateBook
func UpdateBook(ctx *gin.Context) {
	bookRequest := models.Book{}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := ctx.ShouldBindJSON(&bookRequest); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1
	`

	res, err := DB.ExecContext(ctx, sqlStatement, id, bookRequest.Title, bookRequest.Author, bookRequest.Description)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if count, _ := res.RowsAffected(); count != 0 {
		ctx.String(http.StatusOK, "Updated")
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"status":  "NOT FOUND",
		"message": fmt.Sprintf("Book with ID %d not found", id),
	})
}

// Delete Book
func DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	DELETE FROM books
	WHERE id = $1
	`

	res, err := DB.ExecContext(ctx, sqlStatement, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if count, _ := res.RowsAffected(); count != 0 {
		ctx.String(http.StatusOK, "Deleted")
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"status":  "NOT FOUND",
		"message": fmt.Sprintf("Book with ID %d not found", id),
	})
}
