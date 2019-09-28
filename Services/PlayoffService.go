package Services

import (
	"fmt"
	"github.com/Abugangangon/go-rest-api/Database"
	"github.com/Abugangangon/go-rest-api/Model"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func AddPlayoffTeams(tournamentId int) {
	db, err := Database.Connect()
	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	var sixteenth = GetGroupWinners(tournamentId)
	var octaves []Model.Score
	var quarter []Model.Score
	var semi []Model.Score
	var final []Model.Score
	var winner Model.Score
	var gameResult [2]int

	var sqlPlayoffInsert strings.Builder
	var sqlGameInsert strings.Builder

	sqlPlayoffInsert.WriteString(`
		INSERT INTO public."Playoff" ("Tournament_Id", "Team1_Id", "Team2_Id", "Type")
		VALUES `)

	sqlGameInsert.WriteString(`
		INSERT INTO public."Game" ("Tournament_Id", "Is_Group", "Team1_Id", "Team1_Score", "Team2_Id", "Team2_Score")
		VALUES `)

	fmt.Println("SIXTEENTH OF FINALS MATCHES STARTING NOW!!!!!!")
	for i:=0; i< len(sixteenth); i=i+2 {
		sqlPlayoffInsert.WriteString("(")
		sqlPlayoffInsert.WriteString(strconv.Itoa(tournamentId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(sixteenth[i].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(sixteenth[i+1].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString("'Sixteenth'")
		sqlPlayoffInsert.WriteString("),")

		gameResult = eliminationGame()

		if gameResult[0] > gameResult[1] {
			octaves = append(octaves, sixteenth[i])
		} else {
			octaves = append(octaves, sixteenth[i+1])
		}

		sqlGameInsert.WriteString("(")
		sqlGameInsert.WriteString(strconv.Itoa(tournamentId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString("false")
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(sixteenth[i].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[0]))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(sixteenth[i+1].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[1]))
		sqlGameInsert.WriteString("),")
		fmt.Println(sixteenth[i].TeamId,": ", gameResult[0], " -",gameResult[1],":", sixteenth[i+1].TeamId)
	}

	fmt.Println("OCTAVES OF FINALS MATCHES STARTING NOW!!!!!!")
	for i:=0; i< len(octaves); i=i+2 {
		sqlPlayoffInsert.WriteString("(")
		sqlPlayoffInsert.WriteString(strconv.Itoa(tournamentId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(octaves[i].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(octaves[i+1].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString("'Octaves'")
		sqlPlayoffInsert.WriteString("),")

		gameResult = eliminationGame()

		if gameResult[0] > gameResult[1] {
			quarter = append(quarter, octaves[i])
		} else {
			quarter = append(quarter, octaves[i+1])
		}

		sqlGameInsert.WriteString("(")
		sqlGameInsert.WriteString(strconv.Itoa(tournamentId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString("false")
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(octaves[i].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[0]))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(octaves[i+1].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[1]))
		sqlGameInsert.WriteString("),")

		fmt.Println(octaves[i].TeamId,": ", gameResult[0], " -",gameResult[1],":", octaves[i+1].TeamId)
	}

	fmt.Println("QUARTER FINALS MATCHES STARTING NOW!!!!!!")
	for i:=0; i< len(quarter); i=i+2 {
		sqlPlayoffInsert.WriteString("(")
		sqlPlayoffInsert.WriteString(strconv.Itoa(tournamentId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(quarter[i].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(quarter[i+1].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString("'Quarter'")
		sqlPlayoffInsert.WriteString("),")

		gameResult = eliminationGame()

		if gameResult[0] > gameResult[1] {
			semi = append(semi, quarter[i])
		} else {
			semi = append(semi, quarter[i+1])
		}

		sqlGameInsert.WriteString("(")
		sqlGameInsert.WriteString(strconv.Itoa(tournamentId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString("false")
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(quarter[i].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[0]))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(quarter[i+1].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[1]))
		sqlGameInsert.WriteString("),")
		fmt.Println(quarter[i].TeamId,": ", gameResult[0], " -",gameResult[1],":", quarter[i+1].TeamId)
	}

	fmt.Println("SEMI FINAL MATCHES STARTING NOW!!!!!!")
	for i:=0; i< len(semi); i=i+2 {
		sqlPlayoffInsert.WriteString("(")
		sqlPlayoffInsert.WriteString(strconv.Itoa(tournamentId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(semi[i].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString(strconv.Itoa(semi[i+1].TeamId))
		sqlPlayoffInsert.WriteString(",")
		sqlPlayoffInsert.WriteString("'Semi'")
		sqlPlayoffInsert.WriteString("),")

		gameResult = eliminationGame()

		if gameResult[0] > gameResult[1] {
			final = append(final, semi[i])
		} else {
			final = append(final, semi[i+1])
		}

		sqlGameInsert.WriteString("(")
		sqlGameInsert.WriteString(strconv.Itoa(tournamentId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString("false")
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(semi[i].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[0]))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(semi[i+1].TeamId))
		sqlGameInsert.WriteString(",")
		sqlGameInsert.WriteString(strconv.Itoa(gameResult[1]))
		sqlGameInsert.WriteString("),")

		fmt.Println(semi[i].TeamId,": ", gameResult[0], " -",gameResult[1],":", semi[i+1].TeamId)
	}

	fmt.Println("WHO WILL BE THE WINNER?!?!?!")
	fmt.Println("FINAL MATCH STARTING NOW!!!!!!")

	sqlPlayoffInsert.WriteString("(")
	sqlPlayoffInsert.WriteString(strconv.Itoa(tournamentId))
	sqlPlayoffInsert.WriteString(",")
	sqlPlayoffInsert.WriteString(strconv.Itoa(final[0].TeamId))
	sqlPlayoffInsert.WriteString(",")
	sqlPlayoffInsert.WriteString(strconv.Itoa(final[1].TeamId))
	sqlPlayoffInsert.WriteString(",")
	sqlPlayoffInsert.WriteString("'Final'")
	sqlPlayoffInsert.WriteString(")")

	gameResult = eliminationGame()

	sqlGameInsert.WriteString("(")
	sqlGameInsert.WriteString(strconv.Itoa(tournamentId))
	sqlGameInsert.WriteString(",")
	sqlGameInsert.WriteString("false")
	sqlGameInsert.WriteString(",")
	sqlGameInsert.WriteString(strconv.Itoa(final[0].TeamId))
	sqlGameInsert.WriteString(",")
	sqlGameInsert.WriteString(strconv.Itoa(gameResult[0]))
	sqlGameInsert.WriteString(",")
	sqlGameInsert.WriteString(strconv.Itoa(final[1].TeamId))
	sqlGameInsert.WriteString(",")
	sqlGameInsert.WriteString(strconv.Itoa(gameResult[1]))
	sqlGameInsert.WriteString(")")

	fmt.Println(final[0].TeamId,": ", gameResult[0], " -",gameResult[1],":", final[1].TeamId)

	fmt.Println(sqlPlayoffInsert.String())

	_, err = db.Query(sqlGameInsert.String())
	if err != nil {
		panic(err)
	}
	_, err = db.Query(sqlPlayoffInsert.String())
	if err != nil {
		panic(err)
	}

	if gameResult[0] > gameResult[1] {
		winner = final[0]
	} else {
		winner = final[1]
	}

	fmt.Println("AND THE WINNER OF THE TOURNAMENT ID:", tournamentId)
	fmt.Println("IIIISSS", GetTeam(winner.TeamId).Name, "!!!!!!!!!!!!")
	fmt.Println("CONGRATULATIONS!!!!!!!")
}

func GetPlayoffMatchesAndTeams(tournamentId int) []Model.PlayoffMatch {
	db, err := Database.Connect()
	defer db.Close()

	var games []Model.PlayoffMatch

	sqlStatement:=`select "Team1"."Name", "Game"."Team1_Score", "Team2"."Name", "Game"."Team2_Score", "Type" from  "Playoff" 
		join "Game" on "Playoff"."Team1_Id" = "Game"."Team1_Id" and "Playoff"."Team2_Id" = "Game"."Team2_Id"
		join "Team" as "Team1" on "Game"."Team1_Id" = "Team1"."Id" 
		join "Team" as "Team2" on "Game"."Team2_Id" = "Team2"."Id" 
		where "Is_Group" = false and "Game"."Tournament_Id" = ($1)`

	rows, err := db.Query(sqlStatement, tournamentId)
	for rows.Next() {
		var game Model.PlayoffMatch
		if err := rows.Scan(&game.Team1Name, &game.Team1Score, &game.Team2Name, &game.Team2Score, &game.Type); err != nil {
			log.Fatal(err)
		}
		games = append(games, game)
	}
	if err != nil {
		panic(err)
	}

	return games
}

func GetTournamentWinner(tournamentId int) string {
	db, err := Database.Connect()
	defer db.Close()

	var winner string

	sqlStatement:=`select "Team1"."Name", "Game"."Team1_Score", "Team2"."Name", "Game"."Team2_Score" from  "Playoff" 
		join "Game" on "Playoff"."Team1_Id" = "Game"."Team1_Id" and "Playoff"."Team2_Id" = "Game"."Team2_Id"
		join "Team" as "Team1" on "Game"."Team1_Id" = "Team1"."Id" 
		join "Team" as "Team2" on "Game"."Team2_Id" = "Team2"."Id" 
		where "Is_Group" = false and "Game"."Tournament_Id" = ($1) and "Type" = 'Final'`

	rows, err := db.Query(sqlStatement, tournamentId)
	for rows.Next() {
		var game Model.TeamAndMatch
		if err := rows.Scan(&game.Team1Name, &game.Team1Score, &game.Team2Name, &game.Team2Score); err != nil {
			log.Fatal(err)
		}
		if game.Team1Score > game.Team2Score {
			winner = game.Team1Name
		} else {
			winner = game.Team2Name
		}
	}
	if err != nil {
		panic(err)
	}

	return winner
}