package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoxString(t *testing.T) {
	box := collision2d.Box{}
	output := string(box.String())
	assert.Equal(t, "{Pos:{X:0.000000, Y:0.000000}\nWidth:0.000000\nHeight:0.000000}", output, "they should be equal")
}
