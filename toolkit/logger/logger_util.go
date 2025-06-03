package logger

import (
	"encoding/json"
	"fmt"
)

// ParseJSON will transform struct data as json string.
func ParseJSON(data interface{}) string {
	JSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}

// ParsePrettyJSON will transform struct data as json indent string.
func ParsePrettyJSON(data interface{}) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}
