package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jayisaac0/auth-service/src/domain/entity"
)

// Users array
type users []entity.User

// UserHandler array
type UserHandler struct{}

// initialMigration and run
func initialMigration() {
	defer db.Close()
	db.AutoMigrate(&entity.User{})
}

// create user
func (u *UserHandler) create(w http.ResponseWriter, r *http.Request) {

	var user entity.User

	json.NewDecoder(r.Body).Decode(&user)
	err := user.Validate()
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		return
	}

	db.Create(&user)
	json.NewEncoder(w).Encode(&user)
	return
}

// fetchAll user
func (u *UserHandler) fetchAll(w http.ResponseWriter, r *http.Request) {
	var users []entity.User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
	return
}

// fetchRecord user
func (u *UserHandler) fetchRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user entity.User

	if err := db.Where("ID = ?", params["ID"]).First(&user).Error; err != nil {
		// error handling...
		fmt.Fprint(w, err)
		return
	}

	json.NewEncoder(w).Encode(&user)
	return
}

// updateRecord user
func (u *UserHandler) updateRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user entity.User
	var updateuser entity.UserUpdate

	json.NewDecoder(r.Body).Decode(&updateuser)

	err := updateuser.Validate()
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		return
	}

	db.Model(&user).Where("ID = ?", params["ID"]).Update(&updateuser)
	json.NewEncoder(w).Encode("Record updated")
	return
}

// deleteRecord user
func (u *UserHandler) deleteRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user entity.User

	db.Where("ID = ?", params["ID"]).Delete(&user)

	var users []entity.User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
