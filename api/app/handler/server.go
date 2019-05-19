package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nirmoy/vmctl/pkg/cloud/dummy"
)

func GetAllServer(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, dummy.GetAllServers())
}

func CreateServer(w http.ResponseWriter, r *http.Request) {
	server := dummy.Server{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&server); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	if dummy.IsExistServerByName(server.Name) {
		respondError(w, http.StatusConflict, "VM with same name exist")
		return
	}

	server, err := dummy.CreateServer(server.Name)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, dummy.ServerID{ID: server.ID})
}
