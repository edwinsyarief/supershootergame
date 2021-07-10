package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

// Entity interface
type Entity interface {
	EntitySceneBegin(scene Scene)
	EntitySceneEnd(scene Scene)
	EntityAdd(component Component)
	EntityAdded(scene Scene)
	EntityRemove(component Component)
	EntityRemoved(scene Scene)
	EntityAwake(scene Scene)
	EntityUpdate()
	EntityRender(screen *ebiten.Image)
	EntityUpdateLists()
	GetComponentList() []Component
	GetName() string
	GetPosition() f64.Vec2
}

// DefaultEntity : game entity
type DefaultEntity struct {
	Name              string
	Tag               int64
	Scene             Scene
	Position          f64.Vec2
	Active            bool
	Visible           bool
	Collidable        bool
	ComponentList     []Component
	componentToAdd    []Component
	componentToRemove []Component
}

func NewEntity(name string, position f64.Vec2) *DefaultEntity {
	var e = &DefaultEntity{
		Tag:               0,
		Name:              name,
		Position:          position,
		Active:            true,
		Visible:           true,
		ComponentList:     []Component{},
		componentToAdd:    []Component{},
		componentToRemove: []Component{},
	}

	return e
}

func (e *DefaultEntity) GetComponentList() []Component {
	return e.ComponentList
}

func (e *DefaultEntity) GetName() string {
	return e.Name
}

func (e *DefaultEntity) GetPosition() f64.Vec2 {
	return e.Position
}

// EnditySceneBegin : Entity on scene begin
func (e *DefaultEntity) EntitySceneBegin(scene Scene) {
	for _, v := range e.ComponentList {
		v.ComponentSceneBegin(scene)
	}
}

// EntitySceneEnd : Entity on scene end
func (e *DefaultEntity) EntitySceneEnd(scene Scene) {
	for _, v := range e.ComponentList {
		v.ComponentSceneEnd(scene)
	}
}

// EntityAdd : Add component
func (e *DefaultEntity) EntityAdd(component Component) {
	if indexOfComponent(&component, e.ComponentList) == -1 && indexOfComponent(&component, e.componentToAdd) == -1 {
		e.componentToAdd = append(e.componentToAdd, component)
		component.ComponentAdded(e)
	}
}

// EntityAdded : on entity added
func (e *DefaultEntity) EntityAdded(scene Scene) {
	e.EntityUpdateLists()
	for _, v := range e.ComponentList {
		v.ComponentEntityAdded(scene)
	}
}

// EntityRemove : Remove component
func (e *DefaultEntity) EntityRemove(component Component) {
	if indexOfComponent(&component, e.ComponentList) > -1 && indexOfComponent(&component, e.componentToRemove) == -1 {
		e.componentToRemove = append(e.componentToRemove, component)
	}
}

// EntityRemoved : on entity removed
func (e *DefaultEntity) EntityRemoved(scene Scene) {
	for _, v := range e.ComponentList {
		v.ComponentEntityRemoved(scene)
	}
}

// EntityAwake : on entity awake
func (e *DefaultEntity) EntityAwake(scene Scene) {
	for _, v := range e.ComponentList {
		v.ComponentEntityAwake()
	}
}

// EntityUpdate : Update entity component
func (e *DefaultEntity) EntityUpdate() {
	for _, v := range e.ComponentList {
		if v.IsActive() {
			v.ComponentUpdate()
		}
	}
}

// EntityRender : Render entity component
func (e *DefaultEntity) EntityRender(screen *ebiten.Image) {
	for _, v := range e.ComponentList {
		if v.IsVisible() {
			v.ComponentRender(screen)
		}
	}
}

// EntityUpdateLists : Update component list
func (e *DefaultEntity) EntityUpdateLists() {
	if len(e.componentToAdd) > 0 {
		e.ComponentList = append(e.ComponentList, e.componentToAdd...)

		// clear
		e.componentToAdd = []Component{}
	}

	if len(e.componentToRemove) > 0 {

		for _, v := range e.componentToRemove {
			if indexOfComponent(&v, e.ComponentList) > -1 {
				e.ComponentList = removeFromComponentArray(&v, e.ComponentList)
			}
		}

		// clear
		e.componentToRemove = []Component{}
	}
}
