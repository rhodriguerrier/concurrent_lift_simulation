package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image/color"
	"time"
	"math/rand"
)

func main() {
	// Initialise app and container
	a := app.New()
	w := a.NewWindow("Lift System Concurrency")
	w.Resize(fyne.NewSize(300,500))
	container := fyne.NewContainer()

	// Paint floor lines and add to container
	var l *canvas.Line
	for i := 1; i <= 8; i++ {
		l = &canvas.Line{
			Position1: fyne.Position{X: 100, Y: 60*i}, Position2: fyne.Position{X: 200, Y: 60*i},
			StrokeColor: color.Color(color.RGBA{0, 0, 0, 255}), StrokeWidth: float32(0.5),
		}
		container.Objects = append(container.Objects, l)
	}

	// Paint the three lifts and add to container
	var r *canvas.Rectangle
	for i := 0; i < 3; i++ {
		r = &canvas.Rectangle{FillColor: color.Color(color.RGBA{0, 0, 0, 160})}
		r.Resize(fyne.Size{Width: 20, Height: 50})
		r.Move(fyne.Position{X: 100+(i*40), Y: 425})
		container.Objects = append(container.Objects, r)
	}

	// Add container with shapes to window
	w.SetContent(container)

	// Initialise current floor struct to be zero for all
	trackerObject := &CurrentFloors{Current: []int{0, 0, 0}}

	// Initialise channels for goroutines to communicate
	dispatch1 := make(chan int)
	dispatch2 := make(chan int)
	dispatch3 := make(chan int)
	move := make(chan int, 15)

	// Initialise three goroutines for each lift rectangle
	go controlLiftX(container.Objects[8], dispatch1, container, 1, trackerObject)
	go controlLiftX(container.Objects[9], dispatch2, container, 2, trackerObject)
	go controlLiftX(container.Objects[10], dispatch3, container, 3, trackerObject)

	// Gorountine to simulate a random floor call at random times between 1 and 10 seconds
	go func() {
		for {
			rand.Seed(time.Now().UTC().UnixNano())
			move <- rand.Intn(8)
			randomSleep := time.Duration(rand.Intn(10) + 1)
			time.Sleep(time.Second * randomSleep)
		}
	}()

	// Goroutine to relay the floor call to the next available lift
	go dispatchListener(move, dispatch1, dispatch2, dispatch3, trackerObject)

	w.ShowAndRun()
}