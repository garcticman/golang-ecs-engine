package Core

type ReactType byte

const (
	WhenAdded = ReactType(0)
	WhenChanged = ReactType(1)
	WhenRemoved = ReactType(2)
)

type ReactiveSystem interface {
	Execute(Scene, Entity, ReactType)
}