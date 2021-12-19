package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "If we`d go again all the way from the start \nI would try to change"

	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
