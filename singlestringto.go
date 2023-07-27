package cram

import (
	"fmt"
	"reflect"
	"strconv"
	"errors"
	"strings"
)

var (
	WarnEmptyValueIsFalse = errors.New("empty value, treating as false")
	WarnBoolParseError = errors.New("unable to parse bool from string")
)

// convSingleBoolFromSingleString: bool = string
func convSingleBoolFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return booltovalue(dst, false)
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		// could not parse, and is not empty
		return booltovalue(dst, true)
	}
	return booltovalue(dst, b)
}

// convSingleIntFromSingleString: int = string
func convSingleIntFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return int64tovalue(dst, 0)
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return int64tovalue(dst, 0)
	}
	return int64tovalue(dst, i)
}

// convSingleUintFromSingleString: uint = string
func convSingleUintFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return uint64tovalue(dst, 0)
	}
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return uint64tovalue(dst, 0)
	}
	return uint64tovalue(dst, i)
}

// convSingleFloatFromSingleString: float = string
func convSingleFloatFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return float64tovalue(dst, 0)
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64tovalue(dst, 0)
	}
	return float64tovalue(dst, f)
}

// convSingleStringFromSingleString: string = string
func convSingleStringFromSingleString(dst, src reflect.Value) error {
	return stringtovalue(dst, valuetostring(src))
}

// convMultiBoolFromSingleString: []bool = string
func convMultiBoolFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return nil
	}
	ss := strings.Split(s, ",")
	return loopHelper(dst, reflect.ValueOf(ss), convSingleBoolFromSingleString)
}

// convMultiIntFromSingleString: []int = string
func convMultiIntFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return nil
	}
	ss := strings.Split(s, ",")
	return loopHelper(dst, reflect.ValueOf(ss), convSingleIntFromSingleString)
}

// convMultiUintFromSingleString: []uint = string
func convMultiUintFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return nil
	}
	ss := strings.Split(s, ",")
	return loopHelper(dst, reflect.ValueOf(ss), convSingleUintFromSingleString)
}

// convMultiFloatFromSingleString: []float = string
func convMultiFloatFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return nil
	}
	ss := strings.Split(s, ",")
	return loopHelper(dst, reflect.ValueOf(ss), convSingleFloatFromSingleString)
}

// convMultiStringFromSingleString: []string = string
func convMultiStringFromSingleString(dst, src reflect.Value) error {
	s := valuetostring(src)
	if len(s) < 1 {
		return nil
	}
	ss := strings.Split(s, ",")
	return loopHelper(dst, reflect.ValueOf(ss), convSingleStringFromSingleString)
}

// stringtovalue will set a Value (or childvalue to pointers) to string, allocating pointers where possible
func stringtovalue(v reflect.Value, x string) error {
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
	v.SetString(x)
	return nil
}

// valuetostring will get a sting from Value, resolving pointers along the way
func valuetostring(v reflect.Value) string {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	return v.String()
}
