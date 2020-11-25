package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB
var err error

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// User struct
type User struct {
	Base
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// connectionError function
func connectionError() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_test_db sslmode=disable password=postgres")
	if err != nil {
		panic("could not connect to database")
	}
	defer db.Close()
}

// initialMigration and run
func initialMigration() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_test_db sslmode=disable password=postgres")
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNECT")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

// Users array
type users []User

// create user
func create(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

// fetchAll user
func fetchAll(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// fetchRecord user
func fetchRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.Where("Username = ?", params["ID"]).First(&user)
	json.NewEncoder(w).Encode(&user)
}

// updateRecord user
func updateRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	db.Model(&user).Where("ID = ?", params["ID"]).Update(&User{
		Username: user.Username,
		Email:    user.Email,
	})

}

// deleteRecord user
func deleteRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["ID"])
	db.Delete(&user)

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// destroyTableuser table
func destroyTable(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Destroy user table")
}

// handleRequets
func handleRequets() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// db, err = gorm.Open("postgres", os.ExpandEnv("host=${HOST} user=${USER} dbname=${DBNAME} sslmode=disable password=${PASSWORD}"))

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_test_db sslmode=disable password=postgres")
	if err != nil {
		panic("could not connect to database")
	}
	defer db.Close()

	myRouter.HandleFunc("/create", create).Methods("POST")
	myRouter.HandleFunc("/fetch", fetchAll).Methods("GET")
	myRouter.HandleFunc("/fetch/{ID}", fetchRecord).Methods("GET")
	myRouter.HandleFunc("/update/{ID}", updateRecord).Methods("PATCH")
	myRouter.HandleFunc("/delete/{ID}", deleteRecord).Methods("DELETE")
	myRouter.HandleFunc("/destroyTable", destroyTable).Methods("DELETE")

	fmt.Println("Server listening to port 8000")

	if err := http.ListenAndServe(":8000", myRouter); err != nil {
		log.Fatal(http.ListenAndServe(":8000", myRouter))
	}
}

func main() {
	initialMigration()
	handleRequets()
}
