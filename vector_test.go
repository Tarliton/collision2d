package collision2d_test

import (
	// "fmt"
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestVectorString(t *testing.T) {
	vector := collision2d.NewVector(0, 0)
	output := string(vector.String())
	assert.Equal(t, "{X:0.000000, Y:0.000000}\n", output, "they should be equal")
}

func TestVectorNewVectorZeros(t *testing.T) {
	vector := collision2d.NewVector(0, 0)
	assert.Equal(t, float64(0), vector.X, "they should be equal")
	assert.Equal(t, float64(0), vector.Y, "they should be equal")
}

func TestVectorNewVectorNumbers(t *testing.T) {
	vector := collision2d.NewVector(3.46, 2.18)
	assert.Equal(t, float64(3.46), vector.X, "they should be equal")
	assert.Equal(t, float64(2.18), vector.Y, "they should be equal")
}

func TestVectorCopy(t *testing.T) {
	vector1 := collision2d.NewVector(3.46, 2.18)
	vector2 := collision2d.NewVector(-5.1, 4.8)

	vector3 := vector1.Copy(vector2)

	assert.Equal(t, float64(3.46), vector1.X, "they should be equal")
	assert.Equal(t, float64(2.18), vector1.Y, "they should be equal")

	assert.Equal(t, float64(-5.1), vector2.X, "they should be equal")
	assert.Equal(t, float64(4.8), vector2.Y, "they should be equal")

	assert.Equal(t, float64(-5.1), vector3.X, "they should be equal")
	assert.Equal(t, float64(4.8), vector3.Y, "they should be equal")
}

func TestVectorClone(t *testing.T) {
	vector1 := collision2d.NewVector(3.46, 2.18)

	vector2 := vector1.Clone()

	assert.Equal(t, float64(3.46), vector2.X, "they should be equal")
	assert.Equal(t, float64(2.18), vector2.Y, "they should be equal")
}

func TestVectorPerp(t *testing.T) {
	vector1 := collision2d.NewVector(3.46, 2.18)
	vector2 := vector1.Perp()
	assert.Equal(t, float64(2.18), vector2.X, "X should be 1")
	assert.Equal(t, float64(-3.46), vector2.Y, "Y should be 0")
}

func TestVectorRotate(t *testing.T) {
	vector := collision2d.NewVector(5.5, 4.7)
	vector2 := vector.Rotate(math.Pi)
	check1 := nearlyEqual(vector2.X, float64(-5.5), float64(0.00001))
	check2 := nearlyEqual(vector2.Y, float64(-4.7), float64(0.00001))
	assert.Equal(t, true, check1, "should be true")
	assert.Equal(t, true, check2, "should be true")
}

func TestVectorReverse(t *testing.T) {
	vector := collision2d.NewVector(1, -5)
	vector2 := vector.Reverse()
	assert.Equal(t, float64(-1), vector2.X, "should be -1")
	assert.Equal(t, float64(5), vector2.Y, "should be 5")
}
func TestVectorNormalize(t *testing.T) {
	vector := collision2d.NewVector(41, 123.123)
	vector2 := vector.Normalize()
	assert.Equal(t, float64(1), vector2.Len(), "length should be 1")
}
func TestVectorAdd(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := collision2d.NewVector(25, 45)
	vector3 := vector1.Add(vector2)
	assert.Equal(t, float64(40), vector3.X, "should sum right")
	assert.Equal(t, float64(50), vector3.Y, "should sum right")
}
func TestVectorSub(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := collision2d.NewVector(25, 45)
	vector3 := vector1.Sub(vector2)
	assert.Equal(t, float64(-10), vector3.X, "should sum right")
	assert.Equal(t, float64(-40), vector3.Y, "should sum right")
}
func TestVectorScale(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := vector1.Scale(2)
	assert.Equal(t, float64(30), vector2.X, "should sum right")
	assert.Equal(t, float64(10), vector2.Y, "should sum right")
}
func TestVectorScaleDifferent(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := vector1.ScaleDifferent(2, 5)
	assert.Equal(t, float64(30), vector2.X, "should sum right")
	assert.Equal(t, float64(25), vector2.Y, "should sum right")
}
func TestVectorProject(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := collision2d.NewVector(0, 20)
	vector3 := vector1.Project(vector2)
	assert.Equal(t, float64(0), vector3.X, "should sum right")
	assert.Equal(t, float64(5), vector3.Y, "should sum right")
}
func TestVectorProjectN(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := collision2d.NewVector(0, 20)
	vector3 := vector1.ProjectN(vector2)
	assert.Equal(t, float64(0), vector3.X, "should sum right")
	assert.Equal(t, float64(2000), vector3.Y, "should sum right")
}
func TestVectorReflect(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := collision2d.NewVector(0, 20)
	vector3 := vector1.Reflect(vector2)
	assert.Equal(t, float64(-15), vector3.X, "should sum right")
	assert.Equal(t, float64(5), vector3.Y, "should sum right")
}
func TestVectorReflectN(t *testing.T) {
	vector1 := collision2d.NewVector(15, 5)
	vector2 := collision2d.NewVector(0, 20)
	vector3 := vector1.ReflectN(vector2)
	assert.Equal(t, float64(-15), vector3.X, "should sum right")
	assert.Equal(t, float64(3995), vector3.Y, "should sum right")
}
func TestVectorDot(t *testing.T) {
	vector1 := collision2d.NewVector(-4, -9)
	vector2 := collision2d.NewVector(-1, 2)
	result := vector1.Dot(vector2)
	assert.Equal(t, float64(-14), result, "should sum right")
}
func TestVectorLen(t *testing.T) {
	vector1 := collision2d.NewVector(0, 5)
	result := vector1.Len()
	assert.Equal(t, float64(5), result, "should sum right")
}
func TestVectorLen2(t *testing.T) {
	vector1 := collision2d.NewVector(0, 5)
	result := vector1.Len2()
	assert.Equal(t, float64(25), result, "should sum right")
}

func nearlyEqual(a, b, epsilon float64) bool {
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)
	if a == b {
		return true
	} else if a == 0 || b == 0 || diff < math.SmallestNonzeroFloat64 {
		return diff < (epsilon * math.SmallestNonzeroFloat64)
	} else {
		return (diff / (absA + absB)) < epsilon
	}
}
