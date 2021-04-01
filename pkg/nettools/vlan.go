package nettools

import (
	"strconv"
	"strings"
)

func VLANExpand(s string) string {
	strSlice := strings.Split(s, "-")
	var intSlice = []int{}
	var res strings.Builder

	for _, i := range strSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, j)
	}
	lastElem := intSlice[len(intSlice)-1]
	for i := intSlice[0]; i <= lastElem; i++ {
		res.WriteString(strconv.Itoa(i))
		if i != lastElem {
			res.WriteString(",")
		}
	}
	return res.String()
}
