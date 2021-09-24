package storage

import (
	"fmt"
	"log"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	//Postgres Driver imported
	"database/sql"

	_ "github.com/lib/pq"
)

//Connect to DB
var DB *gorm.DB
var MPosDB *sql.DB
var err error

func ConnectDB() *gorm.DB {
	var (
		host     = "localhost"
		user     = "root"
		port     = 5432
		password = "root"
		name     = "albums"
	)

	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     name,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	DB, err = gorm.Open("postgres", dsn.String())

	//Check for Errors in DB
	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}
	if err := DB.DB().Ping(); err != nil {
		log.Fatalln("Error in make ping the DB " + err.Error())
		return nil
	}
	if DB.Error != nil {
		log.Fatalln("Any Error in connect the DB " + err.Error())
		return nil
	}
	log.Println("DB connected")
	return DB
}
