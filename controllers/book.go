package controllers

import (
	"fmt"
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"regexp"
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

	if !isValidURL(book.Image_url) && !isValidYear(book.Release_year) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": []string{
				"URL Tidak Tidak Valid!",
				"Tahun Buku Tidak Valid! Min 1980 Maks 2021",
			},
		})
		return
	} else if !isValidURL(book.Image_url) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": "URL Tidak Tidak Valid!",
		})
		return
	} else if !isValidYear(book.Release_year) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": "Tahun Buku Tidak Valid! Min 1980 Maks 2021",
		})
		return
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	book.Created_at = formattedTime
	book.Updated_at = formattedTime

	book.Thickness = getThickness(book.Total_page)
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

	if !isValidURL(book.Image_url) && !isValidYear(book.Release_year) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": []string{
				"URL Tidak Tidak Valid!",
				"Tahun Buku Tidak Valid! Min 1980 Maks 2021",
			},
		})
		return
	} else if !isValidURL(book.Image_url) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": "URL Tidak Tidak Valid!",
		})
		return
	} else if !isValidYear(book.Release_year) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": "Tahun Buku Tidak Valid! Min 1980 Maks 2021",
		})
		return
	}

	book.Id = int64(id)
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	book.Updated_at = formattedTime
	book.Thickness = getThickness(book.Total_page)

	fmt.Println(book.Thickness)
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

func isValidURL(url string) bool {
	re := regexp.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`)
	return re.MatchString(url)
}

func isValidYear(strYear int64) bool {
	if strYear >= 1980 && strYear <= 2021 {
		return true
	} else {
		return false
	}
}

func getThickness(totalPage int64) string {
	if totalPage <= 100 {
		return "Tipis"
	} else if totalPage > 100 && totalPage <= 200 {
		return "Sedang"
	} else {
		return "Tebal"
	}
}
