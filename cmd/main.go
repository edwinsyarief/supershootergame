package main

import (
	"log"
	"supershootergame/internal/scenes"
	"supershootergame/pkg/engine"
)

func main() {
	var game = engine.NewGame(
		GameTitle,
		ScreenWidth, ScreenHeight, ScreenScale,
		IsExitOnEscapeKeyPress, Debug,
		ClearColor)
	game.SetScene(scenes.NewLoaderScene())

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
