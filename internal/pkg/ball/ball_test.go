package ball

import (
	"github.com/Aoana/ball-sim-go/assets/images"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"testing"
)

func isDifferent(x, y, vx, vy float64, b *Ball) bool {

	if x == b.Obj.X[0] && y == b.Obj.X[1] && vx == b.Obj.V[0] && vy == b.Obj.V[1] {
		return false
	}
	return true
}

func isPresent(x, y, vx, vy, scale float64) bool {
	for i := range BallList {
		b := BallList[i]
		if x == b.Obj.X[0] && y == b.Obj.X[1] && vx == b.Obj.V[0] && vy == b.Obj.V[1] && scale == b.Scale {
			return true
		}
	}
	return false
}

func TestList(t *testing.T) {

	img, _ := gfxutil.LoadPngSlice(images.ImageStar)

	for _, c := range []struct {
		x, y, vx, vy, scale float64
	}{
		{1.1, 2.2, 3.3, 4.4, 0.5},
		{-1.1, 50, 50000, 0, 2.323},
		{0, 0, 0, 0, -99.9},
	} {
		b, err := New(c.x, c.y, c.vx, c.vy, c.scale, img)
		if err != nil {
			t.Errorf("New(%f, %f, %f, %f, %f, %p) error %s", c.x, c.y, c.vx, c.vy, c.scale, img, err)
		}
		err = Add(b)
		if err != nil {
			t.Errorf("Add(%+v) error %s", b, err)
		}
		if !isPresent(c.x, c.y, c.vx, c.vy, c.scale) {
			t.Errorf("Ball missing (%f, %f, %f, %f, %f, %p)", c.x, c.y, c.vx, c.vy, c.scale, img)
		}
	}
	err := Remove(1)
	if err != nil {
		t.Error("Remove(1) failed ", err)
	}
	x, y, vx, vy, scale := -1.1, 50.0, 50000.0, 0.0, 2.323
	if isPresent(x, y, vx, vy, scale) {
		t.Errorf("Ball present (%f, %f, %f, %f, %f, %p)", x, y, vx, vy, scale, img)
	}
}

func TestBoundary(t *testing.T) {

	img, _ := gfxutil.LoadPngSlice(images.ImageStar)

	for _, c := range []struct {
		x0, y0, vx0, vy0 float64
		x, y, vx, vy     float64
	}{
		// Not supposed to bounce
		{1, 1, 1, 1, 1, 1, 1, 1},
		{99, 99, -1, -1, 99, 99, -1, -1},
		// Corner bounce
		{1, 1, -1, -1, 1, 1, 1, 1},
		{99, 99, 1, 1, 99, 99, -1, -1},
	} {
		b, err := New(c.x0, c.y0, c.vx0, c.vy0, 1, img)
		if err != nil {
			t.Errorf("New(%f, %f, %f, %f, %f, %p) error %s", c.x0, c.y0, c.vx0, c.vy0, 1.0, img, err)
		}
		err = Boundary(b, 0.0, 100.0, 0.0, 100.0, 1.0)
		if err != nil {
			t.Errorf("Boundary(%+v, 0, 1, 0, 1) error %s", b.Obj, err)
		}
		if isDifferent(c.x, c.y, c.vx, c.vy, b) {
			t.Errorf("Ball with (%f, %f, %f, %f)) = %+v != (%f, %f, %f, %f)", c.x0, c.y0, c.vx0, c.vy0, b.Obj, c.x, c.y, c.vx, c.vy)
		}
	}
}
