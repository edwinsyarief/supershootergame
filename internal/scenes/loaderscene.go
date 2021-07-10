package scenes

import (
	"image/color"
	"supershootergame/internal/components"
	"supershootergame/pkg/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

var (
	_fadeInEntity    *engine.DefaultEntity
	_fadeInComponent *components.ImageComponent

	_testEntity    *engine.DefaultEntity
	_testComponent *components.ImageComponent
)

// LoaderScene : Scene assets loader
type LoaderScene struct {
	*engine.DefaultScene
}

// NewLoaderScene : Create a new instance of Loader Scene
func NewLoaderScene() *LoaderScene {
	var s = LoaderScene{
		engine.NewDefaultScene(),
	}

	return &s
}

// Begin : Begin function
func (s *LoaderScene) Begin() {

	s.AddRenderer(engine.NewDefaultRenderer())

	_fadeInEntity = engine.NewEntity("fadeIn", f64.Vec2{0, 0})
	s.AddEntity(_fadeInEntity)

	w, h := ebiten.WindowSize()
	_fadeInComponent = components.NewImageComponent(w, h, color.Black)

	_fadeInEntity.EntityAdd(_fadeInComponent)

	_testEntity = engine.NewEntity("test", f64.Vec2{50, 50})
	s.AddEntity(_testEntity)

	_testComponent = components.NewImageComponent(320, 180, color.White)

	_testEntity.EntityAdd(_testComponent)
}

func (s *LoaderScene) Update() {
	s.DefaultScene.Update()

	// TODO: do something
}
