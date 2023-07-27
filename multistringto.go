package cram

import (
	"reflect"
	"fmt"
	"strings"
)

// convSingleBoolFromMultiString: bool = []string
func convSingleBoolFromMultiString(dst, src reflect.Value) error {
	return convSingleBoolFromSingleString(dst, firstHelper(src))
}

// convSingleIntFromMultiString: int = []string
func convSingleIntFromMultiString(dst, src reflect.Value) error {
	return convSingleIntFromSingleString(dst, firstHelper(src))
}

// convSingleUintFromMultiString: uint = []string
func convSingleUintFromMultiString(dst, src reflect.Value) error {
	return convSingleUintFromSingleString(dst, firstHelper(src))
}

// convSingleFloatFromMultiString: float = []string
func convSingleFloatFromMultiString(dst, src reflect.Value) error {
	return convSingleFloatFromSingleString(dst, firstHelper(src))
}

// convSingleStringFromMultiString: string = []string
func convSingleStringFromMultiString(dst, src reflect.Value) error {
	var x []string
	if err := convMultiStringFromMultiString(reflect.ValueOf(&x), src); err != nil {
		return fmt.Errorf("copying strings for conversion: %w", err)
	}
	return stringtovalue(dst, strings.Join(x, ","))
}

// convMultiBoolFromMultiString: []bool = []string
func convMultiBoolFromMultiString(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleBoolFromSingleString)
}

// convMultiIntFromMultiString: []int = []string
func convMultiIntFromMultiString(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleIntFromSingleString)
}

// convMultiUintFromMultiString: []uint = []string
func convMultiUintFromMultiString(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleUintFromSingleString)
}

// convMultiFloatFromMultiString: []float = []string
func convMultiFloatFromMultiString(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleFloatFromSingleString)
}

// convMultiStringFromMultiString: []string = []string
func convMultiStringFromMultiString(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleStringFromSingleString)
}
