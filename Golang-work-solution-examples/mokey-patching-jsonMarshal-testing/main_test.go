package main

import (
	"errors"
	"fmt"
	"mokey-patching-json-testing/helper"
	"testing"
)

var sb []byte = []byte{ 123, 34, 102, 105, 114, 115, 116, 78, 97, 109, 101, 34, 58, 34, 86, 105, 107, 114, 97, 109, 34, 44, 34, 108, 97, 115, 116, 78, 97, 109, 101, 34, 58, 34, 83, 105, 110, 103, 104, 34, 125 }

func TestMarshalWithError(t *testing.T) {
	
	oldMarshal := helper.JsonMarshalFunc
	defer func() { helper.JsonMarshalFunc = oldMarshal }()
	helper.JsonMarshalFunc = func(v interface{}) ([]byte, error) {
		fmt.Println("----------JsonMarshal Success----------")
		return nil, errors.New("Error while Marshalling")
	}

	p1 := Person {
		// FirstName: "Vikram",
		// LastName: "Singh",
	}

	s := MarshalData(p1)
	fmt.Println(s)
}

func TestMarshalWithoutError(t *testing.T) {
	
	oldMarshal := helper.JsonMarshalFunc
	defer func() { helper.JsonMarshalFunc = oldMarshal }()
	helper.JsonMarshalFunc = func(v interface{}) ([]byte, error) {
		fmt.Println("----------JsonMarshal Success----------")
		return sb, nil
	}

	p1 := Person {
		FirstName: "Vikram",
		LastName: "Singh",
	}

	s := MarshalData(p1)
	fmt.Println(s)
}

// func TestMain(t *testing.T) {
// 	main()
// } 