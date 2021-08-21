package main

import (
	"JamEngine/main/Core"
	"JamEngine/main/Core/Components"
	"JamEngine/main/Core/Systems"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	scene := &Core.SceneBase{}
	scene.Init()

	gameEntity, game := createGameEntity(scene)

	createRendering(scene, gameEntity)
	createInput(scene, gameEntity)

	game.AddScene(scene)
	game.Update()
}

func createGameEntity(scene Core.Scene) (Core.Entity, *Components.GameComponent)  {
	game := Components.GameComponent{}

	gameEntity := scene.CreateEntity()

	if err := scene.AddComponent(gameEntity, Components.GameComponentID, &game); err != nil {
		fmt.Println(err)
	}

	return gameEntity, &game
}

func createRendering(scene Core.Scene, entity Core.Entity) {
	window, _ := sdl.CreateWindow("title", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)

	rendering := Components.RenderingComponent{
		Window: window,
		Renderer: renderer,
	}

	if err := scene.AddComponent(entity, Components.RenderingComponentID, rendering); err != nil {
		fmt.Println(err)
	}
	scene.AddSystem(Systems.RenderingSystem{})
}

func createInput(scene Core.Scene, entity Core.Entity) {
	if err := scene.AddComponent(entity, Components.InputComponentID, Components.InputComponent{}); err != nil {
		fmt.Println(err)
	}

	scene.AddSystem(Systems.InputSystem{})
	scene.AddReactiveSystem(Systems.QuitSystem{}, Components.InputComponentID)
}