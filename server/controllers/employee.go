package controllers

import (
	"fmt"
	apss "go-web-template/server/apps/web"
	"go-web-template/server/params/employee"
	"go-web-template/server/servies"
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

type EmployeeController struct {
	Service servies.EmployeeService
}

func NewEmployeeController(service servies.EmployeeService) *EmployeeController {
	return &EmployeeController{Service: service}
}

func (e *EmployeeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	employees, err := e.Service.GetAllEmployees()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	web := apss.RenderWeb{
		Title: "Halaman Employee",
		Data:  employees,
	}

	err = tmpl.Execute(w, web)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (e *EmployeeController) Add(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/add.html"), utils.LayoutMaster)
		if err != nil {
			log.Printf("Failed to parse template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		web := apss.RenderWeb{
			Title: "Tambah Pegawai",
		}

		if err := tmpl.Execute(w, web); err != nil {
			log.Printf("Failed to execute template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:

		if err := r.ParseForm(); err != nil {
			log.Printf("Failed to parse form: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}


		request := params.EmployeeCreate{
			NIP:     r.Form.Get("nip"),
			Address: r.Form.Get("address"),
			Name:    r.Form.Get("name"),
		}


		_, err := e.Service.CreateEmployee(&request)
		if err != nil {
			log.Printf("Failed to create employee: %v", err)
			http.Error(w, fmt.Sprintf("Failed to create employee: %v", err), http.StatusInternalServerError)
			return
		}


		http.Redirect(w, r, "/employees/add?status=success", http.StatusSeeOther)

	default:

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (e *EmployeeController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		e.renderUpdateForm(w, r)
	case http.MethodPost:

		e.processUpdateByID(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}


func (e *EmployeeController) renderUpdateForm(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}


	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Failed to parse ID:", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	employee, err := e.Service.GetEmployeeByID(id)
	if err != nil {
		log.Println("Failed to get employee:", err)
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/update.html"), utils.LayoutMaster)
	if err != nil {
		log.Println("Failed to parse template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	web := apss.RenderWeb{
		Title: "Update Employee",
		Data:  employee,
	}

	if err := tmpl.Execute(w, web); err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}


func (e *EmployeeController) processUpdateByID(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Failed to parse form:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Failed to parse ID:", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}


	request := params.EmployeeUpdate{
		ID:      id,
		NIP:     r.Form.Get("nip"),
		Address: r.Form.Get("address"),
		Name:    r.Form.Get("name"),
	}


	_, err = e.Service.UpdateEmployeeByID(&request)
	if err != nil {
		log.Println("Failed to update employee:", err)
		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
		return
	}


	http.Redirect(w, r, "/employees?status=success", http.StatusSeeOther)
}


func (e *EmployeeController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}


	idStr := r.FormValue("id")
	if idStr == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}


	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}


	err = e.Service.DeleteEmployeeByID(id)
	if err != nil {
		log.Println("Failed to delete employee:", err)
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}


	http.Redirect(w, r, "/employees?status=success", http.StatusSeeOther)
}
