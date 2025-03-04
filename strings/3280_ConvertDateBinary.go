package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	parts := strings.Split(date, "-")

	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	return fmt.Sprintf("%b-%b-%b", year, month, day)
}
