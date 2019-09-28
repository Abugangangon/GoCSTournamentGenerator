package Services

import (
	"fmt"
	"github.com/Abugangangon/go-rest-api/Database"
	"github.com/Abugangangon/go-rest-api/Model"
	"log"
)

func AddGroups(tournamentId int) {
	db, err := Database.Connect()
	defer db.Close()

	sqlStatement := `
		INSERT INTO public."Group" ("Tournament_Id")
		VALUES 
		($1),($1),($1),($1),($1),($1),($1),($1),
		($1),($1),($1),($1),($1),($1),($1),($1)
		RETURNING "Id"`
	id := 0
	err = db.QueryRow(sqlStatement, tournamentId).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tournament groups created for Tournament Id:", tournamentId)
}

func GetGroupsByTournamentId(tournamentId int) []Model.Group {
	db, err := Database.Connect()
	defer db.Close()

	var groups []Model.Group

	sqlStatement := `Select * from public."Group" where "Tournament_Id" = ($1)`

	rows, err := db.Query(sqlStatement, tournamentId)
	for rows.Next() {
		var group Model.Group
		if err := rows.Scan(&group.Id, &group.GroupNumber, &group.TournamentId); err != nil {
			log.Fatal(err)
		}
		groups = append(groups, group)
	}
	if err != nil {
		panic(err)
	}

	for i:= 0; i<len(groups); i++ {
		fmt.Println("Id =", groups[i].Id," Number =", groups[i].GroupNumber, " TournamentId =", groups[i].TournamentId)
	}

	return groups
}

func GetGroupsToShowByTournamentId(tournamentId int) []Model.FaseDeGrupos {
	db, err := Database.Connect()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var FaseDeGrupos []Model.FaseDeGrupos

	sqlStatement := `select "Tournament"."Name", "Group"."Number", "Team"."Name", "Score"."Points" from "Score" join 
		"Group" on "Score"."Group_Id" = "Group"."Id" join "Tournament" on "Tournament"."Id" = "Group"."Tournament_Id" 
		join "Team" on "Team"."Id" = "Score"."Team_Id" 
		where "Tournament_Id" = ($1) order by "Group"."Number", "Score"."Points" desc`

	rows, err := db.Query(sqlStatement, tournamentId)
	for rows.Next() {
		var group Model.FaseDeGrupos
		if err := rows.Scan(&group.TournamentName, &group.GroupNumber, &group.TeamName, &group.GroupPoints); err != nil {
			log.Fatal(err)
		}
		FaseDeGrupos = append(FaseDeGrupos, group)
	}
	if err != nil {
		panic(err)
	}

	return FaseDeGrupos
}