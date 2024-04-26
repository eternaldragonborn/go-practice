package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StringToSliceString(s string) string {
	s = strings.ReplaceAll(s, "[", "{")
	s = strings.ReplaceAll(s, "]", "}")

	depth, maxDepth := 0, 0

	for _, c := range s {
		if c == '{' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if c == '}' {
			depth--
		}
	}

	if strings.Contains(s, "\"") {
		s = "string" + s
	} else if strings.Contains(s, "null") {
		s = strings.ReplaceAll(s, "null", "nil")
		s = "any" + s
	} else if match, _ := regexp.MatchString(`\d*\.\d*`, s); match {
		s = "float64" + s
	} else {
		s = "int" + s
	}
	for i := 0; i < maxDepth; i++ {
		s = "[]" + s
	}

	return s
}

var arrayRegexp = regexp.MustCompile(`^\[.*\]$`)

func StringToSlice(s string) []any {
	result := []any{}

	elements := strings.Split(s[1:len(s)-1], ",")
	for _, ele := range elements {
		ele = strings.TrimSpace(ele)
		if arrayRegexp.MatchString(ele) {
			// is sub-array
			result = append(result, StringToSlice(ele))
		} else if n, err := strconv.Atoi(ele); err == nil {
			// is int
			result = append(result, n)
		} else {
			result = append(result, ele)
		}
	}

	return result
}

func main() {
	s := []string{}
	tmp := ""

	for {
		fmt.Scanln(&tmp)
		if tmp == "q" {
			break
		}
		s = append(s, tmp)
	}

	for _, x := range s {
		fmt.Println(StringToSliceString(x))
		// fmt.Println(x)
	}
}
