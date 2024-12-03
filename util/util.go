package util

import (
	"log"
	"reflect"
)

func GetID(model interface{}) uint {
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName("id")
	if !field.IsValid() {
		log.Println("error to read the field id")
		return 0
	}

	return uint(field.Uint())
}
