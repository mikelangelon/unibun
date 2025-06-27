package game

import (
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
)

type turnManager struct {
	currentTurn      int
	turnOrderDisplay []character
}

func (t turnManager) getPlayerType(playerType config.PlayerType) *entities.Player {
	for _, v := range t.turnOrderDisplay {
		switch item := v.(type) {
		case *entities.Player:
			if item.PlayerType == playerType {
				return item
			}
		}
	}
	return nil
}

func (t turnManager) getPlayerTypes(playerType config.PlayerType) []*entities.Player {
	var types []*entities.Player
	for _, v := range t.turnOrderDisplay {
		switch item := v.(type) {
		case *entities.Player:
			if item.PlayerType == playerType {
				types = append(types, item)
			}
		}
	}
	return types
}
