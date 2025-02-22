package utils

import (
	"encoding/json"
)

func PrintDebugJSON(label string, v interface{}) {
	j, _ := json.MarshalIndent(v, "", "  ")
	println(label, string(j))
}
