package icstructembedding

import "fmt"

type Engine struct {
	HorsePower int
	Type       string
}

func (e Engine) Start() {
	fmt.Printf("Engine (%s, %d HP) started!\n", e.Type, e.HorsePower)
}

func (e Engine) Stop() {
	fmt.Println("Engine stopped.")
}

type Car struct {
	Brand string
	Model string
	Engine
}

func (c Car) Drive() {
	fmt.Printf("%s %s is driving\n", c.Brand, c.Model)
}

func (c Car) Info() {
	fmt.Printf("Car Info: %s %s | Engine: %s (%d HP)\n",
		c.Brand, c.Model, c.Engine.Type, c.Engine.HorsePower)
}

func Car_Embedding() {
	fmt.Println("\n===> Car Embedding")

	myCar := Car{
		Brand: "Toyota",
		Model: "Supra",
		Engine: Engine{
			HorsePower: 355,
			Type:       "Turbocharged I6",
		},
	}

	myCar.Start()
	myCar.Drive()
	myCar.Info()
	myCar.Stop()

}
