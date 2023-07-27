package cram

import (
	"testing"
)

func TestMultiBoolTo(t *testing.T) {
	t.Parallel()

	// bool := []bool
	helper(t, true, []bool{true,false})
	helper(t, false, []bool{false})

	// int := []bool
	helper(t, 1, []bool{true,false})
	helper(t, 0, []bool{false})

	// uint := []bool
	helper(t, uint(1), []bool{true,false})
	helper(t, uint(0), []bool{false})

	// float := []bool
	helper(t, 1.0, []bool{true,false})
	helper(t, 0.0, []bool{false})

	// string := []bool
	helper(t, "true,false", []bool{true,false})
	helper(t, "false", []bool{false})

	// []bool := []bool
	helper(t, []bool{true,false}, []bool{true,false})
	helper(t, []bool{false}, []bool{false})

	// []int := []bool
	helper(t, []int{1,0}, []bool{true,false})
	helper(t, []int{0}, []bool{false})

	// []uint := []bool
	helper(t, []uint{1,0}, []bool{true,false})
	helper(t, []uint{0}, []bool{false})

	// []float := []bool
	helper(t, []float64{1.0,0.0}, []bool{true,false})
	helper(t, []float64{0.0}, []bool{false})

	// []string := []bool
	helper(t, []string{"true","false"}, []bool{true,false})
	helper(t, []string{"false"}, []bool{false})
}
