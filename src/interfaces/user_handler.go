package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jayisaac0/auth-service/src/infrustructure/utility"

	"github.com/gorilla/mux"
	"github.com/jayisaac0/auth-service/src/domain/entity"
)

// Users array
type users []entity.User

// usersyHandler array
type userHandler struct{}

// initialMigration and run
func initialMigration() {
	defer db.Close()
	db.AutoMigrate(&entity.User{})
}

func (u *userHandler) CreateCountry(w http.ResponseWriter, r *http.Request) {
	var request entity.User

	am := utility.Validator(r.Body, &request,
		validation.Field(&request.Username, validation.Required, validation.Length(4, 50)),
		validation.Field(&request.Email, validation.Required, is.Email, validation.Length(4, 50)),
		validation.Field(&request.Password, validation.Required, validation.Length(4, 50)),
	)

	if am != nil {
		log.Printf("Something went wrong with validation: %v\n", am)
		return
	}
}

// create user
func (u *userHandler) create(w http.ResponseWriter, r *http.Request) {

	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	// err := user.Validate()
	// if err != nil {
	// 	json.NewEncoder(w).Encode(&err)
	// 	return
	// }

	db.Create(&user)
	json.NewEncoder(w).Encode(&user)
	return
}

// fetchAll user
func (u *userHandler) fetchAll(w http.ResponseWriter, r *http.Request) {
	var users []entity.User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
	return
}

// fetchRecord user
func (u *userHandler) fetchRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user entity.User

	if err := db.Where("ID = ?", params["ID"]).First(&user).Error; err != nil {
		// error handling...
		fmt.Fprint(w, err)
		return
	}

	// db.Where("Username = ?", params["ID"]).First(&user)
	json.NewEncoder(w).Encode(&user)
	return
}

// updateRecord user
func (u *userHandler) updateRecord(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// var updateuser UserUpdate
	// json.NewDecoder(r.Body).Decode(&updateuser)

	// err := updateuser.Validate()
	// if err != nil {
	// 	json.NewEncoder(w).Encode(&err)
	// 	return
	// }

	// db.Model(&updateuser).Where("ID = ?", params["ID"]).Update(&UserUpdate{
	// 	Username: updateuser.Username,
	// 	Email:    updateuser.Email,
	// })
	// return

	params := mux.Vars(r)

	var user entity.User
	// var updateuser UserUpdate

	json.NewDecoder(r.Body).Decode(&user)

	// err := user.Validate()
	// if err != nil {
	// 	json.NewEncoder(w).Encode(&err)
	// 	return
	// }

	db.Model(&user).Where("ID = ?", params["ID"]).Update(&user)
	return
}

// deleteRecord user
func (u *userHandler) deleteRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user entity.User

	db.Where("ID = ?", params["ID"]).Delete(&user)

	var users []entity.User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
