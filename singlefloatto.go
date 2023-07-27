package cram

import (
	"reflect"
	"fmt"
	"strconv"
)

// convSingleBoolFromSingleFloat: bool = float
func convSingleBoolFromSingleFloat(dst, src reflect.Value) error {
	return booltovalue(dst, !(valuetofloat64(src) == 0.0))
}

// convSingleIntFromSingleFloat: int = float
func convSingleIntFromSingleFloat(dst, src reflect.Value) error {
	return int64tovalue(dst, int64(valuetofloat64(src)))
}

// convSingleUintFromSingleFloat: uint = float
func convSingleUintFromSingleFloat(dst, src reflect.Value) error {
	return uint64tovalue(dst, uint64(valuetofloat64(src)))
}

// convSingleFloatFromSingleFloat: float = float
func convSingleFloatFromSingleFloat(dst, src reflect.Value) error {
	return float64tovalue(dst, valuetofloat64(src))
}

// convSingleStringFromSingleFloat: string = float
func convSingleStringFromSingleFloat(dst, src reflect.Value) error {
	return stringtovalue(dst, strconv.FormatFloat(valuetofloat64(src), 'f', -1, 64))
}

// convMultiBoolFromSingleFloat: []bool = float
func convMultiBoolFromSingleFloat(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleBoolFromSingleFloat)
}

// convMultiIntFromSingleFloat: []int = string
func convMultiIntFromSingleFloat(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleIntFromSingleFloat)
}

// convMultiUintFromSingleFloat: []uint = float
func convMultiUintFromSingleFloat(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleUintFromSingleFloat)
}

// convMultiFloatFromSingleFloat: []float = float
func convMultiFloatFromSingleFloat(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleFloatFromSingleFloat)
}

// convMultiStringFromSingleFloat: []string = float
func convMultiStringFromSingleFloat(dst, src reflect.Value) error {
	return singleListItemConv(dst, src, convSingleStringFromSingleFloat)
}

// floattovalue will set a Value (or childvalue to pointers) to float, allocating pointers where possible
func float64tovalue(v reflect.Value, x float64) error {
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
	v.SetFloat(x)
	return nil
}


// valuetofloat64 will get a float64 from Value, resolving pointers along the way
func valuetofloat64(v reflect.Value) float64 {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return 0.0
		}
		v = v.Elem()
	}
	return v.Float()
}


