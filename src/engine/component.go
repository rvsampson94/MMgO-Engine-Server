package engine

// Component a component can be attached to an entity to give that entity certain functionality
// every component implements an onUpdate method which is called by the engine every frame
type Component interface {
	OnUpdate(float64) error
}
