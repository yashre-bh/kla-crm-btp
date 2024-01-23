package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddEmployee(writer http.ResponseWriter, request *http.Request) {
	var employee types.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = models.AddEmployee(employee)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Failed to add new employee", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("User added successfully")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(employee)

}

func FetchAllEmployees(writer http.ResponseWriter, request *http.Request) {
	employees, err := models.FetchAllEmployees()
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Failed to fetch all employees", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(employees)

}
