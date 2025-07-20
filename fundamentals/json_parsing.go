package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type MyJson struct {
	Test  any    `json:"test"`
	Test3 string `json:"test3"`
}

func JSON_Parsing_Sample() {
	var jsonParsed MyJson
	err := json.Unmarshal([]byte(`{"test": { "test2": [1,2,3] }, "test3": "dummy string"}`), &jsonParsed)
	if err != nil {
		log.Fatal(err)
	}

	switch v := jsonParsed.Test.(type) {
	case map[string]any:
		fmt.Printf("Found Type: %v\n", v)
		field1, ok := v["test2"]
		if ok {
			switch v2 := field1.(type) {
			case []any:
				fmt.Print("I foudn a type: []any\n")
				for _, v2Element := range v2 {
					fmt.Printf("Type: %s\n", reflect.TypeOf(v2Element))
					switch v2ElementValue := v2Element.(type) {
					default:
						fmt.Printf("Nested V2 Element Type not found: %s\n", reflect.TypeOf(v2ElementValue))
					}
				}
			default:
				fmt.Printf("Nested Type not found: %s\n", reflect.TypeOf(v2))
			}
		}
	default:
		fmt.Printf("Type not found: %s\n", reflect.TypeOf(jsonParsed))
	}
}
