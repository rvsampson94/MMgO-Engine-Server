package engine

// InputEvent is a container for input information to be passed to the players RPC
type InputEvent struct {
	position  Vector
	direction Vector
}

// NewInputEvent creates a new input event
func NewInputEvent(position Vector, direction Vector) InputEvent {
	return InputEvent{
		position:  position,
		direction: direction,
	}
}

// RemotePlayerControler is a component which allows an entity to be controlled remotely using control packets
type RemotePlayerControler struct {
	parent    *Entity
	speed     float64
	Direction Vector
	events    []InputEvent //TODO make a queue implementation
}

// NewRemotePlayerController creates a new controller component
func NewRemotePlayerController(parent *Entity, speed float64) *RemotePlayerControler {
	return &RemotePlayerControler{
		parent:    parent,
		speed:     speed,
		Direction: NewVector(0, 0),
	}
}

// AddEvent appends a new input event to the events slice
func (rc *RemotePlayerControler) AddEvent(event InputEvent) {
	rc.events = append(rc.events, event)
}

// OnUpdate update function to be called each frame
func (rc *RemotePlayerControler) OnUpdate(delta float64) error {
	player := rc.parent
	for _, event := range rc.events {
		//TODO validate movement
		player.Position.X = event.position.X
		player.Position.Y = event.position.Y
		rc.Direction.X = event.direction.X
		rc.Direction.Y = event.direction.Y
		player.Update = true
	}
	rc.events = nil
	return nil
}
