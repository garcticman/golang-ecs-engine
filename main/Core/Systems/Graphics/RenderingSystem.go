package Graphics

import (
	"JamEngine/main/Core"
	Components "JamEngine/main/Core/Components"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderingSystem struct{}

func (r RenderingSystem) Filter() Core.Filter {
	return Core.NewFilter([]string{Components.TransformComponentID, Components.SpriteComponentID}, nil)
}

func (r RenderingSystem) Execute(scene Core.Scene, entities []Core.Entity) {

	renderingFilter := Core.NewFilter([]string{Components.RenderingComponentID}, nil)
	renderingEntities := scene.GetEntitiesWithFilter(renderingFilter)

	for _, renderingEntity := range renderingEntities {
		renderingComponent, err := Components.GetRenderingComponent(scene, renderingEntity)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err = clearScreen(renderingComponent); err != nil {
			fmt.Println(err)
			return
		}

		if err = drawSprites(entities, renderingComponent, scene); err != nil {
			fmt.Println(err)
			return
		}

		renderingComponent.Renderer.Present()
		sdl.Delay(1000 / 60)
	}
}

func clearScreen(renderingComponent Components.RenderingComponent) error {

	err := renderingComponent.Renderer.Clear()
	if err != nil {
		return err
	}

	if err = renderingComponent.Renderer.SetDrawColor(0, 0, 0, 0x20); err != nil {
		return err
	}
	return nil
}

func drawSprites(entities []Core.Entity, renderingComponent Components.RenderingComponent, scene Core.Scene) error {
	for _, entity := range entities {
		spriteComponentI, err := scene.GetComponentOfEntity(Components.SpriteComponentID, entity)
		if err != nil {
			return err
		}

		spriteComponent := spriteComponentI.(Components.SpriteComponent)

		transformComponentI, err := scene.GetComponentOfEntity(Components.TransformComponentID, entity)
		if err != nil {
			return err
		}

		transformComponent := transformComponentI.(Components.TransformComponent)

		_, _, width, height, err := spriteComponent.Texture.Query()
		if err != nil {
			return err
		}

		rect := sdl.Rect{
			X: transformComponent.Position.X,
			Y: transformComponent.Position.Y,
			W: width,
			H: height,
		}

		err = renderingComponent.Renderer.Copy(spriteComponent.Texture, nil, &rect)
		if err != nil {
			return err
		}
	}

	return nil
}
