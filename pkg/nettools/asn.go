package nettools

import (
	"fmt"
	"strconv"
	"strings"
)

func ASNPlainToDot(asn string) (string, error) {
	i, _ := strconv.Atoi(asn)
	if (i < 0) || (i > 4294967295) {
		return "", fmt.Errorf("not a valid asn: %d", i)
	}
	var asnDot string
	switch {
	case i <= 65535:
		asnDot = fmt.Sprintf("0.%d", i)
	case i >= 65536:
		decimal := i / 65536
		higherOrder := int(decimal)
		lowerOrder := i - higherOrder*65536
		asnDot = fmt.Sprintf("%d.%d", higherOrder, lowerOrder)
	}
	return asnDot, nil
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
