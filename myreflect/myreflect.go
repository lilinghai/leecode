package myreflect

import (
	"fmt"
	"reflect"
)

func Reflect1(){
	f:=3.14
	vf:=reflect.ValueOf(f)
	fmt.Println(reflect.TypeOf(f),reflect.ValueOf(f))
	fmt.Println(vf.Type(),vf.Kind(),vf.Float(),vf.Kind()==reflect.Float64,reflect.Slice)
	type myint int
	var a myint=10
	fmt.Println(reflect.ValueOf(a).Kind(),reflect.ValueOf(a).Interface())

	vfp:=reflect.ValueOf(&f)
	vfp.Elem().SetFloat(10)
	fmt.Println(vfp.Kind(),vfp.CanSet(),f)

	type T struct {
		A int
		B string
	}
	t:=T{23,"skidoo"}
	s:=reflect.ValueOf(&t).Elem()
	for i:=0;i<s.NumField();i++{
		fmt.Println(s.Field(i),s.Type().Field(i).Name)
	}
}









