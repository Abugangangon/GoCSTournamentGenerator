package Services

import (
	"fmt"
	"github.com/Abugangangon/go-rest-api/Database"
	"github.com/Abugangangon/go-rest-api/Model"
	"log"
	"strconv"
	"strings"
)

func AddScore(tournamentId int) {
	db, err := Database.Connect()
	defer db.Close()

	var teams = GetTeams()
	var groups = GetGroupsByTournamentId(tournamentId)
	groupLimit := 0

	var sqlStatement strings.Builder

	sqlStatement.WriteString(`
		INSERT INTO public."Score" ("Team_Id", "Group_Id")
		VALUES `)

	for i:=0; i< len(teams); i++ {
		if i != 0 && i % 5 == 0 {
			groupLimit++
		}
		sqlStatement.WriteString("(")
		sqlStatement.WriteString(strconv.Itoa(teams[i].Id))
		sqlStatement.WriteString(",")
		sqlStatement.WriteString(strconv.Itoa(groups[groupLimit].Id))
		if i+1<len(teams) {
			sqlStatement.WriteString("),")
		} else {
			sqlStatement.WriteString(")")
		}
	}

	fmt.Println(sqlStatement.String())

	db.Query(sqlStatement.String())
	if err != nil {
		panic(err)
	}

	fmt.Println("Score boards for groups created for Tournament Id:", tournamentId)
}

func GetScoreByTournamentId(tournamentId int) []Model.Score {
	db, err := Database.Connect()
	defer db.Close()

	var scores []Model.Score

	sqlStatement := `select public."Score"."Id", "Team_Id", "Group_Id", "Points" from public."Score" join public."Group"
 		on public."Group"."Id" = "Group_Id" where "Tournament_Id" = ($1) order by "Group_Id" asc, "Points" desc`

	rows, err := db.Query(sqlStatement, tournamentId)
	for rows.Next() {
		var score Model.Score
		if err := rows.Scan(&score.Id, &score.TeamId, &score.GroupId, &score.Points); err != nil {
			log.Fatal(err)
		}
		scores = append(scores, score)
	}
	if err != nil {
		panic(err)
	}

	for i:= 0; i<len(scores); i++ {
		fmt.Println("ScoreId =",scores[i].Id," TeamId =", scores[i].Id," GroupId =", scores[i].GroupId,
			" Points =", scores[i].Points)
	}

	return scores
}

func GetGroupWinners(tournamentId int) []Model.Score {
	db, err := Database.Connect()
	defer db.Close()

	var scores []Model.Score
	var winners []Model.Score
	var winner Model.Score
	winnersAux := 0

	sqlStatement := `select public."Score"."Id", "Team_Id", "Group_Id", "Points" from public."Score" join public."Group"
 		on public."Group"."Id" = "Group_Id" where "Tournament_Id" = ($1) order by "Group_Id" asc, "Points" desc`

	rows, err := db.Query(sqlStatement, tournamentId)
	for rows.Next() {
		var score Model.Score
		if err := rows.Scan(&score.Id, &score.TeamId, &score.GroupId, &score.Points); err != nil {
			log.Fatal(err)
		}
		scores = append(scores, score)
	}
	if err != nil {
		panic(err)
	}

	for i:= 0; i+4<len(scores); i=i+5 {
		winnersAux = len(winners)
		winner = scores[i]
		for j:=i; j<i+5 && j<len(scores); j++ {
			if len(winners) >= winnersAux+2 {
				break
			}
			if winner.Points > scores[j+1].Points && winner.GroupId == scores[j+1].GroupId {
				winners = append(winners, winner)
				winner = scores[j+1]
			} else if winner.Points == scores[j+1].Points && scores[j].GroupId == scores[j+1].GroupId {
				if GetSumMatchScore(tournamentId, winner.TeamId) < GetSumMatchScore(tournamentId, scores[j+1].TeamId) {
					winner = scores[j+1]
				}
			} else {
				winners = append(winners, winner)
			}
		}
	}

	for i:= 0; i<len(winners); i++ {
		fmt.Println("ScoreId =",winners[i].Id," TeamId =", winners[i].Id," GroupId =", winners[i].GroupId,
			" Points =", winners[i].Points)
	}

	return winners
}
