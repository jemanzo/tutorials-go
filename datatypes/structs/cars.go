package main

import (
	"encoding/json"
	"fmt"
)

type (
	Car struct {
		Wheels uint `json:"wheels"`
		HP     uint `json:"hp"`
	}
	SuperCar struct {
		Top bool `json:"top"`
		*Car
		Nitro bool `json:"nitro"`
	}
)

// type AnyCar = Car || SuperCar

func (c *Car) ToJSON() string {
	c.Wheels = 2
	txt, _ := json.Marshal(c)
	return string(txt)
}

func CreateCars() {
	data := []byte(`{"wheels":4,"top":true,"hp":183,"nitro":false}`)
	var car2 SuperCar
	err := json.Unmarshal(data, &car2)
	if err != nil {
		fmt.Println(err)
	}
	// car1 := Car{
	// 	Wheels: 4,
	// 	HP:     233,
	// }
	// car2 := SuperCar{
	// 	&Car{
	// 		Wheels: 4,
	// 		HP:     233,
	// 	},

	// 	true,
	// }
	// fmt.Println(car1.ToJSON())
	fmt.Println(car2.Wheels)
	fmt.Println(car2.ToJSON())
	fmt.Println(car2.Wheels)
}
