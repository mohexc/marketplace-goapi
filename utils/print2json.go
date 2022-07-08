package utils

import (
	"encoding/json"
	"fmt"
)

func Print2json(data interface{}) {
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
