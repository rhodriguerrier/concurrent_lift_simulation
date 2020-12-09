package main

import (
	"math"
)

// Return nearest lift number or return -1 if no lifts are available
func chooseLift(currentFloors *CurrentFloors, dest int) int {
	availableLifts := 0
	chosenLift := -1
	var destDiff float64
	for i := 0; i < len(currentFloors.Current); i++ {
		if currentFloors.Current[i] == -1 {
			continue
		} else if availableLifts < 1 {
			availableLifts += 1
			destDiff = math.Abs(float64(dest - currentFloors.Current[i]))
			chosenLift = i+1
		} else if math.Abs(float64(dest - currentFloors.Current[i])) < destDiff {
			chosenLift = i+1
		}
	}
	return chosenLift
}

// Function to check if any lifts are available so that floor call is not lost in select
func anyAvailable(currentFloors *CurrentFloors) bool {
	for i := 0; i < len(currentFloors.Current); i++ {
		if currentFloors.Current[i] != -1 {
			return true
		}
	}
	return false
}

// Function listens for a floor call and then sends it on to the nearest, available lift
func dispatchListener(floorCall, dispatch1, dispatch2, dispatch3 chan int, floors *CurrentFloors) {
	for {
		// Skip select if no lifts are available so that floorCall value is not lost in select
		if !(anyAvailable(floors)) {
			continue
		} 
		// If at least one lift is available then check for a floor call and then call nearest lift
		select {
		case toFloor := <- floorCall:
			nearestLift := chooseLift(floors, toFloor)
			if nearestLift == 1 {
				dispatch1 <- toFloor
			} else if nearestLift == 2 {
				dispatch2 <- toFloor
			} else if nearestLift == 3 {
				dispatch3 <- toFloor
			}
		}
	}
}