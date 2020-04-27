package es

import "reflect"

func typeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

func Int(i int) *int {
	return &i
}
