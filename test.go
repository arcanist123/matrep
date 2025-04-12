package main

import (
	"fmt"
	"reflect"
)

func main2() {
	// Define fields for the dynamic struct, including one lowercase
	fields := []reflect.StructField{
		{
			Name: "PublicField",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "privateField",
			Type: reflect.TypeOf(0),
		},
	}

	// Create the dynamic struct type
	dynamicStructType := reflect.StructOf(fields)

	// Create an instance of the dynamic struct
	dynamicStructValue := reflect.New(dynamicStructType).Elem()

	// Get reflect.Value for the fields
	publicField := dynamicStructValue.FieldByName("PublicField")
	privateField := dynamicStructValue.FieldByName("privateField")

	fmt.Println("PublicField CanSet:", publicField.IsValid() && publicField.CanSet())
	fmt.Println("privateField CanSet:", privateField.IsValid() && privateField.CanSet())

	// Attempt to set the public field
	if publicField.IsValid() && publicField.CanSet() {
		publicField.SetString("Value set for PublicField")
	} else {
		fmt.Println("Cannot set PublicField")
	}

	// Attempt to set the private field
	if privateField.IsValid() && privateField.CanSet() {
		privateField.SetInt(123)
	} else {
		fmt.Println("Cannot set privateField")
	}

	fmt.Println("Dynamic Struct Value:", dynamicStructValue)
}
