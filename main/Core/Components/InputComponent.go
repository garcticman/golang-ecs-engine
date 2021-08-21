package Components

import "JamEngine/main/Core/Input"

const InputComponentID = "InputComponent"

type InputComponent struct {
	Input.Type
	keys []Input.Key
}