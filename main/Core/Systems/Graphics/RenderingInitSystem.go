package Graphics

import (
	"JamEngine/main/Core"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderingInitSystem struct{}

func (r RenderingInitSystem) Filter() Core.Filter {
	return Core.NewFilter(nil, nil)
}

func (r RenderingInitSystem) Execute(scene Core.Scene, entities []Core.Entity) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println(err)
		return
	}

	if err := img.Init(img.INIT_PNG | img.INIT_JPG); err != nil {
		fmt.Println(err)
		return
	}
}
