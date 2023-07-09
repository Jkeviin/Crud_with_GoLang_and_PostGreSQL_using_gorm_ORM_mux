package main

import (
	"net/http"

	"github.com/Jkeviin/go-gorm-rest-api/db"
	"github.com/Jkeviin/go-gorm-rest-api/models"
	"github.com/Jkeviin/go-gorm-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnect()
	db.DB.AutoMigrate(models.User{}, models.Task{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)


	// Users
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", routes.PutUserHandler).Methods("PUT")

	// Tasks
	r.HandleFunc("/users/{id}/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/users/{id}/tasks/{task_id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/users/{id}/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/users/{id}/tasks/{task_id}", routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}/tasks/{task_id}", routes.PutTaskHandler).Methods("PUT")
	

	http.ListenAndServe(":3000", r)
}
