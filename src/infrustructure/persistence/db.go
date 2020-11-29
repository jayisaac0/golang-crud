package persistence

import "github.com/jinzhu/gorm"

var db *gorm.DB
var err error

// DB function
func DB() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_test_db sslmode=disable password=postgres")
	if err != nil {
		panic("could not connect to database")
	}
	defer db.Close()
}
