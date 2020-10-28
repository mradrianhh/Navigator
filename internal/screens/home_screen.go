package screens

import (
	"fmt"
	"os"

	"github.com/mradrianhh/Navigator/internal"
	"github.com/mradrianhh/Navigator/pkg/navigation"
	"github.com/mradrianhh/Navigator/pkg/navigation/models"
)

// HomeScreen ...
type HomeScreen struct {
	Identifier string
}

var homeScreen = HomeScreen{Identifier: HOME_SCREEN}

// GetHomeScreen ...
func GetHomeScreen() *HomeScreen {
	return &homeScreen
}

// Show ...
func (homeScreen HomeScreen) Show(state *models.State) error {
	fmt.Print("Home\n\n1 - Start Game\n2 - Options\n0 - Exit\n\n")

	var response int
	_, err := fmt.Scan(&response)
	if err != nil {
		return err
	}

	switch response {
	case 1:
		state.State = internal.GAME
	case 2:
		state.State = internal.OPTIONS
	case 0:
		navigation.Clear()
		os.Exit(0)
	default:
		state.State = internal.HOME
	}

	return nil
}
