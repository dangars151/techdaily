package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type A struct{}

type B struct {
	Abc interface{} `json:"abc"`
}

func main() {
	dataMap := map[string]interface{}{"abc": []*A{{}, {}}}
	var data B
	dataByte, _ := json.Marshal(dataMap)
	json.Unmarshal(dataByte, &data)

	_, ok1 := dataMap["abc"].([]*A)
	fmt.Println("ok1:", ok1)

	_, ok2 := data.Abc.([]*A)
	fmt.Println("ok2:", ok2)

	fmt.Printf("Type of data.Abc: %T\n", data.Abc)

	if slice, ok := data.Abc.([]interface{}); ok {
		fmt.Println("Successfully converted to []interface{}")
		fmt.Printf("Length of slice: %d\n", len(slice))
		for i, v := range slice {
			fmt.Printf("Element %d: %T %+v\n", i, v, v)
		}

		sliceType := reflect.TypeOf([]*A{})
		sliceValue := reflect.MakeSlice(sliceType, len(slice), len(slice))

		for i := range slice {
			elem := reflect.New(reflect.TypeOf(A{})).Elem()
			sliceValue.Index(i).Set(elem.Addr())
		}

		abc2 := sliceValue.Interface().([]*A)
		fmt.Printf("Successfully created []*A with length: %d\n", len(abc2))
		for i, a := range abc2 {
			fmt.Printf("Element %d: %p\n", i, a)
		}

		originalSlice := dataMap["abc"].([]*A)
		fmt.Println("\nComparing pointers:")
		for i := range originalSlice {
			fmt.Printf("Original: %p, New: %p, Match: %v\n", originalSlice[i], abc2[i], originalSlice[i] == abc2[i])
		}
	} else {
		fmt.Println("Failed to convert to []interface{}")
	}
}
