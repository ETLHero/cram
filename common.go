package cram

import (
	"errors"
	"reflect"
	"fmt"
)

var (
	ErrCannotSet = errors.New("cannot write to this variable, make sure you have the right pointers")
	ErrNoRoomInList = errors.New("array or slice doesn't have room for any element")
)

// Unpointer will resolve all pointers. A nil returns the nil
func Unpointer(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return v
		}
		v = reflect.Indirect(v)
	}
	return v
}

// loopHelper is the common code for converting slices and arrays
func loopHelper(dst, src reflect.Value, f ConversionFunction) error {
	//TODO: Consider warning when array lengths don't match
	subdst := Unpointer(dst)
	subsrc := Unpointer(src)
	dstlen := subdst.Len()
	srclen := subsrc.Len()

	if subdst.Kind() == reflect.Slice && dstlen < srclen {
		if !subdst.CanSet() {
			return fmt.Errorf("allocating new slice to source len: %w", ErrCannotSet)
		}
		subdst.Set(reflect.MakeSlice(subdst.Type(), srclen, srclen))
		dstlen = srclen
	}

	for i := 0; i < subdst.Len() && i < subsrc.Len(); i++ {
		f(subdst.Index(i), subsrc.Index(i))
	}

	return nil
}

// firstHelper get the first element in array or slice
func firstHelper(v reflect.Value) reflect.Value {
	sv := Unpointer(v)
	if sv.Len() < 1 {
		return reflect.Zero(sv.Type())
	}
	return sv.Index(0)
}

func singleListItem(dst reflect.Value, x reflect.Value) error {
	dst = Unpointer(dst)
	if dst.Kind() == reflect.Slice && (dst.IsNil() || dst.Len() < 1) {
		dst.Set(reflect.MakeSlice(dst.Type(), 1, 1))
	}
	if dst.Len() < 1 {
		return ErrNoRoomInList
	}
	index := dst.Index(0)
	index = Unpointer(index)
	x = Unpointer(x)
	if !index.CanSet() {
		return ErrCannotSet
	}
	index.Set(x)
	return nil
}

func singleListItemConv(dst, src reflect.Value, f ConversionFunction) error {
	dstsub := Unpointer(dst)
	imd := reflect.New(dstsub.Type().Elem())
	err := f(imd, src)
	if err != nil {
		return err
	}
	return singleListItem(dst, imd)
}


