package controllers

import (
	"fmt"
	"math"
	"net/http"
	"quiz-3/structs"

	"github.com/gin-gonic/gin"
)

func HitungLingkaran(c *gin.Context) {
	var lingkaran structs.Lingkaran

	// Mengikat nilai dari query string ke struct Lingkaran
	if err := c.ShouldBindQuery(&lingkaran); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if lingkaran.Hitung == "luas" {
		// Menghitung luas lingkaran
		luas := 3.14 * (math.Pow(float64(lingkaran.Jarijari), 2))
		fmt.Println("Luas Lingkaran", luas)

		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"jari-jari": lingkaran.Jarijari,
			"luas":      luas,
		})
	} else if lingkaran.Hitung == "keliling" {

		var keliling int = int(float64(2) * 3.14 * float64(lingkaran.Jarijari))
		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"jari-jari": lingkaran.Jarijari,
			"keliling":  keliling,
		})
		fmt.Println("Keliling Lingkaran", keliling)
	}

}
