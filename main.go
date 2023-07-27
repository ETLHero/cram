package cram

/*
Struct to map?
Map to struct?
*/

import (
	"reflect"
	"errors"
	"fmt"
)

type ConversionFunction func(reflect.Value, reflect.Value) error

type ConversionLookup struct{
	DM bool        // Destination is a slice or array
	D reflect.Kind // Destination kind
	SM bool        // Source is a slice or array
	S reflect.Kind // Source kind
}

var (
	ErrUnsupportedConversion = errors.New("unsupported conversion")
	ErrPointerRequired = errors.New("pointer needed, can't change it otherwise")
)

var Conversions = map[ConversionLookup]ConversionFunction{
	{DM:false,D:reflect.Bool,SM:false,S:reflect.Bool}:    convSingleBoolFromSingleBool,
	{DM:false,D:reflect.Bool,SM:false,S:reflect.Int64}:   convSingleBoolFromSingleInt,
	{DM:false,D:reflect.Bool,SM:false,S:reflect.Uint64}:  convSingleBoolFromSingleUint,
	{DM:false,D:reflect.Bool,SM:false,S:reflect.Float64}: convSingleBoolFromSingleFloat,
	{DM:false,D:reflect.Bool,SM:false,S:reflect.String}:  convSingleBoolFromSingleString,
	{DM:false,D:reflect.Bool,SM:true,S:reflect.Bool}:     convSingleBoolFromMultiBool,
	{DM:false,D:reflect.Bool,SM:true,S:reflect.Int64}:    convSingleBoolFromMultiInt,
	{DM:false,D:reflect.Bool,SM:true,S:reflect.Uint64}:   convSingleBoolFromMultiUint,
	{DM:false,D:reflect.Bool,SM:true,S:reflect.Float64}:  convSingleBoolFromMultiFloat,
	{DM:false,D:reflect.Bool,SM:true,S:reflect.String}:   convSingleBoolFromMultiString,

	{DM:false,D:reflect.Int64,SM:false,S:reflect.Bool}:    convSingleIntFromSingleBool,
	{DM:false,D:reflect.Int64,SM:false,S:reflect.Int64}:   convSingleIntFromSingleInt,
	{DM:false,D:reflect.Int64,SM:false,S:reflect.Uint64}:  convSingleIntFromSingleUint,
	{DM:false,D:reflect.Int64,SM:false,S:reflect.Float64}: convSingleIntFromSingleFloat,
	{DM:false,D:reflect.Int64,SM:false,S:reflect.String}:  convSingleIntFromSingleString,
	{DM:false,D:reflect.Int64,SM:true,S:reflect.Bool}:     convSingleIntFromMultiBool,
	{DM:false,D:reflect.Int64,SM:true,S:reflect.Int64}:    convSingleIntFromMultiInt,
	{DM:false,D:reflect.Int64,SM:true,S:reflect.Uint64}:   convSingleIntFromMultiUint,
	{DM:false,D:reflect.Int64,SM:true,S:reflect.Float64}:  convSingleIntFromMultiFloat,
	{DM:false,D:reflect.Int64,SM:true,S:reflect.String}:   convSingleIntFromMultiString,

	{DM:false,D:reflect.Uint64,SM:false,S:reflect.Bool}:    convSingleUintFromSingleBool,
	{DM:false,D:reflect.Uint64,SM:false,S:reflect.Int64}:   convSingleUintFromSingleInt,
	{DM:false,D:reflect.Uint64,SM:false,S:reflect.Uint}:    convSingleUintFromSingleUint,
	{DM:false,D:reflect.Uint64,SM:false,S:reflect.Float64}: convSingleUintFromSingleFloat,
	{DM:false,D:reflect.Uint64,SM:false,S:reflect.String}:  convSingleUintFromSingleString,
	{DM:false,D:reflect.Uint64,SM:true,S:reflect.Bool}:     convSingleUintFromMultiBool,
	{DM:false,D:reflect.Uint64,SM:true,S:reflect.Int64}:    convSingleUintFromMultiInt,
	{DM:false,D:reflect.Uint64,SM:true,S:reflect.Uint64}:   convSingleUintFromMultiUint,
	{DM:false,D:reflect.Uint64,SM:true,S:reflect.Float64}:  convSingleUintFromMultiFloat,
	{DM:false,D:reflect.Uint64,SM:true,S:reflect.String}:   convSingleUintFromMultiString,

	{DM:false,D:reflect.Float64,SM:false,S:reflect.Bool}:    convSingleFloatFromSingleBool,
	{DM:false,D:reflect.Float64,SM:false,S:reflect.Int64}:   convSingleFloatFromSingleInt,
	{DM:false,D:reflect.Float64,SM:false,S:reflect.Uint64}:  convSingleFloatFromSingleUint,
	{DM:false,D:reflect.Float64,SM:false,S:reflect.Float64}: convSingleFloatFromSingleFloat,
	{DM:false,D:reflect.Float64,SM:false,S:reflect.String}:  convSingleFloatFromSingleString,
	{DM:false,D:reflect.Float64,SM:true,S:reflect.Bool}:     convSingleFloatFromMultiBool,
	{DM:false,D:reflect.Float64,SM:true,S:reflect.Int64}:    convSingleFloatFromMultiInt,
	{DM:false,D:reflect.Float64,SM:true,S:reflect.Uint64}:   convSingleFloatFromMultiUint,
	{DM:false,D:reflect.Float64,SM:true,S:reflect.Float64}:  convSingleFloatFromMultiFloat,
	{DM:false,D:reflect.Float64,SM:true,S:reflect.String}:   convSingleFloatFromMultiString,

	{DM:false,D:reflect.String,SM:false,S:reflect.Bool}:    convSingleStringFromSingleBool,
	{DM:false,D:reflect.String,SM:false,S:reflect.Int64}:   convSingleStringFromSingleInt,
	{DM:false,D:reflect.String,SM:false,S:reflect.Uint64}:  convSingleStringFromSingleUint,
	{DM:false,D:reflect.String,SM:false,S:reflect.Float64}: convSingleStringFromSingleFloat,
	{DM:false,D:reflect.String,SM:false,S:reflect.String}:  convSingleStringFromSingleString,
	{DM:false,D:reflect.String,SM:true,S:reflect.Bool}:     convSingleStringFromMultiBool,
	{DM:false,D:reflect.String,SM:true,S:reflect.Int64}:    convSingleStringFromMultiInt,
	{DM:false,D:reflect.String,SM:true,S:reflect.Uint64}:   convSingleStringFromMultiUint,
	{DM:false,D:reflect.String,SM:true,S:reflect.Float64}:  convSingleStringFromMultiFloat,
	{DM:false,D:reflect.String,SM:true,S:reflect.String}:   convSingleStringFromMultiString,

	{DM:true,D:reflect.Bool,SM:false,S:reflect.Bool}:    convMultiBoolFromSingleBool,
	{DM:true,D:reflect.Bool,SM:false,S:reflect.Int64}:   convMultiBoolFromSingleInt,
	{DM:true,D:reflect.Bool,SM:false,S:reflect.Uint64}:  convMultiBoolFromSingleUint,
	{DM:true,D:reflect.Bool,SM:false,S:reflect.Float64}: convMultiBoolFromSingleFloat,
	{DM:true,D:reflect.Bool,SM:false,S:reflect.String}:  convMultiBoolFromSingleString,
	{DM:true,D:reflect.Bool,SM:true,S:reflect.Bool}:     convMultiBoolFromMultiBool,
	{DM:true,D:reflect.Bool,SM:true,S:reflect.Int64}:    convMultiBoolFromMultiInt,
	{DM:true,D:reflect.Bool,SM:true,S:reflect.Uint64}:   convMultiBoolFromMultiUint,
	{DM:true,D:reflect.Bool,SM:true,S:reflect.Float64}:  convMultiBoolFromMultiFloat,
	{DM:true,D:reflect.Bool,SM:true,S:reflect.String}:   convMultiBoolFromMultiString,

	{DM:true,D:reflect.Int64,SM:false,S:reflect.Bool}:    convMultiIntFromSingleBool,
	{DM:true,D:reflect.Int64,SM:false,S:reflect.Int64}:   convMultiIntFromSingleInt,
	{DM:true,D:reflect.Int64,SM:false,S:reflect.Uint64}:  convMultiIntFromSingleUint,
	{DM:true,D:reflect.Int64,SM:false,S:reflect.Float64}: convMultiIntFromSingleFloat,
	{DM:true,D:reflect.Int64,SM:false,S:reflect.String}:  convMultiIntFromSingleString,
	{DM:true,D:reflect.Int64,SM:true,S:reflect.Bool}:     convMultiIntFromMultiBool,
	{DM:true,D:reflect.Int64,SM:true,S:reflect.Int64}:    convMultiIntFromMultiInt,
	{DM:true,D:reflect.Int64,SM:true,S:reflect.Uint64}:   convMultiIntFromMultiUint,
	{DM:true,D:reflect.Int64,SM:true,S:reflect.Float64}:  convMultiIntFromMultiFloat,
	{DM:true,D:reflect.Int64,SM:true,S:reflect.String}:   convMultiIntFromMultiString,

	{DM:true,D:reflect.Uint64,SM:false,S:reflect.Bool}:    convMultiUintFromSingleBool,
	{DM:true,D:reflect.Uint64,SM:false,S:reflect.Int64}:   convMultiUintFromSingleInt,
	{DM:true,D:reflect.Uint64,SM:false,S:reflect.Uint}:    convMultiUintFromSingleUint,
	{DM:true,D:reflect.Uint64,SM:false,S:reflect.Float64}: convMultiUintFromSingleFloat,
	{DM:true,D:reflect.Uint64,SM:false,S:reflect.String}:  convMultiUintFromSingleString,
	{DM:true,D:reflect.Uint64,SM:true,S:reflect.Bool}:     convMultiUintFromMultiBool,
	{DM:true,D:reflect.Uint64,SM:true,S:reflect.Int64}:    convMultiUintFromMultiInt,
	{DM:true,D:reflect.Uint64,SM:true,S:reflect.Uint64}:   convMultiUintFromMultiUint,
	{DM:true,D:reflect.Uint64,SM:true,S:reflect.Float64}:  convMultiUintFromMultiFloat,
	{DM:true,D:reflect.Uint64,SM:true,S:reflect.String}:   convMultiUintFromMultiString,

	{DM:true,D:reflect.Float64,SM:false,S:reflect.Bool}:    convMultiFloatFromSingleBool,
	{DM:true,D:reflect.Float64,SM:false,S:reflect.Int64}:   convMultiFloatFromSingleInt,
	{DM:true,D:reflect.Float64,SM:false,S:reflect.Uint64}:  convMultiFloatFromSingleUint,
	{DM:true,D:reflect.Float64,SM:false,S:reflect.Float64}: convMultiFloatFromSingleFloat,
	{DM:true,D:reflect.Float64,SM:false,S:reflect.String}:  convMultiFloatFromSingleString,
	{DM:true,D:reflect.Float64,SM:true,S:reflect.Bool}:     convMultiFloatFromMultiBool,
	{DM:true,D:reflect.Float64,SM:true,S:reflect.Int64}:    convMultiFloatFromMultiInt,
	{DM:true,D:reflect.Float64,SM:true,S:reflect.Uint64}:   convMultiFloatFromMultiUint,
	{DM:true,D:reflect.Float64,SM:true,S:reflect.Float64}:  convMultiFloatFromMultiFloat,
	{DM:true,D:reflect.Float64,SM:true,S:reflect.String}:   convMultiFloatFromMultiString,

	{DM:true,D:reflect.String,SM:false,S:reflect.Bool}:    convMultiStringFromSingleBool,
	{DM:true,D:reflect.String,SM:false,S:reflect.Int64}:   convMultiStringFromSingleInt,
	{DM:true,D:reflect.String,SM:false,S:reflect.Uint64}:  convMultiStringFromSingleUint,
	{DM:true,D:reflect.String,SM:false,S:reflect.Float64}: convMultiStringFromSingleFloat,
	{DM:true,D:reflect.String,SM:false,S:reflect.String}:  convMultiStringFromSingleString,
	{DM:true,D:reflect.String,SM:true,S:reflect.Bool}:     convMultiStringFromMultiBool,
	{DM:true,D:reflect.String,SM:true,S:reflect.Int64}:    convMultiStringFromMultiInt,
	{DM:true,D:reflect.String,SM:true,S:reflect.Uint64}:   convMultiStringFromMultiUint,
	{DM:true,D:reflect.String,SM:true,S:reflect.Float64}:  convMultiStringFromMultiFloat,
	{DM:true,D:reflect.String,SM:true,S:reflect.String}:   convMultiStringFromMultiString,
}

