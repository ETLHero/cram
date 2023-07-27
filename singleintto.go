package cram

import (
	"reflect"
	"fmt"
	"strconv"
)

// convSingleBoolFromSingleInt: bool = int
func convSingleBoolFromSingleInt(dst, src reflect.Value) error {
	return booltovalue(dst, !(valuetoint64(src) == 0))
}

// convSingleIntFromSingleInt: int = int
func convSingleIntFromSingleInt(dst, src reflect.Value) error {
	return int64tovalue(dst, valuetoint64(src))
}

// convSingleIntFromSingleInt: uint = int
func convSingleUintFromSingleInt(dst, src reflect.Value) error {
	return uint64tovalue(dst, uint64(valuetoint64(src)))
}

// convSingleFloatFromSingleInt: float = int
func convSingleFloatFromSingleInt(dst, src reflect.Value) error {
	return float64tovalue(dst, float64(valuetoint64(src)))
}

// convSingleStringFromSingleInt: string = int
func convSingleStringFromSingleInt(dst, src reflect.Value) error {
	return stringtovalue(dst, strconv.FormatInt(valuetoint64(src), 10))
}

// convMultiBoolFromSingleInt: []bool = int
func convMultiBoolFromSingleInt(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleBoolFromSingleInt)
}

// convMultiIntFromSingleInt: []int = int
func convMultiIntFromSingleInt(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleIntFromSingleInt)
}

// convMultiUintFromSingleInt: []uint = int
func convMultiUintFromSingleInt(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleUintFromSingleInt)
}

// convMultiFloatFromSingleInt: []float = int
func convMultiFloatFromSingleInt(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleFloatFromSingleInt)
}

// convMultiStringFromSingleInt: []string = int
func convMultiStringFromSingleInt(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleStringFromSingleInt)
}

////////////////////////////// Uint //////////////////////////////////////////

// convSingleBoolFromSingleUint: bool = uint
func convSingleBoolFromSingleUint(dst, src reflect.Value) error {
	return booltovalue(dst, !(valuetouint64(src) == 0))
}

// convSingleIntFromSingleUint: int = uint
func convSingleIntFromSingleUint(dst, src reflect.Value) error {
	return int64tovalue(dst, int64(valuetouint64(src)))
}

// convSingleUintFromSingleUint: uint = uint
func convSingleUintFromSingleUint(dst, src reflect.Value) error {
	return uint64tovalue(dst, valuetouint64(src))
}

// convSingleFloatFromSingleUint: float = uint
func convSingleFloatFromSingleUint(dst, src reflect.Value) error {
	return float64tovalue(dst, float64(valuetouint64(src)))
}

// convSingleStringFromSingleUint: string = uint
func convSingleStringFromSingleUint(dst, src reflect.Value) error {
	// Consider a unicode conversion
	return stringtovalue(dst, strconv.FormatUint(valuetouint64(src), 10))
}

// convMultiBoolFromSingleUint: []bool = uint
func convMultiBoolFromSingleUint(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleBoolFromSingleUint)
}

// convMultiIntFromSingleUint: []int = uint
func convMultiIntFromSingleUint(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleIntFromSingleUint)
}

// convMultiUintFromSingleUint: []uint = uint
func convMultiUintFromSingleUint(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleUintFromSingleUint)
}

// convMultiFloatFromSingleUint: []float = uint
func convMultiFloatFromSingleUint(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleFloatFromSingleUint)
}

// convMultiStringFromSingleUint: []string = uint
func convMultiStringFromSingleUint(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleStringFromSingleUint)
}

// int64tovalue will set a Value (or childvalue to pointers) to int64, allocating pointers where possible
func int64tovalue(v reflect.Value, x int64) error {
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
		return ErrCannotSet
	}
	v.SetInt(x)
	return nil
}

// uint64tovalue will set a Value (or childvalue to pointers) to uint64, allocating pointers where possible
func uint64tovalue(v reflect.Value, x uint64) error {
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
		return ErrCannotSet
	}
	v.SetUint(x)
	return nil
}

// valuetofloat64 will get a float64 from Value, resolving pointers along the way
func valuetoint64(v reflect.Value) int64 {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return 0
		}
		v = v.Elem()
	}
	return v.Int()
}

// valuetofloat64 will get a float64 from Value, resolving pointers along the way
func valuetouint64(v reflect.Value) uint64 {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return 0
		}
		v = v.Elem()
	}
	return v.Uint()
}
