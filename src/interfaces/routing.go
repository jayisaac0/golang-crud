package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// connectionError function
func connectionError() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_test_db sslmode=disable password=postgres")
	if err != nil {
		panic("could not connect to database")
	}
	defer db.Close()
}

// Routing done here
func Routing() http.Handler {
	myRouter := mux.NewRouter().StrictSlash(true)

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_test_db sslmode=disable password=postgres")
	if err != nil {
		panic("could not connect to database")
	}

	api := myRouter.PathPrefix("/api/v1").Subrouter()

	// Country APIs
	userRoutes := api.PathPrefix("/user").Subrouter()
	user := userHandler{}

	userRoutes.HandleFunc("/create", user.create).Methods("POST")
	userRoutes.HandleFunc("/fetch", user.fetchAll).Methods("GET")
	userRoutes.HandleFunc("/fetch/{ID}", user.fetchRecord).Methods("GET")
	userRoutes.HandleFunc("/update/{ID}", user.updateRecord).Methods("PATCH")
	userRoutes.HandleFunc("/delete/{ID}", user.deleteRecord).Methods("DELETE")

	return myRouter
}
