package routes

import (
	"HealthChain_API/controllers"
	"HealthChain_API/controllers/user"
	"HealthChain_API/middleware"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.AuthMiddleWare)
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/login", user.Login).Methods("POST")
	api.HandleFunc("/register", user.Register).Methods("POST")
	api.HandleFunc("/forgot-password", user.ForgotPwController).Methods("POST")
	api.HandleFunc("/reset-password", user.ResetPwController).Methods("POST")

	//api.HandleFunc("/patients", middleware.Auth(controllers.GetPatiens)).Methods("GET")
	//api.HandleFunc("/patients/{id}", middleware.Auth(controllers.GetPatient)).Methods("GET")
	//api.HandleFunc("/patients", middleware.Auth(controllers.CreatePatient)).Methods("POST")
	//api.HandleFunc("/patients/{id}", middleware.Auth(controllers.UpdatePatient)).Methods("PUT")
	//api.HandleFunc("/patients/{id}", middleware.Auth(controllers.DeletePatient)).Methods("DELETE")

	api.HandleFunc("/add-access", controllers.AddAccess).Methods("POST")
	api.HandleFunc("/get-access/{user_id}", controllers.GetAccess).Methods("GET")
	api.HandleFunc("/delete-access/{access_id}", controllers.DeleteAccess).Methods("DELETE")

	//api.HandleFunc("/audit", middleware.Auth(controllers.GetAuditLogs)).Methods("GET")

	return r
}
