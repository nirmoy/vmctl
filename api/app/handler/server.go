package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nirmoy/vmctl/pkg/cloud/dummy"
)

func GetAllServer(w http.ResponseWriter, r *http.Request) {
	allServer, err := dummy.GetAllServer()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, allServer)
}

func GetServerByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if dummy.IsExistServerByUUID(uuid) {
		server, err := dummy.GetServerByUUID(uuid)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, server)
		return
	}

	respondJSON(w, http.StatusNotFound, nil)
}

func GetServerStatusByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if dummy.IsExistServerByUUID(uuid) {
		serverStatus, err := dummy.GetServerStatusByUUID(uuid)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, serverStatus)
		return
	}

	respondJSON(w, http.StatusNotFound, nil)
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
func DeleteServerByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	success, _ := dummy.DeleteServerByUUID(uuid)
	if !success {
		respondJSON(w, http.StatusInternalServerError, nil)
		return
	}

	respondJSON(w, http.StatusNoContent, nil)
}

func CheckServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if dummy.IsExistServerByName(name) {
		if dummy.IsProhibitedServer(name) {
			respondJSON(w, http.StatusForbidden, nil)
			return
		}

		respondJSON(w, http.StatusOK, nil)
		return
	}

	respondJSON(w, http.StatusInternalServerError, nil)
}
