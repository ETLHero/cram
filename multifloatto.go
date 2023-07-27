package cram

import (
	"reflect"
	"strings"
)

// convSingleBoolFromMultiFloat: bool = []float
func convSingleBoolFromMultiFloat(dst, src reflect.Value) error {
	return convSingleBoolFromSingleFloat(dst, firstHelper(src))
}

// convSingleIntFromMultiFloat: int = []float
func convSingleIntFromMultiFloat(dst, src reflect.Value) error {
	return convSingleIntFromSingleFloat(dst, firstHelper(src))
}

// convSingleUintFromMultiFloat: uint = []float
func convSingleUintFromMultiFloat(dst, src reflect.Value) error {
	return convSingleUintFromSingleFloat(dst, firstHelper(src))
}

// convSingleFloatFromMultiFloat: float = []float
func convSingleFloatFromMultiFloat(dst, src reflect.Value) error {
	return convSingleFloatFromSingleFloat(dst, firstHelper(src))
}

// convSingleStringFromMultiFloat: string = []float
func convSingleStringFromMultiFloat(dst, src reflect.Value) error {
	strs := make([]string, src.Len())
	err := convMultiStringFromMultiFloat(reflect.ValueOf(&strs), src)
	if err != nil {
		return err
	}
	joined := strings.Join(strs, ",")
	return convSingleStringFromSingleString(dst, reflect.ValueOf(joined))
}

// convMultiBoolFromMultiFloat: []bool = []float
func convMultiBoolFromMultiFloat(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleBoolFromSingleFloat)
}

// convMultiIntFromMultiFloat: []int = []float
func convMultiIntFromMultiFloat(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleIntFromSingleFloat)
}

// convMultiUintFromMultiFloat: []uint = []float
func convMultiUintFromMultiFloat(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleUintFromSingleFloat)
}

// convMultiFloatFromMultiFloat: []float = []float
func convMultiFloatFromMultiFloat(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleFloatFromSingleFloat)
}

// convMultiStringFromMultiFloat: []string = []float
func convMultiStringFromMultiFloat(dst, src reflect.Value) error {
	return loopHelper(dst, src, convSingleStringFromSingleFloat)
}
