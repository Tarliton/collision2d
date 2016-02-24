package sat

import (
	"math"
)

type Vector struct {
	X, Y float64
}

func (vector Vector) Copy(other Vector) Vector {
	vector.X = other.X
	vector.Y = other.Y
	return vector
}

func (vector Vector) Clone() Vector {
	another := Vector{}.Copy(vector)
	return another
}

func (vector Vector) Perp() Vector {
	x := vector.X
	vector.X = vector.Y
	vector.Y = x
	return vector
}

func (vector Vector) Rotate(angle float64) Vector {
	x := vector.X
	y := vector.Y
	vector.X = x*math.Cos(angle) - y*math.Sin(angle)
	vector.Y = x*math.Sin(angle) + y*math.Cos(angle)
	return vector
}

func (vector Vector) Reverse() Vector {
	vector.X = -vector.X
	vector.Y = -vector.Y
	return vector
}

func (vector Vector) Normalize() Vector {
	d := vector.Len()
	if d > 0 {
		vector.X /= d
		vector.Y /= d
	}
	return vector
}

func (vector Vector) Add(other Vector) Vector {
	vector.X += other.X
	vector.Y += other.Y
	return vector
}

func (vector Vector) Sub(other Vector) Vector {
	vector.X -= other.X
	vector.Y -= other.Y
	return vector
}

func (vector Vector) Scale(x float64) Vector {
	vector.X *= x
	vector.Y *= x
	return vector
}

func (vector Vector) ScaleDifferent(x, y float64) Vector {
	vector.X *= x
	vector.Y *= y
	return vector
}

func (vector Vector) Project(other Vector) Vector {
	amt := vector.Dot(other) / other.Len2()
	vector.X = amt * other.X
	vector.Y = amt * other.Y
	return vector
}

func (vector Vector) ProjectN(other Vector) Vector {
	amt := vector.Dot(other)
	vector.X = amt * other.X
	vector.Y = amt * other.Y
	return vector
}

func (vector Vector) Reflect(axis Vector) Vector {
	x := vector.X
	y := vector.Y
	resultVector := vector.Project(axis).Scale(2)
	resultVector.X -= x
	resultVector.Y -= y
	return resultVector
}

func (vector Vector) ReflectN(axis Vector) Vector {
	x := vector.X
	y := vector.Y
	resultVector := vector.ProjectN(axis).Scale(2)
	resultVector.X -= x
	resultVector.Y -= y
	return resultVector
}

func (vector Vector) Dot(other Vector) float64 {
	return vector.X*other.X + vector.Y*other.Y
}

func (vector Vector) Len2() float64 {
	return vector.Dot(vector)
}

func (vector Vector) Len() float64 {
	return math.Sqrt(vector.Len2())
}
