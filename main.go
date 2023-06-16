package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Struktur model untuk tabel transaksi
type Transaction struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	Quantity   int
}

func main() {
	// Inisialisasi koneksi database
	dsn := "root:root@tcp(localhost:3306)/database_name?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrasi tabel transaksi
	db.AutoMigrate(&Transaction{})

	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Handler untuk membuat data transaksi
	app.Post("/transactions", func(c *fiber.Ctx) error {
		// Parsing data transaksi dari body request
		var transaction Transaction
		if err := c.BodyParser(&transaction); err != nil {
			return err
		}

		// Menyimpan data transaksi ke database
		result := db.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}

		// Mengembalikan response sukses
		return c.JSON(fiber.Map{
			"message": "Transaction created successfully",
			"data":    transaction,
		})
	})

	// Menjalankan server aplikasi
	app.Listen(":3000")
}
