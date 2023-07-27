package cram

import (
	"testing"
)

func TestSingleBoolTo(t *testing.T) {
	t.Parallel()

	// bool := bool
	helper(t, true, true)
	helper(t, false, false)
	
	// int := bool
	helper(t, int(0), false)
	helper(t, int(1), true)

	// uint := bool
	helper(t, uint(0), false)
	helper(t, uint(1), true)

	// float := bool
	helper(t, float32(0.0), false)
	helper(t, float32(1.0), true)

	// string := bool
	helper(t, "false", false)
	helper(t, "true", true)

	// []bool := bool
	helper(t, []bool{true}, true)
	helper(t, []bool{false}, false)
	
	// []int := bool
	helper(t, []int{0}, false)
	helper(t, []int{1}, true)

	// []uint := bool
	helper(t, []uint{0}, false)
	helper(t, []uint{1}, true)

	// []float := bool
	helper(t, []float64{0.0}, false)
	helper(t, []float64{1.0}, true)

	// []string := bool
	helper(t, []string{"false"}, false)
	helper(t, []string{"true"}, true)
}

