package IO

import (
	"JamEngine/main/Core"
	"JamEngine/main/Core/Components"
	"JamEngine/main/Core/Input"
	"fmt"
)

type QuitSystem struct{}

func (q QuitSystem) Execute(scene Core.Scene, entity Core.Entity, systemType Core.ReactType) {
	if systemType != Core.WhenChanged {
		return
	}

	inputComponentI, err := scene.GetComponentOfEntity(Components.InputComponentID, entity)
	if err != nil {
		fmt.Println(err)
		return
	}

	inputComponent := inputComponentI.(Components.InputComponent)
	if inputComponent.Type == Input.Quit {
		gameComponentI, err := scene.GetComponentOfEntity(Components.GameComponentID, entity)
		if err != nil {
			fmt.Println(err)
			return
		}

		gameComponent := gameComponentI.(*Components.GameComponent)
		gameComponent.Quit()
	}
}
