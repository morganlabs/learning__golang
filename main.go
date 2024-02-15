package main

import "fmt"

func main() {
	var ageOne int = 16 // AT MINIMUM 32-bits, but this is NOT the same as int32.
	var ageTwo = 17
	ageThree := 18

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits
	var numOne int8 = 25    // 8-bits
	var numTwo int16 = 25   // 8-bits
	var numThree int32 = 25 // 8-bits
	var numFour int64 = 25  // 8-bits
	var numFive int8 = 25   // 8-bits
	var numSix int8 = 25    // 8-bits
	fmt.Println(numOne, numTwo, numThree, numFour, numFive, numSix)

	// Unsigned ints
	// These can only be positive numbers
	// This is useful for when you need to use lower bits (say, 8-bits), but need a larger number.
	// Because the negatives aren't acounted for, int8 can go down to -128 and up to *127*,
	// but uint8 can go down to 0 and up to *255*.

	// Floats
	// These are limited to float32 and float64, so 32- and 64-bits relatively
	// These allow you to have a decimal place in a number
	// Floats are SIGNED.
	// f64 has a higher precision than f32, and for this reason, when Go infers the type of a float, f64 is the default.
	var floatOne float32 = 293018348433948938938983478945782748372.12
	var floatTwo float64 = 45738943728947274382973478934728947289473289428947238947238953429812357892238943278957348972423758923475892347892334287432789473289472389574289427890543728957377.21 // this isnt even the limit of a 64-bit float...
	fmt.Println(floatOne, floatTwo)
}
