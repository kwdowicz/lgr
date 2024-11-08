package lgr

import (
	"fmt"
	"log"
	"reflect"
)

func LogMethods(target interface{}) {
	v := reflect.ValueOf(target).Elem()
	t := v.Type()
	log.Printf("In LogMethods: %v\n", v)

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		methodType := t.Method(i)
		log.Printf("method: %v\n", method)
		log.Printf("method type: %v\n", method)
		wrappedMethod := func(in []reflect.Value) []reflect.Value {
			log.Printf("Entering method %s", methodType.Name)
			fmt.Printf("(fmt)Entering method %s", methodType.Name)
			result := method.Call(in)
			return result
		}
		method.Set(reflect.MakeFunc(method.Type(), wrappedMethod))
	}
}
