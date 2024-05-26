package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"quiz-3/controllers"
	"quiz-3/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")

	host := os.Getenv("DB_HOST")
	strPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, _ := strconv.Atoi(strPort)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {

		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(db)

	defer db.Close()

	router := gin.Default()
	// Membuat API endpoint bangun-datar
	router.GET("/segitiga-sama-sisi", controllers.HitungSegitiga)
	router.GET("/persegi", controllers.HitungPersegi)
	router.GET("/persegi-panjang", controllers.HitungPersegiPanjang)
	router.GET("/lingkaran", controllers.HitungLingkaran)

	// CRUD Category
	router.GET("/categories", controllers.GetAllCategory)
	router.GET("/categories/:id/books", controllers.GetAllBookByCategories)

	// CRUD Books
	router.GET("/books", controllers.GetAllBook)

	// Group route dengan middleware BasicAuth
	authorized := router.Group("/", auth())
	{
		// Sub-group untuk books
		books := authorized.Group("/books")
		{
			books.POST("/", controllers.InsertBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		// Sub-group untuk authors
		categories := authorized.Group("/categories")
		{

			categories.POST("/", controllers.InsertCategory)
			categories.PUT("/:id", controllers.UpdateCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
		}
	}

	router.Run(envPortOr("8080"))
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}

// Fungi Log yang berguna sebagai middleware
func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, pwd, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password tidak boleh kosong"})
			c.Abort()
			return
		}

		if (uname == "admin" && pwd == "password") || (uname == "editor" && pwd == "secret") {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak Tidak Memiliki Hak Akses"})
		c.Abort()
		return
	}
}
