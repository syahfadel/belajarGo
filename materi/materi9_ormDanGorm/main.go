package main

/*
ORM (object-Relational Mapping) merupakan teknik mapping antar struktur data relasional dalam database dengan objek dalam bahasa pemograman
dalam ORM tabel dapat direpresentasikan sebagai kelas atau objek, sehingga memudahkan untuk berinteraksi dengan database tanpa per;i
menulis banyak kode SQL secara manual

ORM pada bahasa go yaitu library dari pihak ketiga seperti GORM.
ORM pada GO dapat membuat structur data (struct) yang merepresentasikan tabel dalam database, kemudian menggunakan pustaka ORM untuk melakukan
pemetaan antara strucr dan tabel.
GORM memungkinkan pengguna untuk melakukan operasi CRUD dmenggunakan method yang sudah disediakan

// install gorm dengan go get -u gorm.io/gorm

CRUD pada GORM
Create untuk create
First untuk Read 1 data, jika banyak gunakan find()
Update untuk update
Delete untuk delete

*/

import (
	"errors"
	"fmt"
	"materi9_ormDanGorm/database"
	"materi9_ormDanGorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("syahfadel@email.com")
	// getUserById(3)
	// updateUserById(3, "update@email.com")
	// createProduct(1, "YLO", "YYYY")
	// getUserWithProducts()
	deletProductById(1)

}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	// method create memerlukan argumen yang memiliki tipe data yang sama dengan yang ingin dibuat
	// argumen akan berisi data yang berhasil diupload ke database
	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Println("New User Data: ", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	// first menerima 3 parameter yaitu pointer terhadap data yang ingin dicari, condition dari query, dan data dari condition
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user: ", err)
	}

	fmt.Printf("User Data := %+v \n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	// memerlukan method model  sehingga hasil update bisa langsung discan
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data: ", err)
		return
	}
	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error craeting product data: ", err.Error())
		return
	}

	fmt.Println("New Product Data: ", Product)
}

func getUserWithProducts() {
	db := database.GetDB()

	users := models.User{}
	// method preload digunakan untuk melakukan join statement antara produts dengan user yang merupakan suatu eager loading
	// method preload memiliki parameter berupa nama tabel yang diawali huruf besar
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products: ", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}

func deletProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product: ", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}
