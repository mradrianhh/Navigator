package screens

import (
	"fmt"
	"os"

	"github.com/mradrianhh/Navigator/internal"
	"github.com/mradrianhh/Navigator/pkg/navigation"
	"github.com/mradrianhh/Navigator/pkg/navigation/models"
)

// OptionsScreen ...
type OptionsScreen struct {
	Identifier string
}

var optionsScreen = OptionsScreen{Identifier: OPTIONS_SCREEN}

// GetOptionsScreen ...
func GetOptionsScreen() *OptionsScreen {
	return &optionsScreen
}

// Show ...
func (optionsScreen OptionsScreen) Show(state *models.State) error {
	fmt.Print("Options\n\n1 - Home\n0 - Exit\n\n")

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
		state.State = internal.OPTIONS
	}

	return nil
}
