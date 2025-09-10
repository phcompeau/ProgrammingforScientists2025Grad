package main

import (
	"math"
)

//let's place our gravity simulation functions here.
func SimulateGravity(initialUniverse Universe, numGens int, time float64) []Universe {
	timepoints := make([]Universe,numGens+1)
	timepoints[0] = initialUniverse
	for i := 1; i < numGens+1; i++ {
		timepoints[i] = UpdateUniverse(timepoints[i-1],time)
	}
	return timepoints
}

func UpdateUniverse(currentUniverse Universe, time float64) Universe {
	newUniverse := CopyUniverse(currentUniverse)
	for i,b := range newUniverse.bodies {
		oldAcceleration, oldVelocity := b.acceleration, b.velocity 

		newUniverse.bodies[i].acceleration = UpdateAcceleration(currentUniverse,b)

		newUniverse.bodies[i].velocity = UpdateVelocity(newUniverse.bodies[i],oldAcceleration,time)
		
		newUniverse.bodies[i].position = UpdatePosition(newUniverse.bodies[i],oldAcceleration,oldVelocity,time)
	}
	return newUniverse 
}

func UpdateAcceleration(currentUniverse Universe, b Body) OrderedPair {
	var accel OrderedPair 

	force := ComputeNetForce(currentUniverse, b)

	accel.x = force.x/b.mass
	accel.y = force.y/b.mass 

	return accel
}

func UpdateVelocity(b Body, oldAcceleration OrderedPair, time float64) OrderedPair {
	var currentVelocity OrderedPair 

	currentVelocity.x = b.velocity.x + 0.5*(b.acceleration.x + oldAcceleration.x)*time

	currentVelocity.y = b.velocity.y + 0.5*(b.acceleration.y + oldAcceleration.y)*time

	return currentVelocity
}

func UpdatePosition(b Body, oldAcceleration OrderedPair, oldVelocity OrderedPair, time float64) OrderedPair {
	var pos OrderedPair 

	pos.x = b.position.x + oldVelocity.x*time + 0.5*oldAcceleration.x*time*time 

	pos.y = b.position.y + oldVelocity.y*time + 0.5*oldAcceleration.y*time*time 

	return pos
}

//func ComputeNetForce()

//func ComputeForce()



//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
