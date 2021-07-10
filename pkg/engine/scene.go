package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Begin()
	End()
	BeforeUpdate()
	Update()
	AfterUpdate()
	BeforeRender(screen *ebiten.Image)
	Render(screen *ebiten.Image)
	AfterRender(screen *ebiten.Image)
	RenderEntities(screen *ebiten.Image)
	OnEndOfFrame()
	GetEntityList() map[string]Entity
	GetRendererList() []Renderer
}

type DefaultScene struct {
	EntityList       map[string]Entity
	entityToAdd      map[string]Entity
	entityToRemove   map[string]Entity
	entityToAwake    map[string]Entity
	RendererList     []Renderer
	rendererToAdd    []Renderer
	rendererToRemove []Renderer
}

// NewDefaultScene : Create a new instance of DefaultScene
func NewDefaultScene() *DefaultScene {
	var s = &DefaultScene{
		EntityList:       make(map[string]Entity),
		entityToAdd:      make(map[string]Entity),
		entityToRemove:   make(map[string]Entity),
		entityToAwake:    make(map[string]Entity),
		RendererList:     []Renderer{},
		rendererToAdd:    []Renderer{},
		rendererToRemove: []Renderer{},
	}

	return s
}

// GetEntityList : Get entity list
func (s *DefaultScene) GetEntityList() map[string]Entity {
	return s.EntityList
}

func (s *DefaultScene) GetRendererList() []Renderer {
	return s.RendererList
}

// Begin : Begin function
func (s *DefaultScene) Begin() {
	// Override to do something
}

// End : End function
func (s *DefaultScene) End() {

}

// BeforeUpdate : Before update function
func (s *DefaultScene) BeforeUpdate() {
	s.UpdateEntityLists()
	s.UpdateRendererLists()
}

// Update : Update scene
func (s *DefaultScene) Update() {
	//fmt.Println("Default scene update")
	for _, v := range s.EntityList {
		v.EntityUpdate()
	}

	for _, v := range s.RendererList {
		v.RendererUpdate(s)
	}
}

// AfterUpdate : After update function
func (s *DefaultScene) AfterUpdate() {
	s.OnEndOfFrame()
}

// OnEndOfFrame : Execute end of frame event
func (s *DefaultScene) OnEndOfFrame() {
	// TODO : do something
}

// BeforeRender : Before render function
func (s *DefaultScene) BeforeRender(screen *ebiten.Image) {
	for _, v := range s.RendererList {
		v.RendererBeforeRender(s, screen)
	}
}

// Render : Render
func (s *DefaultScene) Render(screen *ebiten.Image) {
	for _, v := range s.RendererList {
		v.RendererRender(s, screen)
	}
}

// RenderEntities : RenderAllEntities
func (s *DefaultScene) RenderEntities(screen *ebiten.Image) {
	for _, v := range s.EntityList {
		v.EntityRender(screen)
	}
}

// AfterRender : After render function
func (s *DefaultScene) AfterRender(screen *ebiten.Image) {
	for _, v := range s.RendererList {
		v.RendererAfterRender(s, screen)
	}
}

// UpdateEntityLists : Update entity list
func (s *DefaultScene) UpdateEntityLists() {
	if len(s.entityToAdd) > 0 {
		for e, v := range s.entityToAdd {
			s.EntityList[e] = v
			v.EntityAdded(s)
		}
	}

	if len(s.entityToRemove) > 0 {
		for e := range s.entityToRemove {
			delete(s.EntityList, e)
		}

		// clear
		s.entityToRemove = make(map[string]Entity)
	}

	if len(s.entityToAdd) > 0 {
		for e, v := range s.entityToAdd {
			s.entityToAwake[e] = v
			v.EntityAwake(s)
		}

		s.entityToAdd = map[string]Entity{}
		s.entityToAwake = map[string]Entity{}
	}
}

// AddEntity : Add Entity
func (s *DefaultScene) AddEntity(entity Entity) {
	_, addOk := s.entityToAdd[entity.GetName()]
	_, currentOk := s.EntityList[entity.GetName()]

	if !addOk && !currentOk {
		s.entityToAdd[entity.GetName()] = entity
	}
}

// RemoveEntity : Remove Entity
func (s *DefaultScene) RemoveEntity(entity Entity) {
	_, removeOk := s.entityToRemove[entity.GetName()]
	_, currentOk := s.EntityList[entity.GetName()]

	if !removeOk && currentOk {
		s.entityToRemove[entity.GetName()] = entity
	}
}

/* func (s *DefaultScene) entitiesCount() int {
	return len(s.EntityList)
} */

// UpdateRendererLists : Update renderer list
func (s *DefaultScene) UpdateRendererLists() {
	if len(s.rendererToAdd) > 0 {
		s.RendererList = append(s.RendererList, s.rendererToAdd...)

		// clear
		s.rendererToAdd = []Renderer{}
	}

	if len(s.rendererToRemove) > 0 {
		for _, v := range s.rendererToRemove {
			if indexOfRenderer(v, s.RendererList) > -1 {
				s.RendererList = removeFromRendererArray(v, s.RendererList)
			}
		}

		// clear
		s.rendererToRemove = []Renderer{}
	}
}

// AddRenderer : Add Renderer
func (s *DefaultScene) AddRenderer(renderer Renderer) {
	if indexOfRenderer(renderer, s.RendererList) == -1 && indexOfRenderer(renderer, s.rendererToAdd) == -1 {
		s.rendererToAdd = append(s.rendererToAdd, renderer)
	}
}

// RemoveRenderer : Remove Renderer
func (s *DefaultScene) RemoveRenderer(renderer Renderer) {
	if indexOfRenderer(renderer, s.RendererList) > -1 && indexOfRenderer(renderer, s.rendererToRemove) == -1 {
		s.rendererToRemove = append(s.rendererToRemove, renderer)
	}
}
