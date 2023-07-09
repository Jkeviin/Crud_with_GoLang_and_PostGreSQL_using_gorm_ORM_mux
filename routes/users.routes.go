package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Jkeviin/go-gorm-rest-api/db"
	"github.com/Jkeviin/go-gorm-rest-api/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	if len(users) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No hay usuarios")
		return
	}

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 { // Si no se encontró el usuario
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No se encontró el usuario")
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User // Esto es un modelo de la base de datos
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)

	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// Enviar createdUser como respuesta
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 { // Si no se encontró el usuario
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No se encontró el usuario")
		return
	}

	db.DB.Unscoped().Delete(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Usuario eliminado")

}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 { // Si no se encontró el usuario
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No se encontró el usuario")
		return
	}

	json.NewDecoder(r.Body).Decode(&user) // Actualizar el usuario

	db.DB.Save(&user)

	json.NewEncoder(w).Encode(&user)
}
