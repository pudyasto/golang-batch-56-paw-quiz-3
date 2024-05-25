package controllers

import (
	"fmt"
	"net/http"
	"quiz-3/structs"

	"github.com/gin-gonic/gin"
)

func HitungPersegiPanjang(c *gin.Context) {
	var persegiPanjang structs.PersegiPanjang

	// Mengikat nilai dari query string ke struct PersegiPanjang
	if err := c.ShouldBindQuery(&persegiPanjang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if persegiPanjang.Hitung == "luas" {
		// Menghitung luas persegiPanjang
		luas := persegiPanjang.Panjang * (persegiPanjang.Lebar)
		fmt.Println("Luas PersegiPanjang", luas)

		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"panjang": persegiPanjang.Panjang,
			"Lebar":   persegiPanjang.Lebar,
			"luas":    luas,
		})
	} else if persegiPanjang.Hitung == "keliling" {

		var keliling int = 2 * int(persegiPanjang.Panjang+persegiPanjang.Lebar)
		// Mengembalikan hasil sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"panjang":  persegiPanjang.Panjang,
			"Lebar":    persegiPanjang.Lebar,
			"keliling": keliling,
		})
		fmt.Println("Keliling PersegiPanjang", keliling)
	}

}
