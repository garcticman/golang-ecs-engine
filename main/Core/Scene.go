package Core

import (
	"fmt"
)

type Scene interface {
	Init()

	CreateEntity() Entity
	AddComponent(entity Entity, componentID string, componentValue interface{}) error
	SetComponent(entity Entity, componentID string, componentValue interface{}) error
	RemoveComponent(entity Entity, componentID string) error

	GetComponentOfEntity(componentID string, entity Entity) (interface{}, error)
	GetEntitiesWithFilter(filter Filter) []Entity

	GetInitializeSystems() []System
	AddInitializeSystem(system System)

	GetUpdateSystems() []System
	AddUpdateSystem(system System)

	GetReactiveSystems(componentID string) []ReactiveSystem
	AddReactiveSystem(system ReactiveSystem, componentID string)
}

type SceneBase struct {
	entities             []*Entity
	entitiesToComponents []map[string]interface{}
	lastEntityId         Entity

	initializeSystems []System
	systems           []System
	reactiveSystems   map[string][]ReactiveSystem
}

func (s *SceneBase) Init() {
	s.reactiveSystems = make(map[string][]ReactiveSystem)
}

func (s *SceneBase) CreateEntity() Entity {
	newEntity := s.lastEntityId
	s.entities = append(s.entities, &newEntity)
	s.entitiesToComponents = append(s.entitiesToComponents, map[string]interface{}{})

	s.lastEntityId++

	return newEntity
}

func (s *SceneBase) AddComponent(entity Entity, componentID string, componentValue interface{}) error {
	if s.lastEntityId < entity {
		return fmt.Errorf("scene does not contain entity with index: %d", entity)
	}

	s.entitiesToComponents[entity][componentID] = componentValue

	if systems, ok := s.reactiveSystems[componentID]; ok {
		for _, system := range systems {
			system.Execute(s, entity, WhenAdded)
		}
	}

	return nil
}

func (s *SceneBase) SetComponent(entity Entity, componentID string, componentValue interface{}) error {
	if s.lastEntityId < entity {
		return fmt.Errorf("scene does not contain entity with index: %d", entity)
	}

	s.entitiesToComponents[entity][componentID] = componentValue

	if systems, ok := s.reactiveSystems[componentID]; ok {
		for _, system := range systems {
			system.Execute(s, entity, WhenChanged)
		}
	}

	return nil
}

func (s *SceneBase) RemoveComponent(entity Entity, componentID string) error {
	if s.lastEntityId < entity {
		return fmt.Errorf("scene does not contain entity with index: %d", entity)
	}

	delete(s.entitiesToComponents[entity], componentID)

	if systems, ok := s.reactiveSystems[componentID]; ok {
		for _, system := range systems {
			system.Execute(s, entity, WhenRemoved)
		}
	}

	return nil
}

func (s SceneBase) GetComponentOfEntity(componentID string, entity Entity) (interface{}, error) {
	if s.lastEntityId < entity {
		return nil, fmt.Errorf("scene does not contain entity with index: %d", entity)
	}

	components := s.entitiesToComponents[entity]

	value, ok := components[componentID]
	if !ok {
		return nil, fmt.Errorf("component with id: %s not cotains in entity %d", componentID, entity)
	}

	return value, nil
}

func (s SceneBase) GetEntitiesWithFilter(filter Filter) (entities []Entity) {
	for entityIndex, components := range s.entitiesToComponents {

		if filter.Filter(components) {
			entities = append(entities, Entity(entityIndex))
		}
	}

	return
}

func (s SceneBase) GetInitializeSystems() []System {
	return s.initializeSystems
}

func (s *SceneBase) AddInitializeSystem(system System) {
	s.initializeSystems = append(s.initializeSystems, system)
}

func (s SceneBase) GetUpdateSystems() []System {
	return s.systems
}

func (s *SceneBase) AddUpdateSystem(system System) {
	s.systems = append(s.systems, system)
}

func (s *SceneBase) AddReactiveSystem(system ReactiveSystem, componentID string) {
	s.reactiveSystems[componentID] = append(s.reactiveSystems[componentID], system)
}

func (s SceneBase) GetReactiveSystems(componentID string) []ReactiveSystem {
	return s.reactiveSystems[componentID]
}
