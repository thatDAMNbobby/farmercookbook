package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"math"
)

func PrintDebugJSON(label string, v interface{}) {
	j, _ := json.MarshalIndent(v, "", "  ")
	log.Debug(label, string(j))
}

func Min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
