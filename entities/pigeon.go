package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"math/rand"
)

type Pigeon struct {
	*Enemy
}

func NewPigeon(startX, startY int) *Pigeon {
	return &Pigeon{
		Enemy: NewEnemy(startX, startY, []*ebiten.Image{common.GetImage(assets.Pidgeon)}),
	}
}

func nextRandomMove(level Level, enemy Enemier) (int, int) {
	possibleMoves := []struct{ dx, dy int }{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}
	rand.Shuffle(len(possibleMoves), func(i, j int) {
		possibleMoves[i], possibleMoves[j] = possibleMoves[j], possibleMoves[i]
	})
	for _, move := range possibleMoves {
		x, y := enemy.Position()
		targetX, targetY := x+move.dx, y+move.dy
		if !level.IsWalkable(targetX, targetY) {
			continue
		}
		return targetX, targetY
	}
	return enemy.Position()
}