// Into is the primary function for this library. Cram source into destination.
func Into(destination, source any) error {
	dst := reflect.ValueOf(destination)
	src := reflect.ValueOf(source)
	if dst.Kind() != reflect.Pointer {
		return ErrPointerRequired
	}
	dm, d := Resolve(dst.Type())
	sm, s := Resolve(src.Type())
	composit := ConversionLookup{DM:dm,D:d,SM:sm,S:s}
	f, ok := Conversions[composit]
	if !ok {
		return fmt.Errorf("%w: %s", ErrUnsupportedConversion, composit)
	}
	return f(dst, src)
}

// Resolve will dectect slices and arrays then resolve pointers
// For instance:
//   Resolve(*[]*bool) -> true, reflect.Bool
//   Resolve(***float32) -> false, reflect.Float64
// Multipe slices does not work:
//   Resolve([][]string) -> true, reflect.Slice
func Resolve(x reflect.Type) (multi bool, k reflect.Kind) {
	for x.Kind() == reflect.Pointer {
		x = x.Elem()
	}

	k = x.Kind()
	multi = k == reflect.Array || k == reflect.Slice
	if !multi {
		k = Refine(k)
		return
	}

	x = x.Elem()
	for x.Kind() == reflect.Pointer {
		x = x.Elem()
	}

	k = Refine(x.Kind())
	return
}

// Refine will dumb down specific types that the reflect package requires you to use
func Refine(focus reflect.Kind) reflect.Kind {
	switch focus {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return reflect.Int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return reflect.Uint64
	case reflect.Float32:
		return reflect.Float64
	}
	return focus
}

func (cl ConversionLookup) String() string {
	n := map[bool]string{false:"single",true:"multi"}
	return fmt.Sprintf("%s %s to %s %s", n[cl.SM], cl.S, n[cl.DM], cl.D)
}
