package cram

import (
	"testing"
)

func TestMultiFloatTo(t *testing.T) {
	t.Parallel()

	// bool := []float
	helper(t, true, []float32{1.0,0.0})
	helper(t, false, []float32{0.0})
}
