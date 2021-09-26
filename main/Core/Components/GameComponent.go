package Components

import "JamEngine/main/Core"

const GameComponentID = "GameComponent"

type GameComponent struct {
	scenes       []Core.Scene
	currentScene uint64
	quit         bool
}

func (game *GameComponent) Start() {
	scene := game.scenes[game.currentScene]
	for !game.quit {
		for _, system := range scene.GetInitializeSystems() {
			componentIDs := system.Filter()

			entities := scene.GetEntitiesWithFilter(componentIDs)

			system.Execute(scene, entities)
		}
	}
}

func (game *GameComponent) Update() {
	scene := game.scenes[game.currentScene]
	for !game.quit {
		for _, system := range scene.GetUpdateSystems() {
			componentIDs := system.Filter()

			entities := scene.GetEntitiesWithFilter(componentIDs)

			system.Execute(scene, entities)
		}
	}
}

func (game *GameComponent) AddScene(scene Core.Scene) {
	game.scenes = append(game.scenes, scene)
}

func (game *GameComponent) Quit() {
	game.quit = true
}
