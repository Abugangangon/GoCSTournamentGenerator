package main

import (
	"fmt"
	"github.com/Abugangangon/go-rest-api/Controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Primeiro app em GO! Pega Leve na avaliação =D... Os exercicios são os endpoints: /groups/{id} , /games/{id} , /playoffs/{id} , /winner/{id}")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	//TEAMS
	router.HandleFunc("/teams", Controllers.GetAllTeams).Methods("GET")
	router.HandleFunc("/team/{id}", Controllers.GetOneTeam).Methods("GET")
	//Tournaments
	router.HandleFunc("/tournament", Controllers.CreateTournament).Methods("POST")
	router.HandleFunc("/tournaments", Controllers.GetAllTournaments).Methods("GET")
	router.HandleFunc("/tournament/{id}", Controllers.GetOneTournament).Methods("GET")
	//Exercicio 1:
	router.HandleFunc("/groups/{id}", Controllers.GetGroupsAndPoints).Methods("GET")
	router.HandleFunc("/games/{id}", Controllers.GetGamesAndScores).Methods("GET")
	//Exercicio 2:
	router.HandleFunc("/playoffs/{id}", Controllers.GetPlayoffMatchesAndTeams).Methods("GET")
	router.HandleFunc("/winner/{id}", Controllers.GetTournamentWinner).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}