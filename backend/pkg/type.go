package pkg

import (
	"fmt"
	"reflect"
	"time"
)

func IsEmpty(i interface{}) bool {
	ret := i == nil || &i == nil

	if !ret {
		switch i.(type) {
		case *string:
			value := i.(*string)
			return value == nil || *value == ""
		case *int:
			value := i.(*int)
			return value == nil || *value == 0
		case *float32:
			value := i.(*float32)
			return value == nil || (*value-0) < 0
		case *float64:
			value := i.(*float64)
			return value == nil || (*value-0) < 0
		case *time.Time:
			value := i.(*time.Time)
			return value == nil || *value == time.Time{}
		default:
			fmt.Println(reflect.TypeOf(i))
		}
	}

	return ret
}

func IsErrorType(value interface{}) bool {
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	return reflect.TypeOf(value).Implements(errorInterface)
}
