package nettools

import (
	"strings"
)

// MacStripper takes a string and removes
// typical seperators used in MAC address
// formats.
func MacStripper(mac string) string {
	seperators := []rune{':', '-', '.'}
	var res strings.Builder
	for _, c := range mac {
		match := false
		for _, s := range seperators {
			if c == s {
				match = true
			}
		}
		if !match {
			res.WriteRune(c)
		}
	}
	return res.String()
}
