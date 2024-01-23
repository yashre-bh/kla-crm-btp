package controller

import (
	"fmt"
	"net/http"

	"github.com/yashre-bh/kla-crm-btp/pkg/models"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "backend connection established")
}

func CheckDBConnection(w http.ResponseWriter, r *http.Request) {
	_, err := models.Connection()
	if err != nil {
		fmt.Fprintln(w, err)
	}
}
