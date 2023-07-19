package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// Encode JSON

func logJson(data interface{})  {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(data), ": ", string(bytes))
}

func TestEncode(t *testing.T)  {
	logJson("Rocket")
	logJson(1)
	logJson(true)
	logJson([]string{"Rocket", "Racoon"})
}
