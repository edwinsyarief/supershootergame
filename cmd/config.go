package main

import "image/color"

const (
	// GameTitle : title of game
	GameTitle = "Shadow Storm"
	// ScreenWidth : game window width
	ScreenWidth = 1280
	// ScreenHeight : game window height
	ScreenHeight = 720
	// ScreenScale : game window scale
	ScreenScale = 1
	// Debug : set this to false when production
	Debug = true
	// IsExitOnEscapeKeyPress : set this to true will exit when user press Escape Key
	IsExitOnEscapeKeyPress = true
)

var (
	// ClearColor : background color
	ClearColor color.Color = color.NRGBA{0xff, 0xff, 0xff, 0xff}
)
