package main

import (
	"fmt"
	"strconv"
)

func modulo4WithZeroToFour(value byte) int {
	mod := (value - '0') % 4
	if mod == 0 {
		mod = 4
	}
	return int(mod)
}

func processFileContent(content string) string {
	if len(content) < 8 {
		return ""
	}

	output := "5"
	for i := 1; i <= 7; i++ {
		output += strconv.Itoa(modulo4WithZeroToFour(content[i]))
	}
	output += strconv.Itoa(modulo4WithZeroToFour(content[7]))
	output += strconv.Itoa(modulo4WithZeroToFour(content[6]))
	output += strconv.Itoa(modulo4WithZeroToFour(content[5]))
	output += strconv.Itoa(modulo4WithZeroToFour(content[4]))

	return output
}

func main() {
	content := "896987769"
	
	fmt.Println(processFileContent(content))
}
