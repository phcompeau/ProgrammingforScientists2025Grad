package main

import (
	"canvas"
	"image"
)



// let's place our drawing functions here.
func AnimateSystem(timePoints []Universe, canvasWidth, drawingFrequency int) []image.Image {
	images := make([]image.Image,0)

	trails := make(map[int][]OrderedPair)

	for i,u := range timePoints {
		if i%drawingFrequency == 0 {
			for bodyIndex, body := range u.bodies {
				trails[bodyIndex] = append(trails[bodyIndex], body.position)
				if len(trails[bodyIndex]) > numberOfTrailFrames*trailFrequency {
					trails[bodyIndex] = trails[bodyIndex][1:] // Limit trail length to MaxTrailLengthFactor
				}
			}
		}
		if i%drawingFrequency == 0 {
			images = append(images, DrawToCanvas(u, canvasWidth, trails, i))
		}
	}
	return images	

}



// DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
// object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(u Universe, canvasWidth int, trails map[int][]OrderedPair, frameCounter int) image.Image {
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// set canvas to white
	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)

	// Draw trails for all bodies
	//DrawTrails(&c, trails, frameCounter, u.width, float64(canvasWidth), u.bodies)

	// Draw the bodies themselves
	for _, b := range u.bodies {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		centerX := (b.position.x / u.width) * float64(canvasWidth)
		centerY := (b.position.y / u.width) * float64(canvasWidth)
		r := (b.radius / u.width) * float64(canvasWidth)

		if b.name == "Io" || b.name == "Ganymede" || b.name == "Callisto" || b.name == "Europa" {
			c.Circle(centerX, centerY, jupiterMoonMultiplier*r)
		} else {
			c.Circle(centerX, centerY, r)
		}
		c.Fill()
	}

	return c.GetImage()
}



const jupiterMoonMultiplier = 10.0
const trailFrequency = 10
const numberOfTrailFrames = 100
const trailThicknessFactor = 0.2
