package cram

import (
	"testing"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func helper[e any](t *testing.T, expected e, provided any) {
	t.Helper()
	var dst e
	if err := Into(&dst, provided); err != nil {
		t.Errorf("Unexpected error running conversion function: %s", err)
		return
	}
	if msg := cmp.Diff(expected, dst); msg != "" {
		t.Error(msg)
	}
}

func ExampleInto() {
	inital := []byte{0x37,0x32,0x2c,0x31,0x30,0x31,0x2c,0x31,0x30,
	0x38,0x2c,0x31,0x30,0x38,0x2c,0x31,0x31,0x31,0x2c,0x33,0x32,
	0x2c,0x31,0x32,0x37,0x37,0x35,0x38}
	var str string
	var uni []uint
	var final string
	
	Into(&str, inital)
	fmt.Println(str)
	Into(&uni, str)
	fmt.Println(uni)
	Into(&final, uni)
	fmt.Println(final)

	// Output:
	// 72,101,108,108,111,32,127758
	// [72 101 108 108 111 32 127758]
	// Hello 🌎
}
 
func TestPointersDst(t *testing.T) {
	t.Parallel()
	deep := new(**int)
	if msg := cmp.Diff(Into(deep, 777), nil); msg != "" {
		t.Fatal(msg)
	}
	if deep == nil {
		t.Fatal("nil at depth 0")
	}
	if *deep == nil {
		t.Fatal("nil at depth 1")
	}
	if **deep == nil {
		t.Fatal("nil at depth 2")
	}
	if msg := cmp.Diff(***deep, 777); msg != "" {
		t.Error(msg)
	}
}

func TestPointersSrc(t *testing.T) {
	t.Parallel()
	source1 := 888
	source2 := &source1
	source3 := &source2
	var simple int
	if msg := cmp.Diff(Into(&simple, source3), nil); msg != "" {
		t.Fatal(msg)
	}
	if msg := cmp.Diff(simple, source1); msg != "" {
		t.Error(msg)
	}
}
