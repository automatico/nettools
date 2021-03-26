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

// MacFormatter takes a mac and a desired
// mac format and returns the mac in the
// desired format
func MacFormatter(mac string, format string) string {
	stripped := MacStripper(mac)
	var formatted string

	switch format {
	case "raw":
		formatted = stripped
	case "eui":
		formatted = StringInserter(stripped, ":", 2)
	case "unix":
		formatted = StringInserter(stripped, ":", 2)
	case "dot":
		formatted = StringInserter(stripped, ".", 4)
	}
	return formatted
}

// StringInserter takes a source string and inserts
// a string (s) every N distance (d).
func StringInserter(source string, s string, d int) string {
	for i := d; i < len(source); i += d + 1 {
		source = source[:i] + s + source[i:]
	}
	return source
}
