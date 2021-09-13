package dataset

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Game struct {
	Name  string  `csv:"Name"`
	Year  int     `csv:"Year"`
	Sales float64 `csv:"Global_Sales"`
}

func (g *Game) GetName() string {
	return g.Name
}

func (g *Game) GetYear() int {
	return g.Year
}

func (g *Game) GetSales() float64 {
	return g.Sales
}

func ExtractData() []Game {
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
