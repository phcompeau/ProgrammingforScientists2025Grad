package main

import (
	"canvas"
	"image"
)

// constants used to animate the Jupiter simulation
const jupiterMoonMultiplier = 10.0
const trailFrequency = 10
const numberOfTrailFrames = 100
const trailThicknessFactor = 0.2

// let's place our drawing functions here.

// Draw trails behind the bodies in our animation
// Inputs: canvas, map of trails for each body, frame counter, canvas and universe width, bodies
// Note that this uses a pointer to the canvas as input.
func DrawTrails(c *canvas.Canvas, trails map[int][]OrderedPair, frameCounter int, uWidth, canvasWidth float64, bodies []Body) {
	for bodyIndex, b := range bodies {
		trail := trails[bodyIndex]
		numTrails := len(trail)

		lineWidth := (b.radius/uWidth)*float64(canvasWidth)*trailThicknessFactor
		if b.name == "Io" || b.name == "Ganymede" || b.name == "Callisto" || b.name == "Europa" {
			lineWidth *= jupiterMoonMultiplier
		}
		c.SetLineWidth(lineWidth)

		for j := 0; j < numTrails-1; j++ {
			alpha := 255.0*float64(j)/float64(numTrails)
			red := uint8((1-alpha/255.0)*255.0 + (alpha/255.0)*float64(b.red))
			green := uint8((1-alpha/255.0)*255.0 + (alpha/255.0)*float64(b.green))
			blue := uint8((1-alpha/255.0)*255.0 + (alpha/255.0)*float64(b.blue))

			c.SetStrokeColor(canvas.MakeColor(red,green,blue))

			startX := (trail[j].x / uWidth) * canvasWidth
			startY := (trail[j].y / uWidth) * canvasWidth
			endX := (trail[j+1].x / uWidth) * canvasWidth
			endY := (trail[j+1].y / uWidth) * canvasWidth

			c.MoveTo(startX,startY)
			c.LineTo(endX,endY)
			c.Stroke()
		}

	}
}

// Animate the simulation frame by frame
// Input: slice of universe timepoints, width of canvas, frequency of frames drawn from timepoints 
// Output: slice of images
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
// Input: universe, width of canvas, map of position trails for each body, counter corresponding to simulation frames
// Output: image
func DrawToCanvas(u Universe, canvasWidth int, trails map[int][]OrderedPair, frameCounter int) image.Image {
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// set canvas to white
	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)

	// Draw trails for all bodies
	DrawTrails(&c, trails, frameCounter, u.width, float64(canvasWidth), u.bodies)

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


