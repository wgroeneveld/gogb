package util

import "reflect"

func StrToFields(interf interface{}, fieldNames ...string) []reflect.Value {
	r := reflect.Indirect(reflect.ValueOf(interf))
	fields := make([]reflect.Value, len(fieldNames))

	for i, name := range fieldNames {
		fields[i] = r.FieldByName(name)
	}

	return fields
}

