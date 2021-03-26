package nettools

import (
	"fmt"
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

// StringInserter takes a source string and inserts
// a string (s) every N distance (d).
func StringInserter(source string, s string, d int) string {
	for i := d; i < len(source); i += d + 1 {
		source = source[:i] + s + source[i:]
	}
	return source
}

// MacFormatter takes a mac and a desired
// mac format and returns the mac in the
// desired format
// Valid Formats:
// raw : 111122223333
// eui = 11:11:22:22:33:33
// unix = 11-11-22-22-33-33
// dot = 1111.2222.3333
func MacFormatter(mac string, format string) (string, error) {
	stripped := MacStripper(mac)

	if len(stripped) != 12 {
		return "", fmt.Errorf("not a valid mac address: %s", mac)
	}

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
	return formatted, nil
}
