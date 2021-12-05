package challenge

import (
	"aoc/utils"
	"strconv"

	"strings"
)

func Challenge0101(data string) string {
	var previous, total int

	depths := utils.StringToNumberConvertion(strings.Split(data, "\n"))

	for index, depth := range depths {
		if index != 0 && depth > previous {
			total += 1
		}
		previous = depth
	}

	return strconv.Itoa(total)
}
