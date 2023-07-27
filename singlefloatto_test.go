package cram

import (
	"testing"
)

func TestSingleFloatTo(t *testing.T) {
	t.Parallel()

	// bool := float
	helper(t, false, 0.0)
	helper(t, true, 1.0)
	helper(t, true, -0.2)
	helper(t, true, -2.0)
	helper(t, true, 7.0)
	helper(t, true, 1.01)
	
	// int := float
	helper(t, int(0), 0.0)
	helper(t, int(1), 1.0)
	helper(t, int(-1), -1.0)
	helper(t, int(-9), -9.02)
	helper(t, int(12), 12.98)

	// uint := float
	helper(t, uint(0), 0.0)
	helper(t, uint(1), 1.0)
	helper(t, uint(19), 19.9991)
	helper(t, uint(24), 24.1)

	// float := float
	helper(t, float64(0.0), 0.0)
	helper(t, float64(1.0), 1.0)
	helper(t, float64(1.3), 1.3)
	helper(t, float64(-22.8), -22.8)

	// string := float
	helper(t, "1", 1.0)
	helper(t, "0", 0.0)
	helper(t, "-1", -1.0)
	helper(t, "-3.3", -3.3)
	helper(t, "9.98", 9.98)

	// []bool := float
	helper(t, []bool{false}, 0.0)
	helper(t, []bool{true}, 1.0)
	helper(t, []bool{true}, -0.2)
	helper(t, []bool{true}, -2.0)
	helper(t, []bool{true}, 7.0)
	helper(t, []bool{true}, 1.01)
	
	// []int := float
	helper(t, []int{0}, 0.0)
	helper(t, []int{1}, 1.0)
	helper(t, []int{-1}, -1.0)
	helper(t, []int{-9}, -9.02)
	helper(t, []int{12}, 12.98)

	// []uint := float
	helper(t, []uint{0}, 0.0)
	helper(t, []uint{1}, 1.0)
	helper(t, []uint{19}, 19.9991)
	helper(t, []uint{24}, 24.1)

	// []float := float
	helper(t, []float64{0.0}, 0.0)
	helper(t, []float64{1.0}, 1.0)
	helper(t, []float64{1.3}, 1.3)
	helper(t, []float64{-22.8}, -22.8)

	// []string := float
	helper(t, []string{"1"}, 1.0)
	helper(t, []string{"0"}, 0.0)
	helper(t, []string{"-1"}, -1.0)
	helper(t, []string{"-3.3"}, -3.3)
	helper(t, []string{"9.98"}, 9.98)
}

