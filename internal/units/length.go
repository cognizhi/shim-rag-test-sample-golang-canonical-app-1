package units

import (
	"fmt"
	"sort"
	"strings"
)

var lengthUnitsInMeters = map[string]float64{
	"centimeter": 0.01,
	"cm":         0.01,
	"foot":       0.3048,
	"ft":         0.3048,
	"inch":       0.0254,
	"in":         0.0254,
	"kilometer":  1000,
	"km":         1000,
	"meter":      1,
	"m":          1,
	"mile":       1609.344,
	"millimeter": 0.001,
	"mm":         0.001,
	"yard":       0.9144,
	"yd":         0.9144,
}

func ConvertLength(value float64, fromUnit string, toUnit string) (float64, error) {
	fromMultiplier, exists := lengthUnitsInMeters[strings.ToLower(fromUnit)]
	if !exists {
		return 0, fmt.Errorf("unsupported source unit %q", fromUnit)
	}

	toMultiplier, exists := lengthUnitsInMeters[strings.ToLower(toUnit)]
	if !exists {
		return 0, fmt.Errorf("unsupported target unit %q", toUnit)
	}

	return value * fromMultiplier / toMultiplier, nil
}

func SupportedLengthUnits() []string {
	units := make([]string, 0, len(lengthUnitsInMeters))
	for unit := range lengthUnitsInMeters {
		units = append(units, unit)
	}
	sort.Strings(units)
	return units
}
