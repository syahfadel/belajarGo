package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

// router untuk menghubungkan client kepada endpoint
// func untuk menjalankan server dari aplikasi
// mengembalikan struct *gin.Engine dari Gin untuk menjalankan server sebagai multiplexer dari routing
func StartServer() *gin.Engine {
	// router menampung engine dari router yang kita dapatkan dari function gin.Default
	router := gin.Default()

	// menggunakan method POST untyk menghubungkan client dengan endpoint CreateCar
	// parameter pertama berupa route ath dan kedua  merupakan handler atau endpointnya
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.GET("/cars", controllers.GetCars)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
