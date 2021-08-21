package Core

type Filter struct {
	include []string
	exclude []string
}

func NewFilter(include []string, exclude []string) Filter {
	return Filter{include: include, exclude: exclude}
}

func (f Filter) Filter(components map[string]interface{}) bool {
	success := true

	for _, componentID := range f.include {
		if _, success = components[componentID]; success {
			break
		}
	}

	if !success {
		return false
	}

	for _, componentID := range f.exclude {
		if _, success = components[componentID]; success {
			return false
		}
	}

	return true
}
