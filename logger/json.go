package logger

import (
	json2 "encoding/json"
	"fmt"
)

type json struct {
	app string
}

func NewJSON(a string) Logger {
	return &json{app: a}
}

func (j json) Info(key string, data map[string]interface{}) {
	data["app"] = j.app
	data["key"] = key
	val, err := json2.Marshal(data)
	if err != nil {
		fmt.Println("Deu problema")
	}

	fmt.Println(string(val))
}

func (j json) Error(key string, data map[string]interface{}) {
	fmt.Println("Deu ruim")
}
