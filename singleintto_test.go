package cram

import (
	"testing"
)

func TestSingleIntTo(t *testing.T) {
	t.Parallel()

	// bool := int
	helper(t, false, 0)
	helper(t, true, 1)
	helper(t, true, -1)
	helper(t, true, 8)

	// int := int
	helper(t, 0, 0)
	helper(t, -1, -1)
	helper(t, 1, 1)
	helper(t, 7, 7)
	helper(t, 999999999, 999999999)

	// uint := int
	helper(t, uint(0), 0)
	helper(t, uint(1), 1)
	helper(t, uint(7), 7)
	helper(t, uint64(18446744073709551615), -1)
	helper(t, uint32(4294967295), -1)

	// float = int
	helper(t, float64(0), 0)
	helper(t, float32(0), 0)
	helper(t, float32(5.0), 5)
	helper(t, float64(-11.0), -11)

	// string = int
	helper(t, "0", 0)
	helper(t, "1", 1)
	helper(t, "-1", -1)
	helper(t, "-987654321", -987654321)

	// []bool := int
	helper(t, []bool{false}, 0)
	helper(t, []bool{true}, 1)
	helper(t, []bool{true}, -1)
	helper(t, []bool{true}, 8)

	// []int := int
	helper(t, []int{0}, 0)
	helper(t, []int{-1}, -1)
	helper(t, []int{1}, 1)
	helper(t, []int{7}, 7)
	helper(t, []int{999999999}, 999999999)

	// []uint := int
	helper(t, []uint{0}, 0)
	helper(t, []uint{1}, 1)
	helper(t, []uint{7}, 7)
	helper(t, []uint64{18446744073709551615}, -1)
	helper(t, []uint32{4294967295}, -1)

	// []float = int
	helper(t, []float64{0}, 0)
	helper(t, []float32{0}, 0)
	helper(t, []float32{5.0}, 5)
	helper(t, []float64{-11.0}, -11)

	// []string = int
	helper(t, []string{"0"}, 0)
	helper(t, []string{"1"}, 1)
	helper(t, []string{"-1"}, -1)
	helper(t, []string{"-987654321"}, -987654321)

}
