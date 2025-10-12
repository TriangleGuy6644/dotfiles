package main

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake   string
	carModel  string
	modelYear int32
	engine    T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:   "Honda",
		carModel:  "Civic",
		modelYear: 2007,
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}
	var electricCar = car[electricEngine]{
		carMake:   "Tesla",
		carModel:  "Model Y",
		modelYear: 2024,
		engine: electricEngine{
			kwh:   57.7,
			mpkwh: 4.17,
		},
	}

}
