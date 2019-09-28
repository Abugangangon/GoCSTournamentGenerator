package Controllers

import (
	"encoding/json"
	"github.com/Abugangangon/go-rest-api/Services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetGamesAndScores(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err!=nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(Services.GetGamesToShowByTournamentId(eventID))
}
