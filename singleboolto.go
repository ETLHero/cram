package cram

import (
	"reflect"
	"fmt"
	"strconv"
)

// convSingleBoolFromSingleBool: bool = bool
func convSingleBoolFromSingleBool(dst, src reflect.Value) error {
	return booltovalue(dst, valuetobool(src))
}

// convSingleIntFromSingleBool: int = bool
func convSingleIntFromSingleBool(dst, src reflect.Value) error {
	if valuetobool(src) {
		return int64tovalue(dst, 1)
	}
	return int64tovalue(dst, 0)
}

// convSingleUintFromSingleBool: uint = bool
func convSingleUintFromSingleBool(dst, src reflect.Value) error {
	if valuetobool(src) {
		return uint64tovalue(dst, 1)
	}
	return uint64tovalue(dst, 0)
}

// convSingleFloatFromSingleBool: float = bool
func convSingleFloatFromSingleBool(dst, src reflect.Value) error {
	if valuetobool(src) {
		return float64tovalue(dst, 1.0)
	}
	return float64tovalue(dst, 0.0)
}

// convSingleStringFromSingleBool: string = bool
func convSingleStringFromSingleBool(dst, src reflect.Value) error {
	return stringtovalue(dst, strconv.FormatBool(valuetobool(src)))
}

// convMultiBoolFromSingleBool: []bool = bool
func convMultiBoolFromSingleBool(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleBoolFromSingleBool)
}

// convMultiIntFromSingleBool: []int = bool
func convMultiIntFromSingleBool(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleIntFromSingleBool)
}

// convMultiUintFromSingleBool: []uint = bool
func convMultiUintFromSingleBool(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleUintFromSingleBool)
}

// convMultiFloatFromSingleBool: []float = bool
func convMultiFloatFromSingleBool(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleFloatFromSingleBool)
}

// convMultiStringFromSingleBool: []string = bool
func convMultiStringFromSingleBool(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleStringFromSingleBool)
}

// valuetobool will retreve a bool from Value, treversing pointers where needed
func valuetobool(v reflect.Value) bool {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}
	return v.Bool()
}

// booltovalue will set a Value (or childvalue to pointers) to bool, allocating pointers where possible
func booltovalue(v reflect.Value, x bool) error {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			if !v.CanSet() {
				return fmt.Errorf("nil pointer in destination: %w", ErrCannotSet)
			}
			n := reflect.New(v.Type().Elem())
			v.Set(n)
		}
		v = v.Elem()
	}

	if !v.CanSet() {
		return fmt.Errorf("%w: %#v to %#v", ErrCannotSet, x, v.Interface())
	}
	v.SetBool(x)
	return nil
}
