/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectDemo(t *testing.T) {
	var x float64 = 3.4
	var xV = reflect.ValueOf(x)
	fmt.Println(reflect.TypeOf(x))    //float64
	fmt.Println(xV.Kind())            //float64
	fmt.Println(xV.Interface().(int)) //还原成原值

	p := reflect.ValueOf(&x)
	fmt.Println(p.Kind())                        //ptr
	fmt.Println("type of p:", p.Type())          //*float64
	fmt.Println("settability of p:", p.CanSet()) //false
	v := p.Elem()
	v.SetFloat(1.6)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

type FamilyMember struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Parents []string
}

func (u FamilyMember) GetMPrint(name string, age int) {
	fmt.Println("GetMPrint->Name", name, "Print--age-->", age)
}

func TestMarshalTagDemo(t *testing.T) {
	var m = FamilyMember{}
	var mv = reflect.TypeOf(&m)
	var mvv = mv.Elem()
	for i := 0; i < mvv.NumField(); i++ {
		tagInfo := mvv.Field(i).Tag.Get("json")
		fmt.Println("tag:", tagInfo)
	}

	var mv1 = reflect.ValueOf(&m)
	var mvv1 = mv1.Type().Elem() //It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice
	for i := 0; i < mvv1.NumField(); i++ {
		tagInfo := mvv1.Field(i).Tag.Get("json")
		fmt.Println("tag:", tagInfo)
	}

	var mv2 = reflect.TypeOf(m)
	for i := 0; i < mv2.NumField(); i++ {
		tagInfo := mv2.Field(i).Tag.Get("json")
		fmt.Println("tag:", tagInfo)
	}

	var mv3 = reflect.ValueOf(m)
	var mvv3 = mv3.Type()
	for i := 0; i < mvv3.NumField(); i++ {
		tagInfo := mvv3.Field(i).Tag.Get("json")
		fmt.Println("tag:", tagInfo)
	}
}

func TestReflectMethodCall(t *testing.T) {
	var m = FamilyMember{}
	var oo = reflect.ValueOf(m)
	var me = oo.MethodByName("GetMPrint")
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	me.Call(args)
}

type user struct {
	UserId string `model:"pk" type:"string"`
	Name   string
	Lvl    int
}

func TestReflectStructSlice(t *testing.T) {
	var (
		model    *user
		modelSet []*user
		st       reflect.Type
		elem     reflect.Value
		slice    reflect.Value
	)
	st = reflect.TypeOf(model).Elem()       //反射后的类型依然是指针类型，这时Elem()指向的就是结构体类型
	elem = reflect.New(st).Elem()           //reflect.New(st)创建了一个结构体对象，New返回值是指向结构体指针的反射；Elem()是取结构体类型反射值
	model = elem.Addr().Interface().(*user) //结构体地址反射值
	elem.FieldByName("UserId").SetString("12345678")
	elem.FieldByName("Name").SetString("nickname")
	elem.FieldByName("Lvl").SetInt(1)
	t.Log("model", model)

	st = reflect.TypeOf(modelSet)
	t.Log("reflect.TypeOf", st.Kind().String())
	slice = reflect.ValueOf(&modelSet).Elem()
	slice.Set(reflect.MakeSlice(st, 0, 16))
	slice.SetLen(1)
	t.Log("slice.len", slice.Len())
	slice.Index(0).Set(elem.Addr())
	t.Log("slice[0].kind", slice.Index(0).Kind().String())
	slice.Index(0).Elem().FieldByName("Lvl").SetInt(2)
	t.Log("slice[0]", slice.Index(0).Interface())
}

func TestReflectStruct(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("nickname")
	sv.FieldByName("Lvl").SetInt(1)
	t.Log("model", model)
}

type S struct {
	Field1 int
	field2 string
}

func TestReflectDeepEqualDemo(t *testing.T) {
	Array1 := []string{"hello1", "hello2"}
	Array2 := []string{"hello1", "hello2"}
	fmt.Println(reflect.DeepEqual(Array1, Array2))

	s1 := S{Field1: 1, field2: "hello"}
	s2 := S{Field1: 1, field2: "hello"}
	fmt.Println(reflect.DeepEqual(s1, s2))
}
