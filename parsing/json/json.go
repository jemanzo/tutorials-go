package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// -------------------------------------------
// JSON Datatypes - (Marshal/Unmarshal)
// Docs: https://pkg.go.dev/encoding/json
// -------------------------------------------
// JSON objects    map[string]interface{}
// JSON arrays     []interface{}
// JSON numbers    float64
// JSON strings    string
// JSON booleans   bool
// JSON null       nil
// -------------------------------------------
// JSON Tags
// `json:"<fieldName>,<fieldOptionOrType>)"`
// -------------------------------------------
// `json:"name"`           Field appears in JSON as key "name"
// `json:"name,omitempty"` Field is omitted from the object if its value is empty
// `json:",omitempty"`     Field is omitted from the object if its value is empty
// `json:"-"` 		       Field is ignored (empty or not)
// `json:"-,"` 		       Field appears in JSON as key "-"
// -------------------------------------------

type Item struct {
	Text1 string `json:"text1"`
	Text2 string `json:"text2"`
	Num1  int16  `json:"num1"`
	Num2  int16  `json:"num2,string"`
}

type List struct {
	Name1 string `json:"name1"`
	Name2 string `json:"name2" myTag:"MyName2"`
	Prop1 string `json:"-"`
	Prop2 string
	Items []Item `json:"items"`
}

func main() {
	ReadingStructTags()
	// ShowFloatToIntParsingError()
	// UnmarshalRawObject()
	// UnmarshalRawArray()
}

func CreateList() List {
	return List{
		Name1: "myList",
		Prop1: "DoNotMarshalThisProperty",
		Prop2: "MarshalThisPropertyWithOriginalUppercaseName",
		Items: []Item{
			{"a1", "b1", 123, 321},
			{"a2", "b2", 234, 432},
			{"a3", "b3", 345, 543},
		},
	}
}

func ReadingStructTags() {
	list := CreateList()

	typ := reflect.TypeOf(list)
	f, _ := typ.FieldByName("Name2")
	list.Name2 = f.Tag.Get("myTag")

	listJson, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(listJson))
}

func ShowFloatToIntParsingError() {
	// Parse JSON string
	jsonStr := []byte(`{
		"name1":"myList",
		"name2":"MyName2",
		"Prop2":"MarshalThisPropertyWithOriginalUppercaseName",
		"items":[{
			"text1":true,
			"text2":"b",
			"num1":123.22,
			"num2":456
		}]
	}`)

	var list2 List
	err := json.Unmarshal(jsonStr, &list2)
	if err != nil {
		log.Println(err)
	}
	log.Println(list2)
}

func UnmarshalRawObject() {
	jsonStr := []byte(`{
		"name1":"myList",
		"name2":"MyName2",
		"Prop2":"MarshalThisPropertyWithOriginalUppercaseName",
		"items":[{
			"text1":true,
			"text2":"b",
			"num1":123.22,
			"num2":456
		}]
	}`)

	var obj map[string]interface{}

	err := json.Unmarshal(jsonStr, &obj)
	if err != nil {
		log.Println(err)
	}
	log.Println(obj)
	log.Printf("Type of %q %T", "object", obj)
	log.Printf("Type of %q %T", "items", obj["items"])
	firstItem := obj["items"].([]interface{})[0].(map[string]interface{})
	log.Printf("Type of %q %T %v", "items.text1", firstItem["text1"], firstItem["text1"])
}

func UnmarshalRawArray() {
	jsonStr := []byte(`[ 1, 2, {
		"items":[{
			"text1":true,
			"text2":"b",
			"num1":123.22,
			"num2":456
		}]
	}]`)

	var arr []interface{}

	err := json.Unmarshal(jsonStr, &arr)
	if err != nil {
		log.Println(err)
	}
	log.Println(arr)
	log.Printf("Type of %q %T", "array", arr)
	log.Printf("Type of %q %T %v", "array[0]", arr[0], arr[0])
	log.Printf("Type of %q %T %v", "array[1]", arr[1], arr[1])
	log.Printf("Type of %q %T %v", "array[2]", arr[2], arr[2])

	obj := arr[2].(map[string]interface{})
	log.Printf("Type of %q %T %v", "items", obj["items"], obj)
	firstItem := obj["items"].([]interface{})[0].(map[string]interface{})
	log.Printf("Type of %q %T %v", "items.text1", firstItem["text1"], firstItem["text1"])
}
