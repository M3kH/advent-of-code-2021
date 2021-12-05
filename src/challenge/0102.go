package challenge

import (
	"aoc/utils"
	"strconv"
	"strings"
)

func Challenge0102(data string) string {
	var previous, total int

	depths := utils.StringToNumberConvertion(strings.Split(data, "\n"))

	for index, depth := range depths {
		current := depth

		if index+1 < len(depths) {
			current = current + depths[index+1]
		}

		if index+2 < len(depths) {
			current = current + depths[index+2]
		}

		if index != 0 && current > previous {
			total += 1
		}

		previous = current
	}

	return strconv.Itoa(total)
}
