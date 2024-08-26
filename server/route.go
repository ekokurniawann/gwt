package routes

import (
	"database/sql"
	"go-web-template/server/controllers"
	repositories "go-web-template/server/repositories/employee"
	"go-web-template/server/servies"
	"net/http"

	"github.com/gorilla/mux"
)


func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	
	employeeRepo := repositories.NewEmployeeRepository(db)

	
	employeeService := servies.NewEmployeeService(employeeRepo)

	
	employeeController := controllers.NewEmployeeController(*employeeService)
	homeController := controllers.NewHomeController()

	
	router.HandleFunc("/employees", employeeController.Index).Methods(http.MethodGet)
	router.HandleFunc("/employees/add", employeeController.Add).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/employees/update", employeeController.UpdateByID).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/employees/delete", employeeController.DeleteByID).Methods(http.MethodPost)

	router.HandleFunc("/", homeController.Index)

	return router
}
