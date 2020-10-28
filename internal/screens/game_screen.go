package screens

import (
	"fmt"
	"os"

	"github.com/mradrianhh/Navigator/internal"
	"github.com/mradrianhh/Navigator/pkg/navigation"
	"github.com/mradrianhh/Navigator/pkg/navigation/models"
)

// GameScreen ...
type GameScreen struct {
	Identifier string
}

var gameScreen = GameScreen{Identifier: GAME_SCREEN}

// GetGameScreen ...
func GetGameScreen() *GameScreen {
	return &gameScreen
}

// Show ...
func (gameScreen GameScreen) Show(state *models.State) error {
	fmt.Print("Game\n\n1-Home\n0 - Exit\n\n")

	var response int
	_, err := fmt.Scan(&response)
	if err != nil {
		return err
	}

	switch response {
	case 1:
		state.State = internal.HOME
	case 0:
		navigation.Clear()
		os.Exit(0)
	default:
		state.State = internal.GAME
	}

	return nil
}
