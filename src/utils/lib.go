package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringToNumberConvertion(str []string) (values []int) {
	values = make([]int, 0, len(str))
	for _, raw := range str {
		v, err := strconv.Atoi(raw)
		if err != nil {
			continue
		}
		values = append(values, v)
	}
	return
}

func GetStdinString() string {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return string(output)
}

func HasStdinInput() bool {
	info, err := os.Stdin.Stat()

	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice == 0 && info.Size() > 0
}
