package solution

import (
	m "math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomType int

const SidesTriangle intCustomType = 3
const SidesSquare intCustomType = 4
const SidesCircle intCustomType = 0
const Pi = 3.14159265359

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var result float64

	switch sidesNum {
	case 3:
		result = m.Sqrt(3) / 4 * m.Pow(sideLen, 2)
	case 4:
		result = m.Pow(sideLen, 2)
	case 0:
		result = Pi * m.Pow(sideLen, 2)
	}

	return result
}
