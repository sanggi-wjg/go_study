package main

import "fmt"

type Bicycle struct {
}

func (b Bicycle) Spare() {
	fmt.Println("Bicycle Spare")
}

func (b Bicycle) Run() {
	fmt.Println("Bicycle Run")
}

type MountainBicycle struct {
	Bicycle
}

func (m MountainBicycle) Spare() {
	fmt.Println("Mountain Bicycle Spare")
}

func (m MountainBicycle) Jump() {
	fmt.Println("Mountain Bicycle Jump")
}

func main() {
	bike := Bicycle{}
	bike.Spare()
	bike.Run()

	mBike := MountainBicycle{}
	mBike.Spare()
	mBike.Jump()
}
