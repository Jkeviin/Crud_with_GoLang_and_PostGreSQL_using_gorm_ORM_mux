package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Jkeviin/go-gorm-rest-api/db"
	"github.com/Jkeviin/go-gorm-rest-api/models"
	"github.com/gorilla/mux"
)

func getUserByID(userID string) (models.User, error) {
	var user models.User
	result := db.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	user, err := getUserByID(userID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("El usuario no existe")
		return
	}

	/* type Task struct {
		gorm.Model
		Title       string `gorm:"not null;uniqueIndex"`
		Description string
		Done        bool `gorm:"default:false"`
		UserID      uint
	}
	*/

	var tasks []models.Task
	db.DB.Where("user_id = ?", user.ID).Find(&tasks) // se pone user_id en vez de UserID porque en la base de datos se guarda así

	if len(tasks) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("El usuario no tiene tareas")
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]
	taskID := params["task_id"]

	_, err := getUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("El usuario no existe")
		return
	}

	var task models.Task
	db.DB.Where("user_id = ? AND id = ?", userID, taskID).Find(&task)

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("La tarea no existe")
		return
	}

	json.NewEncoder(w).Encode(task)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	user, err := getUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("El usuario no existe")
		return
	}

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	result := db.DB.Where("user_id = ? AND title = ?", userID, task.Title).Find(&task) // Check if task already exists

	if result.RowsAffected > 0 {
		json.NewEncoder(w).Encode("La tarea ya existe")
		return
	}

	task.UserID = user.ID

	createdTask := db.DB.Create(&task)

	err = createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]
	taskID := params["task_id"]

	_, err := getUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("El usuario no existe")
		return
	}

	var task models.Task
	db.DB.Where("user_id = ? AND id = ?", userID, taskID).Find(&task)

	db.DB.Unscoped().Delete(&task) // Unscoped() permite eliminar el registro de la base de datos

	json.NewEncoder(w).Encode(task)
}

func PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	_, err := getUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("El usuario no existe")
		return
	}

	taskID := params["task_id"]

	var task models.Task
	db.DB.Where("user_id = ? AND id = ?", userID, taskID).Find(&task)

	json.NewDecoder(r.Body).Decode(&task)

	db.DB.Save(&task)

	json.NewEncoder(w).Encode(task)
}
