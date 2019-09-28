package Services

import (
	"fmt"
	"github.com/Abugangangon/go-rest-api/Database"
	"github.com/Abugangangon/go-rest-api/Model"
	"log"
)

func AddTeam(name string) {
	db, err := Database.Connect()

	sqlStatement := `
		INSERT INTO public."Team" ("Name")
		VALUES ($1)
		RETURNING "Id"`
	id := 0
	err = db.QueryRow(sqlStatement, name).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)
	Database.Close(db)
}

func RemoveTeam(idToRemove int) {
	db, err := Database.Connect()

	sqlStatement := `
		delete from public."Team"
		where "Id" = $1
		RETURNING "Id"`
	id := 0
	err = db.QueryRow(sqlStatement, idToRemove).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted record ID is:", id)
	Database.Close(db)
}

func UpdateTeam(team Model.Team) {
	db, err := Database.Connect()

	sqlStatement := `
		update public."Team"
		set "Name" = $2
		where "Id" = $1
		RETURNING "Id"`
	id := 0
	err = db.QueryRow(sqlStatement, team.Id, team.Name).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated record ID is:", id)
	Database.Close(db)
}

func GetTeam(id int) Model.Team {
	db, err := Database.Connect()
	defer db.Close()
	var team Model.Team

	sqlStatement := `Select * from public."Team" where "Id" = $1`

	err = db.QueryRow(sqlStatement, id).Scan(&team.Id, &team.Name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Id =", team.Id," Name =", team.Name)

	return team
}

func GetTeams() []Model.Team {
	db, err := Database.Connect()
	defer db.Close()

	var teams []Model.Team

	sqlStatement := `Select * from public."Team" LIMIT 80`

	rows, err := db.Query(sqlStatement)
	for rows.Next() {
		var team Model.Team
		if err := rows.Scan(&team.Id, &team.Name); err != nil {
			log.Fatal(err)
		}
		teams = append(teams, team)
	}
	if err != nil {
		panic(err)
	}

	for i:= 0; i<len(teams); i++ {
		fmt.Println("Id =", teams[i].Id," Name =", teams[i].Name)
	}

	return teams
}
