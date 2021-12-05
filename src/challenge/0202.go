package challenge

import (
	"aoc/utils"
	"strconv"
	"strings"
)

func Challenge0202(data string) string {
	var x, y, aim int

	records := strings.Split(data, "\n")

	for _, record := range records {
		instructions := strings.Split(record, " ")

		if len(instructions) == 1 {
			continue
		}

		moveTo := instructions[0]
		atSpeed, err := strconv.Atoi(instructions[1])
		utils.Check(err)

		if moveTo == "forward" {
			x = x + atSpeed
			y = y + (aim * atSpeed)
			continue
		}

		if moveTo == "up" {
			aim = aim - atSpeed
			continue
		}

		if moveTo == "down" {
			aim = aim + atSpeed
			continue
		}
	}

	return strconv.Itoa(x * y)
}
