package Services

import (
	"database/sql"
	"fmt"
	"github.com/Abugangangon/go-rest-api/Database"
	"github.com/Abugangangon/go-rest-api/Model"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func AddGroupGames(tournamentId int) {
	db, err := Database.Connect()
	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	var scores = GetScoreByTournamentId(tournamentId)

	var sqlGameInsert strings.Builder
	var sqlGroupUpdate strings.Builder

	sqlGameInsert.WriteString(`
		INSERT INTO public."Game" ("Tournament_Id", "Is_Group", "Team1_Id", "Team1_Score", "Team2_Id", "Team2_Score")
		VALUES `)

	sqlGroupUpdate.WriteString(`update public."Score" set "Points" = ($2) where "Team_Id" = ($1)`)

	for i:=0; i< len(scores); i++ {
		for j:=i+1; j<i+5 && j<len(scores); j++{
			if scores[i].GroupId != scores[j].GroupId {
				break
			}
			var gameResult = playGame()
			if gameResult[0] > gameResult[1] {
				scores[i].Points++
				_, err := db.Exec(sqlGroupUpdate.String(), scores[i].TeamId, scores[i].Points)
				if err != nil {
					panic(err)
				}
			} else if gameResult[1] > gameResult[0] {
				scores[j].Points++
				_, err := db.Exec(sqlGroupUpdate.String(), scores[j].TeamId, scores[j].Points)
				if err != nil {
					panic(err)
				}
			}

			sqlGameInsert.WriteString("(")
			sqlGameInsert.WriteString(strconv.Itoa(tournamentId))
			sqlGameInsert.WriteString(",")
			sqlGameInsert.WriteString("true")
			sqlGameInsert.WriteString(",")
			sqlGameInsert.WriteString(strconv.Itoa(scores[i].TeamId))
			sqlGameInsert.WriteString(",")
			sqlGameInsert.WriteString(strconv.Itoa(gameResult[0]))
			sqlGameInsert.WriteString(",")
			sqlGameInsert.WriteString(strconv.Itoa(scores[j].TeamId))
			sqlGameInsert.WriteString(",")
			sqlGameInsert.WriteString(strconv.Itoa(gameResult[1]))
			if i+2<len(scores) {
				sqlGameInsert.WriteString("),")
			} else {
				sqlGameInsert.WriteString(")")
			}
			fmt.Println(scores[i].TeamId,": ", gameResult[0], " -",gameResult[1],":", scores[j].TeamId)
		}
	}

	_, err = db.Query(sqlGameInsert.String())
	if err != nil {
		panic(err)
	}

	fmt.Println("Games from group stage registered for Tournament Id:", tournamentId)
	fmt.Println("Group points updated for Tournament Id:", tournamentId)
}

func GetGroupGamesByTournamentId(tournamentId int) []Model.Game {
	db, err := Database.Connect()
	defer db.Close()

	var games []Model.Game

	sqlStatement := `select * from public."Game" where "Is_Group" = ($2) and "Tournament_Id" = ($1)`

	rows, err := db.Query(sqlStatement, tournamentId, "true")
	for rows.Next() {
		var game Model.Game
		if err := rows.Scan(&game.Id, &game.TournamentId, &game.IsGroup, &game.Team1, &game.Score1, &game.Team2, &game.Score2); err != nil {
			log.Fatal(err)
		}
		games = append(games, game)
	}
	if err != nil {
		panic(err)
	}

	for i:= 0; i<len(games); i++ {
		fmt.Println("Team",games[i].Team1,"|", games[i].Score1," - ", games[i].Score2,"|", " Team", games[i].Team2)
	}

	return games
}

func GetGamesToShowByTournamentId(tournamentId int) []Model.TeamAndMatch {
	db, err := Database.Connect()
	defer db.Close()

	var games []Model.TeamAndMatch

	sqlStatement := `select "Team1"."Name", "Game"."Team1_Score", "Team2"."Name", "Game"."Team2_Score" from "Game" 
		inner join "Team" as "Team1" on "Game"."Team1_Id" = "Team1"."Id" 
		inner join "Team" as "Team2" on "Game"."Team2_Id" = "Team2"."Id" 
		where "Is_Group" = ($2) and "Tournament_Id" = ($1)`

	rows, err := db.Query(sqlStatement, tournamentId, "true")
	for rows.Next() {
		var game Model.TeamAndMatch
		if err := rows.Scan(&game.Team1Name, &game.Team1Score, &game.Team2Name, &game.Team2Score); err != nil {
			log.Fatal(err)
		}
		games = append(games, game)
	}
	if err != nil {
		panic(err)
	}

	return games
}

func playGame() [2]int {
	var gameScore [2]int
	for i:=0; i<30; i++ {
		gameScore[rand.Intn(2)]++
		if gameScore[0] == 16 || gameScore[1] == 16	{
			break
		}
	}
	return gameScore
}

func eliminationGame() [2]int {
	var gameScore [2]int
	elimination := 16
	for {
		gameScore[rand.Intn(2)]++
		if gameScore[0] == elimination || gameScore[1] == elimination	{
			break
		} else if gameScore[0] == gameScore[1] && gameScore[0] == elimination-1{
			elimination = elimination+3
		}
	}
	return gameScore
}

func GetSumMatchScore(tournamentId int, teamId int) int {
	db, err := Database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sum := 0
	sqlTeam1Query := `select sum("Team1_Score") from public."Game" where "Team1_Id" = ($1) and "Tournament_Id" = ($2)`
	sqlTeam2Query := `select sum("Team2_Score") from public."Game" where "Team2_Id" = ($1) and "Tournament_Id" = ($2)`

	rows, err := db.Query(sqlTeam1Query, teamId, tournamentId)
	for rows.Next() {
		var score sql.NullInt64
		if err := rows.Scan(&score); err != nil {
			log.Fatal(err)
		}
		sum = sum + int(score.Int64)
	}
	if err != nil {
		panic(err)
	}
	rows, err = db.Query(sqlTeam2Query, teamId, tournamentId)
	for rows.Next() {
		var score sql.NullInt64
		if err := rows.Scan(&score); err != nil {
			log.Fatal(err)
		}
		sum = sum + int(score.Int64)
	}
	if err != nil {
		panic(err)
	}

	return sum
}
