package main

import (
	"math"
	"fyne.io/fyne"
	"time"
	"math/rand"
	"fmt"
)

// Function to determine how much lift moves before meeting desired floor
func distToFloor(pos, dest, floorHeight int) int {
	return (dest - pos)*floorHeight
}

// Function to change direction of lift movement
func directionOfLift(dist int) int {
	if dist == 0 {
		return 1
	} else {
		return dist / int(math.Abs(float64(dist)))
	}
}

// Function to simulate random choice by passenger that is not the current floor
func passengerFloorChoice(currentFloor int) int {
	var dest int
	for {
		dest = rand.Intn(8)
		if dest != currentFloor {
			return dest
		}
	}
}

// Function to move lift shape to desired floor
func goTo(pos, dest, floorHeight int, lift fyne.CanvasObject, container *fyne.Container) int {
	toMove := (dest - pos)*floorHeight
	startYPos := lift.Position().Y
	direction := directionOfLift(toMove)
	for {
		currentPosX, currentPosY := lift.Position().X, lift.Position().Y
		if currentPosY == (startYPos - toMove) {
			pos = dest
			// Sleep for a second to simulate stopping at floor
			time.Sleep(time.Second)
			break
		}
		lift.Move(fyne.Position{X: currentPosX, Y: currentPosY-direction})
		container.Refresh()
		time.Sleep(time.Millisecond * 50)
	}
	return pos
}

// Function to listen to a floor call from dispatcher and send lift
func controlLiftX(lift fyne.CanvasObject, call chan int,
	container *fyne.Container, liftNum int, floors *CurrentFloors) {
	currentFloor := 0
	for {
		select {
		case toFloor := <- call:
			floors.changeCurrentFloor(liftNum-1, -1)
			// Simulate travelling to call
			fmt.Println("Call at floor: ", toFloor)
			currentFloor = goTo(currentFloor, toFloor, 60, lift, container)

			// Simulate travelling to chosen floor
			toFloor = passengerFloorChoice(currentFloor)
			fmt.Println("Passenger chose: ", toFloor)
			currentFloor = goTo(currentFloor, toFloor, 60, lift, container)
			floors.changeCurrentFloor(liftNum-1, currentFloor)
			fmt.Println("Current floors are: ", floors.Current)
		}
	}
}