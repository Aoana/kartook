package objects

import (
	"testing"
)

func isDifferent(a, b Object) bool {
	if a.X == b.X && a.Y == b.Y && a.VX == b.VX && a.VY == b.VY {
		return false
	}
	return true
}

func TestPosition(t *testing.T) {
	for _, c := range []struct {
		in, want Object
		dt       float64
	}{
		// Position update in all directions
		{Object{X: 0, Y: 0, VX: 0, VY: 0}, Object{X: 0, Y: 0, VX: 0, VY: 0}, 1},
		{Object{X: 0, Y: 0, VX: 1, VY: 0}, Object{X: 1, Y: 0, VX: 1, VY: 0}, 1},
		{Object{X: 0, Y: 0, VX: 0, VY: 1}, Object{X: 0, Y: 1, VX: 0, VY: 1}, 1},
		{Object{X: 0, Y: 0, VX: 1, VY: 1}, Object{X: 1, Y: 1, VX: 1, VY: 1}, 1},
		{Object{X: 0, Y: 0, VX: -1, VY: -1}, Object{X: -1, Y: -1, VX: -1, VY: -1}, 1},
		{Object{X: 0, Y: 0, VX: -1, VY: 1}, Object{X: -1, Y: 1, VX: -1, VY: 1}, 1},
		{Object{X: 0, Y: 0, VX: 1, VY: -1}, Object{X: 1, Y: -1, VX: 1, VY: -1}, 1},
		// Float values
		{Object{X: 1.1, Y: 2.2, VX: 1.1, VY: 2.2}, Object{X: 2.2, Y: 4.4, VX: 1.1, VY: 2.2}, 1},
		{Object{X: 1.1, Y: 2.2, VX: -1.1, VY: 2.2}, Object{X: 0, Y: 4.4, VX: -1.1, VY: 2.2}, 1},
		{Object{X: 1.1, Y: 2.2, VX: 1.1, VY: -2.2}, Object{X: 2.2, Y: 0, VX: 1.1, VY: -2.2}, 1},
		{Object{X: 1.1, Y: 2.2, VX: -1.1, VY: -2.2}, Object{X: 0, Y: 0, VX: -1.1, VY: -2.2}, 1},
		// dt variation
		{Object{X: 100, Y: 200, VX: 20, VY: 40}, Object{X: 102, Y: 204, VX: 20, VY: 40}, 10},
		{Object{X: 100, Y: 200, VX: 20, VY: 40}, Object{X: 108, Y: 216, VX: 20, VY: 40}, 2.5},
	} {
		c.in.Position(c.dt)
		if isDifferent(c.in, c.want) {
			t.Error("Position() failed", c.in, c.want)
		}
	}
}

func TestVelocity(t *testing.T) {

}

func TestNew(t *testing.T) {

}