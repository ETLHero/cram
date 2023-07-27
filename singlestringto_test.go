package cram

import (
	"testing"
)

func TestSingleStringTo(t *testing.T) {
	t.Parallel()

	// bool := string
	helper(t, true, "true")
	helper(t, true, "t")
	helper(t, true, "1")
	helper(t, false, "false")
	helper(t, false, "f")
	helper(t, false, "0")
	helper(t, false, "")
	helper(t, true, "ou3b023br02b3i")

	// int := string
	helper(t, 0, "0")
	helper(t, 1, "1")
	helper(t, -1, "-1")
	helper(t, -888, "-888")

	// uint := string
	helper(t, uint(0), "0")
	helper(t, uint(1), "1")
	helper(t, uint(808), "808")

	// float := string
	helper(t, float64(0), "0000.00000")
	helper(t, float64(1.0), "001.00")
	helper(t, float64(6.0), "6.00")
	helper(t, float64(-11.22), "-11.22")
	helper(t, float64(-12.55), "-012.550")

	// string := string
	helper(t, ":D", ":D")
	helper(t, "💩", "💩")
	helper(t, "0", "0")
	helper(t, "", "")
	helper(t, "This is the story for a girl...", "This is the story for a girl...")

	// []bool := string
	helper(t, []bool{true,true,false,true}, "true,1,,t")
	helper(t, []bool{true,false,true}, "wiqdbqiubdwiu,f,falsefalse")
	helper(t, []bool(nil), "")

	// []int := string
	helper(t, []int{1,0,-45,222}, "1,0,-45,222")
	helper(t, []int{5}, "5")
	helper(t, []int(nil), "")

	// []uint := string
	helper(t, []uint{1,0,222}, "1,0,222")
	helper(t, []uint(nil), "")

	// []float := string
	helper(t, []float64{1,0.0,101.202}, "01,0,101.202")
	helper(t, []float64(nil), "")

	// []string := string
	helper(t, []string{"abc","_"}, "abc,_")
	helper(t, []string(nil), "")
}
