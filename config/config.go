package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {
	env := ".env"

	errorENV := godotenv.Load(env)

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//	}
	//fmt.Println(dir)
	//	environmentPath := filepath.Join(dir, ".env")
	//	err = godotenv.Load(environmentPath)
	//	panic(err)
	//	fmt.Println(err)
	//errorENV := godotenv.Load()
	//errorENV := godotenv.Load(filepath.Join(path_dir, ".env"))
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(dbUser)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	}

	return db
}

//DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
