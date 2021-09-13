package dataset

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/projeto-de-algoritmos/D-C_GameSort/src/model"
)

func ExtractData() []model.Game {
	gamesFile, err := os.OpenFile("games.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer gamesFile.Close()

	games := []*model.Game{}

	if err := gocsv.UnmarshalFile(gamesFile, &games); err != nil {
		panic(err)
	}

	gamesList := make([]model.Game, 0)

	for _, v := range games {
		gamesList = append(gamesList, *v)
	}

	return gamesList
}
