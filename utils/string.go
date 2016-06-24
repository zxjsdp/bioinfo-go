package utils

import (
	"strings"
	"strconv"
)

func GenerateSpaces(spaceNum int) string {
	return strings.Repeat(" ", spaceNum)
}

func IsStringIntType(stringToCheck string) bool {
	if _, err := strconv.Atoi(stringToCheck); err == nil {
		return true;
	}
	return false;
}
