package engine

import "github.com/hajimehoshi/ebiten/v2"

// Component : entity's component
type Component interface {
	IsActive() bool
	IsVisible() bool
	ComponentAdded(entity Entity)
	ComponentRemoved()
	ComponentSceneBegin(scene Scene)
	ComponentSceneEnd(scene Scene)
	ComponentEntityAwake()
	ComponentEntityAdded(scene Scene)
	ComponentEntityRemoved(scene Scene)
	ComponentUpdate()
	ComponentRender(screen *ebiten.Image)
}

// DefaultComponent : default component (e.g.: image, mover, etc.)
type DefaultComponent struct {
	Entity  Entity
	Active  bool
	Visible bool
}

func NewComponent() *DefaultComponent {
	var c = &DefaultComponent{
		Active:  true,
		Visible: true,
	}

	return c
}

func (c *DefaultComponent) ComponentAdded(entity Entity) {
	c.Entity = entity
}

func (c *DefaultComponent) ComponentRemoved() {

}

func (c *DefaultComponent) ComponentSceneBegin(scene Scene) {

}

func (c *DefaultComponent) ComponentSceneEnd(scene Scene) {

}

func (c *DefaultComponent) ComponentEntityAwake() {

}

func (c *DefaultComponent) ComponentEntityAdded(scene Scene) {

}

func (c *DefaultComponent) ComponentEntityRemoved(scene Scene) {

}

func (c *DefaultComponent) ComponentUpdate() {

}

func (c *DefaultComponent) ComponentRender(screen *ebiten.Image) {

}

func (c *DefaultComponent) IsActive() bool {
	return c.Active
}

func (c *DefaultComponent) IsVisible() bool {
	return c.Visible
}
