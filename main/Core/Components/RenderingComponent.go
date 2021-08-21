package Components

import "github.com/veandco/go-sdl2/sdl"

const RenderingComponentID = "RenderingComponent"

type RenderingComponent struct {
	Renderer *sdl.Renderer
	Window *sdl.Window
	Width int32
	Height int32
}