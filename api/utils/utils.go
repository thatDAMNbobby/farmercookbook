package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func PrintDebugJSON(label string, v interface{}) {
	j, _ := json.MarshalIndent(v, "", "  ")
	log.Debug(label, string(j))
}
