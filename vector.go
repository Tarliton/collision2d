package collision2d

import (
	"fmt"
	"math"
)

//Vector is a simple 2D vector/point struct.
type Vector struct {
	X, Y float64
}

func (vector Vector) String() string {
	return fmt.Sprintf("{X:%f, Y:%f}\n", vector.X, vector.Y)
}

//NewVector create a new vector with the values of x and y
func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

//Copy the value of other vector to a new one.
func (vector Vector) Copy(other Vector) Vector {
	vector.X = other.X
	vector.Y = other.Y
	return vector
}

//Clone this vector coordinates to a new vector with the same coordinates as this one.
func (vector Vector) Clone() Vector {
	another := Vector{}.Copy(vector)
	return another
}

//Perp returns a new vector perpendicular from this one.
func (vector Vector) Perp() Vector {
	x := vector.X
	vector.X = vector.Y
	vector.Y = -x
	return vector
}

//Rotate returns a new vector rotated counter-clockwise by the specified number of radians.
func (vector Vector) Rotate(angle float64) Vector {
	x := vector.X
	y := vector.Y
	vector.X = x*math.Cos(angle) - y*math.Sin(angle)
	vector.Y = x*math.Sin(angle) + y*math.Cos(angle)
	return vector
}

//Reverse returns a new vector that is reversed from this one.
func (vector Vector) Reverse() Vector {
	vector.X = -vector.X
	vector.Y = -vector.Y
	return vector
}

//Normalize returns a new unit-length vector.
func (vector Vector) Normalize() Vector {
	d := vector.Len()
	if d > 0 {
		vector.X /= d
		vector.Y /= d
	}
	return vector
}

//Add returns a new vector with the result of adding another vector to this one.
func (vector Vector) Add(other Vector) Vector {
	vector.X += other.X
	vector.Y += other.Y
	return vector
}

//Sub returns a new vector that is the result of subtracting another vector from this one.
func (vector Vector) Sub(other Vector) Vector {
	vector.X -= other.X
	vector.Y -= other.Y
	return vector
}

//Scale returns a new vector scaled in the direction of X and Y by value x
func (vector Vector) Scale(x float64) Vector {
	vector.X *= x
	vector.Y *= x
	return vector
}

//ScaleDifferent returns a new vector scaled in the direction of X and Y by value x and y respectively
func (vector Vector) ScaleDifferent(x, y float64) Vector {
	vector.X *= x
	vector.Y *= y
	return vector
}

//Project this vector onto another one
func (vector Vector) Project(other Vector) Vector {
	amt := vector.Dot(other) / other.Len2()
	vector.X = amt * other.X
	vector.Y = amt * other.Y
	return vector
}

//ProjectN this vector onto a unit vector
func (vector Vector) ProjectN(other Vector) Vector {
	amt := vector.Dot(other)
	vector.X = amt * other.X
	vector.Y = amt * other.Y
	return vector
}

//Reflect this vector on an arbitrary axis vector
func (vector Vector) Reflect(axis Vector) Vector {
	x := vector.X
	y := vector.Y
	resultVector := vector.Project(axis).Scale(2)
	resultVector.X -= x
	resultVector.Y -= y
	return resultVector
}

//ReflectN this vector on an arbitrary axis unit vector
func (vector Vector) ReflectN(axis Vector) Vector {
	x := vector.X
	y := vector.Y
	resultVector := vector.ProjectN(axis).Scale(2)
	resultVector.X -= x
	resultVector.Y -= y
	return resultVector
}

//Dot returns the dot product of this vector and another
func (vector Vector) Dot(other Vector) float64 {
	return vector.X*other.X + vector.Y*other.Y
}

//Len2 returns the squared length of this vector
func (vector Vector) Len2() float64 {
	return vector.Dot(vector)
}

//Len returns the length of this vector
func (vector Vector) Len() float64 {
	return math.Sqrt(vector.Len2())
}
