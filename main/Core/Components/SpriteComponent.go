package Components

import "github.com/veandco/go-sdl2/sdl"

const SpriteComponentID = "SpriteComponent"

type SpriteComponent struct {
	Texture *sdl.Texture
}
