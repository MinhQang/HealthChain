package routes

import (
	"HealthChain_API/controllers"
	"HealthChain_API/middleware"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.AuthMiddleWare)
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/login", controllers.Login).Methods("POST")
	api.HandleFunc("/register", controllers.Register).Methods("POST")
	api.HandleFunc("/forgot-password", controllers.ForgotPwController).Methods("POST")
	api.HandleFunc("/reset-password", controllers.ResetPwController).Methods("POST")
	api.HandleFunc("/verify-otp", controllers.VerifyOTPController).Methods("POST")

	api.HandleFunc("/patients", controllers.CreatePatientsController).Methods("POST")
	api.HandleFunc("/patients/{id}", controllers.GetPatientsController).Methods("GET")
	api.HandleFunc("/patients/{id}", controllers.UpdatePatientsController).Methods("PUT")
	api.HandleFunc("/patients/{id}", controllers.DeletePatientsController).Methods("DELETE")

	api.HandleFunc("/add-access", controllers.AddAccess).Methods("POST")
	api.HandleFunc("/get-access/{user_id}", controllers.GetAccess).Methods("GET")
	api.HandleFunc("/delete-access/{access_id}", controllers.DeleteAccess).Methods("DELETE")
	return r
}
