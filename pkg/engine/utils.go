package engine

func indexOfComponent(component *Component, componentList []Component) int {
	for i, v := range componentList {
		if component == &v {
			return i
		}
	}
	return -1
}

func indexOfRenderer(renderer Renderer, rendererList []Renderer) int {
	for i, v := range rendererList {
		if renderer == v {
			return i
		}
	}
	return -1
}

func removeFromComponentArray(component *Component, componentList []Component) []Component {
	var index = indexOfComponent(component, componentList)
	componentList[index] = componentList[len(componentList)-1]
	return componentList[:len(componentList)-1]
}

func removeFromRendererArray(renderer Renderer, rendererList []Renderer) []Renderer {
	var index = indexOfRenderer(renderer, rendererList)
	rendererList[index] = rendererList[len(rendererList)-1]
	return rendererList[:len(rendererList)-1]
}
