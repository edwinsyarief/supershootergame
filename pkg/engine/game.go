package engine

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"runtime"
	"supershootergame/pkg/engine/customfonts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

var (
	normalFont font.Face
)

// Game struct
type Game struct {
	GameTitle              string
	Scene                  Scene
	NextScene              Scene
	ClearColor             color.Color
	ScreenWidth            int
	ScreenHeight           int
	ScreenScale            int
	IsDebug                bool
	IsExitOnEscapeKeyPress bool
}

// NewGame : create a new instance of Game
func NewGame(
	title string, screenWidth, screenHeight, screenScale int,
	isExitOnEscapeKeyPress, isDebug bool,
	clearColor color.Color) *Game {
	var g = &Game{
		GameTitle:              title,
		ScreenWidth:            screenWidth,
		ScreenHeight:           screenHeight,
		ScreenScale:            screenScale,
		IsDebug:                isDebug,
		IsExitOnEscapeKeyPress: isExitOnEscapeKeyPress,
		ClearColor:             clearColor,
	}

	g.Initialize()

	return g
}

func (g *Game) GetScene() Scene {
	return g.Scene
}

func (g *Game) SetScene(scene Scene) {
	g.NextScene = scene
}

// Update : main game loop
func (g *Game) Update() error {
	// key press check which only available on debug mode
	err := g.debugKeyPressCheck()
	if err != nil {
		return err
	}

	if g.Scene != nil {
		g.Scene.BeforeUpdate()
		g.Scene.Update()
		g.Scene.AfterUpdate()
	}

	// changing scene
	if g.Scene != g.NextScene {
		var lastScene = g.Scene
		if g.Scene != nil {
			g.Scene.End()
		}
		g.Scene = g.NextScene
		g.OnSceneTransition(lastScene, g.NextScene)
		if g.Scene != nil {
			g.Scene.Begin()
		}
	}

	return nil
}

func (g *Game) OnSceneTransition(lastScene Scene, nextScene Scene) {}

// Draw : main game render objects
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	if g.Scene != nil {
		g.Scene.BeforeRender(screen)
	}

	screen.Fill(g.ClearColor)

	if g.Scene != nil {
		g.Scene.Render(screen)
		g.Scene.AfterRender(screen)
	}

	// Show additional information on window's title
	g.generateDebugInfo(screen)
}

// Layout : game screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

// Initialize : game initialization
func (g *Game) Initialize() {
	ebiten.SetWindowSize(g.ScreenWidth*g.ScreenScale, g.ScreenHeight*g.ScreenScale)
	ebiten.SetFullscreen(!g.IsDebug)
	ebiten.SetWindowTitle(g.GameTitle)
	ebiten.SetWindowResizable(g.IsDebug)

	tt, err := opentype.Parse(customfonts.MinecraftiaRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.Hinting(font.HintingFull),
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Run : run game
func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

func (g *Game) debugKeyPressCheck() error {
	if g.IsExitOnEscapeKeyPress && inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Game ended")
	}

	if g.IsDebug && ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyR) {
		log.Print("Game restarted")
		g.Initialize()
	}

	return nil
}

func (g *Game) generateDebugInfo(screen *ebiten.Image) {
	if g.IsDebug {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		var title = fmt.Sprintf("%s - %s",
			g.GameTitle,
			fmt.Sprintf("FPS: %.0f, RAM: %d MB", ebiten.CurrentFPS(), bToMb(m.Alloc)))
		ebiten.SetWindowTitle(title)

		text.Draw(screen, fmt.Sprintf("FPS: %.0f", ebiten.CurrentFPS()), normalFont, 5, 30, color.Black)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
