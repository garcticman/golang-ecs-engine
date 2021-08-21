package Systems

import (
	"JamEngine/main/Core"
	"JamEngine/main/Core/Components"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderingSystem struct{}

func (r RenderingSystem) Filter() Core.Filter {
	return Core.NewFilter([]string{Components.RenderingComponentID}, nil)
}

func (r RenderingSystem) Execute(scene Core.Scene, entities []Core.Entity) {
	for _, entity := range entities {
		renderingComponentI, err := scene.GetComponentOfEntity(Components.RenderingComponentID, entity)
		if err != nil {
			fmt.Println(err)
			return
		}

		renderingComponent := renderingComponentI.(Components.RenderingComponent)

		err = renderingComponent.Renderer.Clear()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = renderingComponent.Renderer.SetDrawColor(0, 0, 0, 0x20)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = renderingComponent.Renderer.FillRect(&sdl.Rect{W: renderingComponent.Width, H: renderingComponent.Height})
		if err != nil {
			fmt.Println(err)
			return
		}

		// TODO render scene entities

		renderingComponent.Renderer.Present()
		sdl.Delay(1000 / 60)
	}
}
