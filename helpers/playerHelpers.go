package helpers

import (
	"fmt"
	"tripoley-server/models"
)

func ProcessAddPlayer(game *models.ServerData, player models.Player) {
	if GetPlayer(game, player.Name) == nil {
		game.Players = append(game.Players, models.Player{
			Name:  player.Name,
			Money: 150,
			Ready: false,
		})
		fmt.Println("Player added: ", player.Name)
	}
}

func GetPlayer(game *models.ServerData, name string) *models.Player {
	for i, player := range game.Players {
		if player.Name == name {
			return &game.Players[i]
		}
	}
	return nil
}
