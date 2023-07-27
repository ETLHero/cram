package cram

import (
	"reflect"
	"strings"
)

// convSingleBoolFromMultiBool: bool = []bool
func convSingleBoolFromMultiBool(dst, src reflect.Value) error {
	return convSingleBoolFromSingleBool(dst, firstHelper(src))
}

// convSingleIntFromMultiBool: int = []bool
func convSingleIntFromMultiBool(dst, src reflect.Value) error {
	return convSingleIntFromSingleBool(dst, firstHelper(src))
}

// convSingleUintFromMultiBool: uint = []bool
func convSingleUintFromMultiBool(dst, src reflect.Value) error {
	return convSingleUintFromSingleBool(dst, firstHelper(src))
}

// convSingleFloatFromMultiBool: float = []bool
func convSingleFloatFromMultiBool(dst, src reflect.Value) error {
	return convSingleFloatFromSingleBool(dst, firstHelper(src))
}

// convSingleStringFromMultiBool: string = []bool
func convSingleStringFromMultiBool(dst, src reflect.Value) error {
	strs := make([]string, src.Len())
	err := convMultiStringFromMultiBool(reflect.ValueOf(&strs), src)
	if err != nil {
		return err
	}
	joined := strings.Join(strs, ",")
	return convSingleStringFromSingleString(dst, reflect.ValueOf(joined))
}

// convMultiBoolFromMultiBool: []bool = []bool
func convMultiBoolFromMultiBool(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleBoolFromSingleBool)
}

// convMultiIntFromMultiBool: []int = []bool
func convMultiIntFromMultiBool(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleIntFromSingleBool)
}

// convMultiUintFromMultiBool: []uint = []bool
func convMultiUintFromMultiBool(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleUintFromSingleBool)
}

// convMultiFloatFromMultiBool: []float = []bool
func convMultiFloatFromMultiBool(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleFloatFromSingleBool)
}

// convMultiStringFromMultiBool: []string = []bool
func convMultiStringFromMultiBool(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleStringFromSingleBool)
}
