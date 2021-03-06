package bounce

import (
	"github.com/Aoana/go-ball-sim/assets/images"
	"github.com/Aoana/go-ball-sim/internal/pkg/ball"
	"github.com/Aoana/go-ball-sim/pkg/gfxutil"
	"github.com/Aoana/go-ball-sim/pkg/mathutil"
	"github.com/hajimehoshi/ebiten"
)

// Simulation variables
var (
	// Mathematical values
	dt, g float64 = 10.0, 9.80665
	// Images
	backgroundImage, leftWallImage, rightWallImage *ebiten.Image
	BallImage                                      *ebiten.Image
	// Screen Resolution
	ScreenWidth  = 1600
	ScreenHeight = 900
	// Starting values for balls
	X0            = float64(ScreenWidth) / 2
	Y0            = float64(ScreenHeight) / 10
	MinV0 float64 = -50
	MaxV0 float64 = 50
)

func init() {
	// Pre-Load images
	backgroundImage, _ = gfxutil.LoadPngSlice(images.ImageSky)
	leftWallImage, _ = gfxutil.LoadPngSlice(images.ImageWallLeft)
	rightWallImage, _ = gfxutil.LoadPngSlice(images.ImageWallRight)
	BallImage, _ = gfxutil.LoadPngSlice(images.ImageBasketBall)
}

// StartValues set starting position and velocity for a slice of balls
// Fixed starting position and velocity is random
func StartValues(nballs int) error {

	for i := 0; i < nballs; i++ {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(MinV0, MaxV0)
		vy0, _ := mathutil.RandInRange(MinV0, MaxV0)
		// Ball constructor
		b, err := ball.New(X0, Y0, vx0, vy0, 0.05, BallImage)
		if err != nil {
			return err
		}
		err = ball.Add(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// DrawScenery is a helper function to draw background and walls
func DrawScenery(screen *ebiten.Image) {
	// Draw background
	gfxutil.PrintImage(screen, backgroundImage, 0, 0, 3.0, 2.3)
	// Draw walls
	gfxutil.PrintImage(screen, leftWallImage, -50, 60, 1.9, 0.9)
	gfxutil.PrintImage(screen, rightWallImage, 1400, 50, 1.7, 0.9)
}

// Timestep is a helper function to perform a timestep with position and velocity updates
func Timestep(b *ball.Ball) {
	b.Obj.Position(dt)
	b.Obj.Velocity(0, g, dt)
}

// OutOfBound is a helper function to set the right boundary
// The values are simply set to fit the scenery
func OutOfBound(b *ball.Ball) {
	ball.Boundary(b, 100, float64(ScreenWidth)-60, -500, float64(ScreenHeight)-100, 0.8)
}
