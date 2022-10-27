package routes

import (
	"golang-test/handlers"
	"golang-test/pkg/mysql"
	"golang-test/repositories"

	"github.com/gorilla/mux"
)

func TodoRoutes(r *mux.Router) {
	todoRepository := repositories.RepositoryTodo(mysql.DB)
	h := handlers.HandlerTodo(todoRepository)

	r.HandleFunc("/activity-groups", h.FindTodo).Methods("GET")
	r.HandleFunc("/activity-groups/{id}", h.GetTodo).Methods("GET")
	r.HandleFunc("/activity-groups", h.CreateTodo).Methods("POST")
	r.HandleFunc("/activity_groups/{id}", h.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/activity_groups/{id}", h.DeleteTodo).Methods("DELETE")
}
