package engine

type collisionRect struct {
	parent        *Entity
	width, height float64
}

func newCollisionRect(parent *Entity, w float64, h float64) *collisionRect {
	return &collisionRect{
		parent: parent,
		width:  w,
		height: h,
	}
}

func (cb *collisionRect) OnUpdate(delta float64) error {
	return nil
}
