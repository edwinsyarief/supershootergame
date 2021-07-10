package components

import (
	"image/color"
	_ "image/png"
	"supershootergame/pkg/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageComponent struct {
	*engine.DefaultComponent
	Image  *ebiten.Image
	Width  int
	Height int
}

func NewImageComponent(width int, height int, color color.Color) *ImageComponent {
	timage := ebiten.NewImage(width, height)
	timage.Fill(color)
	var ic = &ImageComponent{
		DefaultComponent: engine.NewComponent(),
		Image:            timage,
		Width:            width,
		Height:           height,
	}

	return ic
}

func (c *ImageComponent) ComponentRender(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	pos := c.Entity.GetPosition()
	op.GeoM.Translate(pos[0], pos[1])

	screen.DrawImage(c.Image, op)
}
