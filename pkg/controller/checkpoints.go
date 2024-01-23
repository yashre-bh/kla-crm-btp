package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddCheckpoint(writer http.ResponseWriter, request *http.Request) {
	var checkpoint types.Checkpoint
	err := json.NewDecoder(request.Body).Decode(&checkpoint)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = models.AddCheckpoint(checkpoint)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Failed to add new checkpoint", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Checkpoint added successfully")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(checkpoint)
}

func FetchAllCheckpoints(writer http.ResponseWriter, request *http.Request) {
	checkpoints, err := models.FetchAllCheckpoints()
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Failed to fetch all checkpoints", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(checkpoints)

}
