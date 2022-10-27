package routes

import (
	"golang-test/handlers"
	"golang-test/pkg/mysql"
	"golang-test/repositories"

	"github.com/gorilla/mux"
)

func ActivityRoutes(r *mux.Router) {
	activityRepository := repositories.RepositoryActivity(mysql.DB)
	h := handlers.HandlerActivity(activityRepository)

	r.HandleFunc("/activity-groups", h.FindActivity).Methods("GET")
	r.HandleFunc("/activity-groups/{id}", h.GetActivity).Methods("GET")
	r.HandleFunc("/activity-groups", h.CreateActivity).Methods("POST")
	r.HandleFunc("/activity_groups/{id}", h.UpdateActivity).Methods("PATCH")
	r.HandleFunc("/activity_groups/{id}", h.DeleteActivity).Methods("DELETE")
}
