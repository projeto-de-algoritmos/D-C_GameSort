package main

import "github.com/projeto-de-algoritmos/D-C_GameSort/src/model"

func BinarySearch(gameName string, gameList []model.Game) (bool, int, model.Game) {
	position := 0
	max := len(gameList)

	for position <= max {
		median := (position + max) / 2

		if gameList[median].GetName() == gameName {
			return true, median, gameList[median]
		} else if gameList[median].GetName() > gameName {
			max = median - 1
		} else {
			position = median + 1
		}
	}
	return false, -1, model.Game{"test", 1, 1}
}
