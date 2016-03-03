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

//Copy the value of another vector to a new one.
func (vector Vector) Copy(another Vector) Vector {
	return Vector{another.X, another.Y}
}

//Clone this vector coordinates to a new vector with the same coordinates as this one.
func (vector Vector) Clone() Vector {
	return Vector{}.Copy(vector)
}

//Perp returns a new vector perpendicular from this one.
func (vector Vector) Perp() Vector {
	return Vector{vector.Y, -vector.X}
}

//Rotate returns a new vector rotated counter-clockwise by the specified number of radians.
func (vector Vector) Rotate(angle float64) Vector {
	return Vector{
		vector.X*math.Cos(angle) - vector.Y*math.Sin(angle),
		vector.X*math.Sin(angle) + vector.Y*math.Cos(angle)}
}

//Reverse returns a new vector that is reversed from this one.
func (vector Vector) Reverse() Vector {
	return Vector{-vector.X, -vector.Y}
}

//Normalize returns a new unit-length vector.
func (vector Vector) Normalize() Vector {
	d := vector.Len()
	return Vector{vector.X / d, vector.Y / d}
}

//Add returns a new vector with the result of adding another vector to this one.
func (vector Vector) Add(another Vector) Vector {
	return Vector{vector.X + another.X, vector.Y + another.Y}
}

//Sub returns a new vector that is the result of subtracting another vector from this one.
func (vector Vector) Sub(another Vector) Vector {
	return Vector{vector.X - another.X, vector.Y - another.Y}
}

//Scale returns a new vector scaled in the direction of X and Y by value x
func (vector Vector) Scale(x float64) Vector {
	return Vector{vector.X * x, vector.Y * x}
}

//ScaleDifferent returns a new vector scaled in the direction of X and Y by value x and y respectively
func (vector Vector) ScaleDifferent(x, y float64) Vector {
	return Vector{vector.X * x, vector.Y * y}
}

//Project this vector onto another one
func (vector Vector) Project(another Vector) Vector {
	amt := vector.Dot(another) / another.Len2()
	return Vector{amt * another.X, amt * another.Y}
}

//ProjectN this vector onto a unit vector
func (vector Vector) ProjectN(another Vector) Vector {
	amt := vector.Dot(another)
	return Vector{amt * another.X, amt * another.Y}
}

//Reflect this vector on an arbitrary axis vector
func (vector Vector) Reflect(axis Vector) Vector {
	return vector.Project(axis).Scale(2).Sub(vector)
}

//ReflectN this vector on an arbitrary axis unit vector
func (vector Vector) ReflectN(axis Vector) Vector {
	return vector.ProjectN(axis).Scale(2).Sub(vector)
}

//Dot returns the dot product of this vector and another
func (vector Vector) Dot(another Vector) float64 {
	return vector.X*another.X + vector.Y*another.Y
}

//Len2 returns the squared length of this vector
func (vector Vector) Len2() float64 {
	return vector.Dot(vector)
}

//Len returns the length of this vector
func (vector Vector) Len() float64 {
	return math.Sqrt(vector.Len2())
}
