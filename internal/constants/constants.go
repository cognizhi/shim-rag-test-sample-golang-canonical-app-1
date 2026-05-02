package constants

import (
	"math"
	"sort"
	"strings"
)

var values = map[string]float64{
	"e":   math.E,
	"pi":  math.Pi,
	"tau": 2 * math.Pi,
}

func Lookup(name string) (float64, bool) {
	value, exists := values[strings.ToLower(name)]
	return value, exists
}

func Names() []string {
	names := make([]string, 0, len(values))
	for name := range values {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
