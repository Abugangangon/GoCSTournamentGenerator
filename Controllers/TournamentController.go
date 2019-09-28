package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Abugangangon/go-rest-api/Services"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetOneTournament(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err!=nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(Services.GetTournament(eventID))
}

func GetAllTournaments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Services.GetTournaments())
}

func CreateTournament(w http.ResponseWriter, r *http.Request) {
	var tournamentName string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the tournament title in order to create")
	}

	json.Unmarshal(reqBody, &tournamentName)
	Services.AddTournament(tournamentName)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(tournamentName)
}