package engine

import (
	"fmt"
	"reflect"
)

// Entity is a container for all game objects
type Entity struct {
	Position   Vector
	Components []Component
	Update     bool
}

// NewEntity creates a new entity at position (x, y)
func NewEntity(x float64, y float64) *Entity {
	return &Entity{
		Position: NewVector(x, y),
		Update:   true,
	}
}

// AddComponent adds a new component to the entity
func (e *Entity) AddComponent(new Component) {
	for _, existing := range e.Components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("Component of type %v already exists on entity", reflect.TypeOf(new)))
		}
	}
	e.Components = append(e.Components, new)
}

// GetComponent if a component of the given type is found on the entity this method will return that component
// otherwise the function panics
func (e *Entity) GetComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range e.Components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	panic(fmt.Sprintf("No component with type %v on entity", typ))
}
