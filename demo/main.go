package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	a := []map[string]interface{}{
		map[string]interface{}{"name": "xmzhang"},
		map[string]interface{}{"age": 24},
	}

	marshal, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(marshal))
	}
}
