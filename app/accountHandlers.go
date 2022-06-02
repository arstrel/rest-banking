package app

import (
	"encoding/json"
	"net/http"

	"github.com/arstrel/rest-banking/dto"
	"github.com/arstrel/rest-banking/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	request := dto.NewAccountRequest{CustomerId: vars["id"]}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	acc, appErr := ah.service.NewAccount(request)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, acc)

}
