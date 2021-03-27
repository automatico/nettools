package nettools

import (
	"fmt"
	"strconv"
	"strings"
)

func ASNPlainToDot(asn string) string {
	i, _ := strconv.Atoi(asn)
	var asnDot string
	switch {
	case i == 65536:
		asnDot = "1.0"
	case i < 65536:
		asnDot = fmt.Sprintf("0.%d", i)
	case i > 65536:
		decimal := i / 65536
		higherOrder := int(decimal)
		lowerOrder := i - higherOrder*65536
		asnDot = fmt.Sprintf("%d.%d", higherOrder, lowerOrder)
	}
	return asnDot
}

func ASNDotToPlain(asn string) string {
	split := strings.Split(asn, ".")
	higherOrder, _ := strconv.Atoi(split[0])
	lowerOrder, _ := strconv.Atoi(split[1])
	var asnPlain string

	switch {
	case higherOrder == 0:
		asnPlain = fmt.Sprintf("%d", lowerOrder)
	case higherOrder >= 1:
		total := higherOrder*65536 + lowerOrder
		asnPlain = fmt.Sprintf("%d", total)
	}

	return asnPlain
}
