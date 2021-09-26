package IO

import (
	"JamEngine/main/Core"
	"JamEngine/main/Core/Components"
	"JamEngine/main/Core/Input"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type InputSystem struct{}

func (i InputSystem) Filter() Core.Filter {
	return Core.NewFilter([]string{Components.InputComponentID}, nil)
}

func (i InputSystem) Execute(scene Core.Scene, entities []Core.Entity) {
	event := sdl.PollEvent()

	var inputType Input.Type

	switch event.(type) {
	case *sdl.QuitEvent:
		inputType = Input.Quit
	default:
		return
	}

	for _, entity := range entities {
		inputComponentI, err := scene.GetComponentOfEntity(Components.InputComponentID, entity)
		if err != nil {
			fmt.Println(err)
			return
		}

		inputComponent := inputComponentI.(Components.InputComponent)
		inputComponent.Type = inputType

		err = scene.SetComponent(entity, Components.InputComponentID, inputComponent)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
