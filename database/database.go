package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
	"github.com/joho/godotenv"
	"os"
)

var DB *gorm.DB

func Connect() {
	// Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: gagal load .env file")
    }

    // Ambil variabel dari env
    db_user := os.Getenv("DB_USER")
    db_pass := os.Getenv("DB_PASS")
    db_name := os.Getenv("DB_NAME")
    db_host := os.Getenv("DB_HOST")
    db_port := os.Getenv("DB_PORT")

    dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal koneksi ke database:", err)
    }

    DB = db
    log.Println("Berhasil koneksi ke database")
}
