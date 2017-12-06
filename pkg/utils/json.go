package utils

import "encoding/json"

//MustJSON Marshal json
func MustJSON(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}
