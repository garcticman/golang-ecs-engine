package Components

import (
	"JamEngine/main/Core"
	"github.com/veandco/go-sdl2/sdl"
)

const RenderingComponentID = "RenderingComponent"

type RenderingComponent struct {
	Renderer *sdl.Renderer
	Window   *sdl.Window
	Width    int32
	Height   int32
}

func GetRenderingComponent(scene Core.Scene, renderingEntity Core.Entity) (RenderingComponent, error) {

	renderingComponentI, err := scene.GetComponentOfEntity(RenderingComponentID, renderingEntity)
	if err != nil {
		return RenderingComponent{}, err
	}

	renderingComponent := renderingComponentI.(RenderingComponent)
	return renderingComponent, nil
}
