package Graphics

import (
	"JamEngine/main/Core"
	"JamEngine/main/Core/Components"
	"fmt"
	"github.com/veandco/go-sdl2/img"
)

type SpriteLoadSystem struct{}

func (s SpriteLoadSystem) Execute(scene Core.Scene, entity Core.Entity, reactType Core.ReactType) {
	if reactType != Core.WhenAdded {
		return
	}

	loadSpriteComponentI, err := scene.GetComponentOfEntity(Components.LoadSpriteComponentID, entity)
	if err != nil {
		fmt.Println(err)
		return
	}

	loadSpriteComponent := loadSpriteComponentI.(Components.LoadSpriteComponent)

	renderingFilter := Core.NewFilter([]string{Components.RenderingComponentID}, nil)
	renderingEntities := scene.GetEntitiesWithFilter(renderingFilter)

	for _, renderingEntity := range renderingEntities {
		renderingComponent, err := Components.GetRenderingComponent(scene, renderingEntity)
		if err != nil {
			fmt.Println(err)
			return
		}

		texture, err := img.LoadTexture(renderingComponent.Renderer, loadSpriteComponent.Path)
		if err != nil {
			fmt.Println(err)
			return
		}

		spriteComponent := Components.SpriteComponent{Texture: texture}

		if err = scene.AddComponent(entity, Components.SpriteComponentID, spriteComponent); err != nil {
			fmt.Println(err)
			return
		}

		if err = scene.RemoveComponent(entity, Components.LoadSpriteComponentID); err != nil {
			fmt.Println(err)
			return
		}
	}
}
