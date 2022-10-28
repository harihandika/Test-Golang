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

	r.HandleFunc("/todo-items", h.FindTodo).Methods("GET")
	r.HandleFunc("/todo-items/{id}", h.GetTodo).Methods("GET")
	r.HandleFunc("/todo-items", h.CreateTodo).Methods("POST")
	r.HandleFunc("/todo-items/{id}", h.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todo-items/{id}", h.DeleteTodo).Methods("DELETE")
}
