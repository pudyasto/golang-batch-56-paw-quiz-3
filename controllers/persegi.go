package controllers

import (
	"fmt"
	"net/http"
	"quiz-3/structs"

	"github.com/gin-gonic/gin"
)

func HitungPersegi(c *gin.Context) {
	var persegi structs.Persegi

	// Mengikat nilai dari query string ke struct Persegi
	if err := c.ShouldBindQuery(&persegi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if persegi.Hitung == "luas" {
		// Menghitung luas persegi
		luas := persegi.Sisi * (persegi.Sisi)
		fmt.Println("Luas Persegi", luas)

		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"sisi": persegi.Sisi,
			"luas": luas,
		})
	} else if persegi.Hitung == "keliling" {

		var keliling int = 4 * int(persegi.Sisi)
		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"alas":     persegi.Sisi,
			"keliling": keliling,
		})
		fmt.Println("Keliling Persegi", keliling)
	}

}
