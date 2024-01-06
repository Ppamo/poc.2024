package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	jsonFile, err := os.Open("res/e03.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	o := map[string]interface{}{}
	json.Unmarshal(byteValue, &o)
	fmt.Printf(">> Showing json info\n")
	getMapFropMap("/", o, 0)
}

func getMapFropMap(propName string, propValue any, level int) {
	var (
		oTypeName, subTypeName, key string
	)
	printNestedIbjectInfo(level, propName, propValue)
	oTypeName = getTypeName(propValue)
	level++
	if oTypeName == "map[string]interface {}" {
		for key, value := range propValue.(map[string]interface{}) {
			subTypeName = getTypeName(value)
			switch subTypeName {
			case "string", "float64", "int", "bool":
				printNestedIbjectInfo(level, key, value)
			default:
				getMapFropMap(key, value, level)
			}
		}
	}
	if oTypeName == "[]interface {}" {
		for index, value := range propValue.([]interface{}) {
			key = fmt.Sprintf("%d", index)
			subTypeName = getTypeName(value)
			switch subTypeName {
			case "string", "float64", "int", "bool":
				printNestedIbjectInfo(level, key, value)
			default:
				getMapFropMap(key, value, level)
			}
		}
	}
}

func getTypeName(o any) string {
	var (
		oTypeName string
		oType     reflect.Type
	)
	oType = reflect.TypeOf(o)
	oTypeName = oType.Name()
	if len(oTypeName) == 0 {
		oTypeName = oType.String()
	}
	return oTypeName
}

func printNestedIbjectInfo(level int, key string, value any) {
	printNestedMessage(level, fmt.Sprintf("%s: %s", key, getTypeName(value)))
}

func printNestedMessage(level int, message string) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("- %s\n", message)
}
