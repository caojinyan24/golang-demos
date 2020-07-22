package basic

import (
	"fmt"
	"reflect"
)

func Reflect()  {
	var a="abc"
	a_type:=reflect.TypeOf(a)
	fmt.Println("a.type=",a_type)
	fmt.Println("a.type.name=",a_type.Name())
	a_value:=reflect.ValueOf(a)
	fmt.Println("a.value=",a_value)
	fmt.Println("a.value.name=",a_value.Interface())


}
