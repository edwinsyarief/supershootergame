package engine

import "github.com/hajimehoshi/ebiten/v2"

// Renderer : Grapics renderer
type Renderer interface {
	RendererSceneBegin(scene Scene)
	RendererSceneEnd(scene Scene)
	RendererAdded(scene Scene)
	RendererRemoved(scene Scene)
	RendererUpdate(scene Scene)
	RendererBeforeRender(scene Scene, screen *ebiten.Image)
	RendererRender(scene Scene, screen *ebiten.Image)
	RendererAfterRender(scene Scene, screen *ebiten.Image)
}

// DefaultRenderer : Will render all entities
type DefaultRenderer struct {
	Visible bool
}

func NewDefaultRenderer() *DefaultRenderer {
	var r = &DefaultRenderer{
		Visible: true,
	}

	return r
}

func (r *DefaultRenderer) RendererSceneBegin(scene Scene) {

}

func (r *DefaultRenderer) RendererSceneEnd(scene Scene) {

}

func (r *DefaultRenderer) RendererAdded(scene Scene) {

}

func (r *DefaultRenderer) RendererRemoved(scene Scene) {

}

// RendererUpdate : Renderer update function
func (r *DefaultRenderer) RendererUpdate(scene Scene) {

}

// RendererBeforeRender : Renderer before render function
func (r *DefaultRenderer) RendererBeforeRender(scene Scene, screen *ebiten.Image) {

}

// RendererRender : Renderer render function
func (r *DefaultRenderer) RendererRender(scene Scene, screen *ebiten.Image) {
	scene.RenderEntities(screen)
}

// RendererAfterRender : Renderer after render function
func (r *DefaultRenderer) RendererAfterRender(scene Scene, screen *ebiten.Image) {

}
