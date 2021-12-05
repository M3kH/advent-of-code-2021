package challenge

import (
	"aoc/utils"
	"strconv"
	"strings"
)

var move = map[string]func(position int, modifier int) int{
	"down": func(position int, modifier int) int {
		return position + modifier
	},
	"forward": func(position int, modifier int) int {
		return position + modifier
	},
	"up": func(position int, modifier int) int {
		return position - modifier
	},
}

func Challenge0201(data string) string {
	var x, y int

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
			x = move["forward"](x, atSpeed)
			continue
		}

		y = move[moveTo](y, atSpeed)
	}

	return strconv.Itoa(x * y)
}
