package main

import (
	"fmt"
	"log"
	"soltani.ma/go/design-patterns/creational"
)

func main() {
	// Show how repeated calls share the same singleton instance.
	fmt.Println("Singleton")
	for i := 0; i < 30; i++ {
		go creational.GetInstance()
	}

	// Assemble vehicles one part at a time with different builders.
	fmt.Println("Builder")
	manufacturingComplex := creational.ManufacturingDirector{}
	carBuilder := &creational.CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	fmt.Println(car)
	bikeBuilder := &creational.BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	fmt.Println(motorbike)

	// Select a payment implementation from the requested method.
	fmt.Println("Factory")
	payment, err := creational.GetPaymentMethod(creational.Cash)
	if err != nil {
		log.Fatal("Error getting payment method Cash")
	}
	msg := payment.Pay(10.30)
	fmt.Println(msg)
	payment, err = creational.GetPaymentMethod(creational.CreditCard)
	msg = payment.Pay(22.30)
	fmt.Println(msg)

	// Clone shirt prototypes and confirm each copy can be changed independently.
	fmt.Println("Prototype")

	shirtCache := creational.GetShirtsCloner()
	if shirtCache == nil {
		log.Fatal("Received cache was nil")
	}
	item1, err := shirtCache.GetClone(creational.White)
	if err != nil {
		log.Fatal(err)
	}
	if item1 == whitePrototype {
		log.Fatal("item1 cannot be equal to the white prototype")
	}
	shirt1, ok := item1.(*creational.Shirt)
	if !ok {
		log.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"
	item2, err := shirtCache.GetClone(creational.White)
	if err != nil {
		log.Fatal(err)
	}
	shirt2, ok := item2.(*creational.Shirt)
	if !ok {
		log.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	if shirt1.SKU == shirt2.SKU {
		log.Fatal("SKU's of shirt1 and shirt2 must be different")
	}
	if shirt1 == shirt2 {
		log.Fatal("Shirt 1 cannot be equal to Shirt 2")
	}

}
