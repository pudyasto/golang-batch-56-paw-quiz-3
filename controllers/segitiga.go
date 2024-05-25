package controllers

import (
	"fmt"
	"net/http"
	"quiz-3/structs"

	"github.com/gin-gonic/gin"
)

func HitungSegitiga(c *gin.Context) {
	var segitiga structs.Segitiga

	// Mengikat nilai dari query string ke struct Segitiga
	if err := c.ShouldBindQuery(&segitiga); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if segitiga.Hitung == "luas" {
		// Menghitung luas segitiga
		luas := int(0.5 * float64(segitiga.Alas) * float64(segitiga.Tinggi))
		fmt.Println("Luas Segitiga", luas)

		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"alas":   segitiga.Alas,
			"tinggi": segitiga.Tinggi,
			"luas":   luas,
		})
	} else if segitiga.Hitung == "keliling" {

		var keliling int = int(3 * float64(segitiga.Alas))
		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"alas":     segitiga.Alas,
			"keliling": keliling,
		})
		fmt.Println("Keliling Segitiga", keliling)
	}

}
