package lgr

import (
	"reflect"
	"log"
)

func LogMethods(target interface{}) {
	v := reflect.ValueOf(target).Elem()
	t := v.Type()

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		methodType := t.Method(i)
		wrappedMethod := func(in []reflect.Value) []reflect.Value {
			log.Printf("Entering method %s", methodType.Name)
			result := method.Call(in)
			return result
		}
		method.Set(reflect.MakeFunc(method.Type(), wrappedMethod))
	}
}