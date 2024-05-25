package controllers

import (
	"fmt"
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	book.Created_at = formattedTime
	book.Updated_at = formattedTime

	fmt.Println(book)
	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.Id = int64(id)
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	book.Updated_at = formattedTime

	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book

	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.Id = int64(id)

	err = repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Book",
	})
}
