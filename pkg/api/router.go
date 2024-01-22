package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yashre-bh/kla-crm-btp/pkg/controller"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/api/connection", controller.GetData).Methods("GET")
	r.HandleFunc("/api/db", controller.CheckDBConnection).Methods("GET")
	r.HandleFunc("/api/add-employee", controller.AddEmployee).Methods("POST")
	r.HandleFunc("/api/fetch-employees", controller.FetchAllEmployees).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
