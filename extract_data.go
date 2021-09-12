package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

func extractData() []Game {
	gamesFile, err := os.OpenFile("games.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer gamesFile.Close()

	games := []*Game{}

	if err := gocsv.UnmarshalFile(gamesFile, &games); err != nil {
		panic(err)
	}

	gamesList := make([]Game, 0)

	for _, v := range games {
		gamesList = append(gamesList, *v)
	}

	return gamesList
}
