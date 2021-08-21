package Core

type System interface {
	Filter() Filter
	Execute(Scene, []Entity)
}
