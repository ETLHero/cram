package cram

import (
	"reflect"
	"strings"
	"fmt"
)

// convSingleBoolFromMultiInt: bool = []int
func convSingleBoolFromMultiInt(dst, src reflect.Value) error {
	return convSingleBoolFromSingleInt(dst, firstHelper(src))
}

// convSingleIntFromMultiInt: int = []int
func convSingleIntFromMultiInt(dst, src reflect.Value) error {
	return convSingleIntFromSingleInt(dst, firstHelper(src))
}

// convSingleUintFromMultiInt: uint = []int
func convSingleUintFromMultiInt(dst, src reflect.Value) error {
	return convSingleUintFromSingleInt(dst, firstHelper(src))
}

// convSingleFloatFromMultiInt: float = []int
func convSingleFloatFromMultiInt(dst, src reflect.Value) error {
	return convSingleFloatFromSingleInt(dst, firstHelper(src))
}

// convSingleStringFromMultiInt: string = []int
func convSingleStringFromMultiInt(dst, src reflect.Value) error {
	strs := make([]string, src.Len())
	err := convMultiStringFromMultiInt(reflect.ValueOf(&strs), src)
	if err != nil {
		return err
	}
	joined := strings.Join(strs, ",")
	return convSingleStringFromSingleString(dst, reflect.ValueOf(joined))
}

// convMultiBoolFromMultiInt: []bool = []int
func convMultiBoolFromMultiInt(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleBoolFromSingleInt)
}

// convMultiIntFromMultiInt: []int = []int
func convMultiIntFromMultiInt(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleIntFromSingleInt)
}

// convMultiUintFromMultiInt: []uint = []int
func convMultiUintFromMultiInt(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleUintFromSingleInt)
}

// convMultiStringFromMultiInt: []string = []int
func convMultiStringFromMultiInt(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleStringFromSingleInt)
}

// convMultiFloatFromMultiInt: []float = []int
func convMultiFloatFromMultiInt(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleFloatFromSingleInt)
}

////////////////////////////// Uint /////////////////////////////////////////


// convSingleBoolFromMultiUint: bool = []uint
func convSingleBoolFromMultiUint(dst, src reflect.Value) error {
	return convSingleBoolFromSingleUint(dst, firstHelper(src))
}

// convSingleIntFromMultiUint: int = []uint
func convSingleIntFromMultiUint(dst, src reflect.Value) error {
	return convSingleIntFromSingleUint(dst, firstHelper(src))
}

// convSingleUintFromMultiUint: uint = []uint
func convSingleUintFromMultiUint(dst, src reflect.Value) error {
	return convSingleUintFromSingleUint(dst, firstHelper(src))
}

// convSingleFloatFromMultiUint: float = []uint
func convSingleFloatFromMultiUint(dst, src reflect.Value) error {
	return convSingleFloatFromSingleUint(dst, firstHelper(src))
}

// convSingleStringFromMultiUint: string = []uint
func convSingleStringFromMultiUint(dst, src reflect.Value) error {
	_, subkind := Resolve(src.Type())
	if subkind == reflect.Uint8 {
		var x []byte
		if err := convMultiUintFromMultiUint(reflect.ValueOf(&x), src); err != nil {
			return fmt.Errorf("copying bytes for string conversion: %w", err)
		}
		return stringtovalue(dst, string(x))
	}

	var x []uint64
	if err := convMultiUintFromMultiUint(reflect.ValueOf(&x), src); err != nil {
		return fmt.Errorf("copying runes for string conversion: %w", err)
	}
	rs := make([]rune, len(x))
	for i := range x {
		rs[i] = rune(x[i])
	}
	return stringtovalue(dst, string(rs))
}

// convMultiBoolFromMultiUint: []bool = []uint
func convMultiBoolFromMultiUint(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleBoolFromSingleUint)
}

// convMultiIntFromMultiUint: []int = []uint
func convMultiIntFromMultiUint(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleIntFromSingleUint)
}

// convMultiUintFromMultiUint: []uint = []uint
func convMultiUintFromMultiUint(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleUintFromSingleUint)
}

// convMultiFloatFromMultiUint: []float = []uint
func convMultiFloatFromMultiUint(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleFloatFromSingleUint)
}

// convMultiStringFromMultiUint: []string = []uint
func convMultiStringFromMultiUint(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleStringFromSingleUint)
}
