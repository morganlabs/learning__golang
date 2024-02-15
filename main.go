package main

import (
	"fmt"
)

func main() {
	age := 17
	rights := calculateRights(age)
	canDrive, canDrink, canDrinkUS := rights[0], rights[1], rights[2]

	fmt.Printf("This person is %v-years-old. This means that they can:\nDrive: %v,\nDrink: %v,\nDrink in the US: %v\n", age, canDrive, canDrink, canDrinkUS)

	age = 18
	rights = calculateRights(age)
	canDrive, canDrink, canDrinkUS = rights[0], rights[1], rights[2]

	fmt.Printf("This person is %v-years-old. This means that they can:\nDrive: %v,\nDrink: %v,\nDrink in the US: %v\n", age, canDrive, canDrink, canDrinkUS)

	age = 21
	rights = calculateRights(age)
	canDrive, canDrink, canDrinkUS = rights[0], rights[1], rights[2]

	fmt.Printf("This person is %v-years-old. This means that they can:\nDrive: %v,\nDrink: %v,\nDrink in the US: %v\n", age, canDrive, canDrink, canDrinkUS)
}

func calculateRights(age int) [3]bool {
	canDrive := age >= 17
	canDrink := age >= 18
	canDrinkUS := age >= 21

	return [3]bool{canDrive, canDrink, canDrinkUS}
}
