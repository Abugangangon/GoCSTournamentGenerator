package Services

import (
	"fmt"
	"github.com/Abugangangon/go-rest-api/Database"
	"github.com/Abugangangon/go-rest-api/Model"
	"log"
)

func AddTournament(name string) {
	db, err := Database.Connect()
	defer db.Close()

	sqlStatement := `
		INSERT INTO public."Tournament" ("Name")
		VALUES ($1)
		RETURNING "Id"`
	id := 0
	err = db.QueryRow(sqlStatement, name).Scan(&id)
	if err != nil {
		panic(err)
	}

	AddGroups(id)
	AddScore(id)
	AddGroupGames(id)
	AddPlayoffTeams(id)

	fmt.Println("New Tournament ID is:", id)
}

func RemoveTournament(idToRemove int) {
	db, err := Database.Connect()
	defer db.Close()

	sqlStatement := `
		delete from public."Tournament"
		where "Id" = $1
		RETURNING "Id"`
	id := 0
	err = db.QueryRow(sqlStatement, idToRemove).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted record ID is:", id)
}

func UpdateTournament(tournament Model.Tournament) {
	db, err := Database.Connect()
	defer db.Close()

	sqlStatement := `
		update public."Tournament"
		set "Name" = $2
		where "Id" = $1
		RETURNING "Id"`
	id := 0
	err = db.QueryRow(sqlStatement, tournament.Id, tournament.Name).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated record ID is:", id)
}

func GetTournament(id int) Model.Tournament {
	db, err := Database.Connect()
	defer db.Close()
	var tournament Model.Tournament

	sqlStatement := `Select * from public."Tournament" where "Id" = $1`

	err = db.QueryRow(sqlStatement, id).Scan(&tournament.Id, &tournament.Name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Id =", tournament.Id," Name =", tournament.Name)

	return tournament
}

func GetTournaments() []Model.Tournament {
	db, err := Database.Connect()
	defer db.Close()

	var tournaments []Model.Tournament

	sqlStatement := `Select * from public."Tournament" LIMIT 80`

	rows, err := db.Query(sqlStatement)
	for rows.Next() {
		var tournament Model.Tournament
		if err := rows.Scan(&tournament.Id, &tournament.Name); err != nil {
			log.Fatal(err)
		}
		tournaments = append(tournaments, tournament)
	}
	if err != nil {
		panic(err)
	}

	for i:= 0; i<len(tournaments); i++ {
		fmt.Println("Id =", tournaments[i].Id," Name =", tournaments[i].Name)
	}

	return tournaments
}