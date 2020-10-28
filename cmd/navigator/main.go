package main

import (
	"fmt"
	"os"

	"github.com/mradrianhh/Navigator/internal"
	"github.com/mradrianhh/Navigator/internal/screens"
	"github.com/mradrianhh/Navigator/pkg/navigation"
	"github.com/mradrianhh/Navigator/pkg/navigation/models"
)

var state models.State
var navigator navigation.Navigator

func main() {
	setup()
	loop()
}

func setup() {
	navigator = navigation.NewNavigator()
	state.State = internal.HOME

	navigator.AddScreen(screens.GetHomeScreen().Identifier, screens.GetHomeScreen())
	navigator.AddScreen(screens.GetGameScreen().Identifier, screens.GetGameScreen())
	navigator.AddScreen(screens.GetOptionsScreen().Identifier, screens.GetOptionsScreen())
}

func loop() {
	for {
		switch state.State {
		case internal.HOME:
			err := navigator.RemoveScreensAndPush(&state, screens.HOME_SCREEN)
			checkError(err)
		case internal.GAME:
			err := navigator.RemoveScreensAndPush(&state, screens.GAME_SCREEN)
			checkError(err)
		case internal.OPTIONS:
			err := navigator.RemoveScreensAndPush(&state, screens.OPTIONS_SCREEN)
			checkError(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
