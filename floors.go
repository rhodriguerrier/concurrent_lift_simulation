package main

// Struct to hold current floor for each lift
type CurrentFloors struct {
	Current []int
}

// Function to store current floor of each lift, -1 is placed when lift is in use
func (c *CurrentFloors) changeCurrentFloor(index, newVal int) {
	c.Current[index] = newVal
}