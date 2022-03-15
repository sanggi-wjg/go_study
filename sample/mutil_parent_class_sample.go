package main

import "fmt"

type Phone struct {
	ModelName string
}

func (p Phone) Call() {
	fmt.Println(p.ModelName + ": 부르르")
}

type Camera struct {
	ModelName string
}

func (c Camera) TakePicture() {
	fmt.Println(c.ModelName + ": 찰칵")
}

type SmartPhone struct {
	Phone
	Camera
}

func main() {
	smartPhone := SmartPhone{
		Phone{ModelName: "MyPhone"},
		Camera{ModelName: "MyCamera"},
	}
	smartPhone.Call()
	smartPhone.TakePicture()
}
