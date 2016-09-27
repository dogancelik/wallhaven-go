package search

import (
	"strings"
)

var availRes = [17]string{
	"1024x768",
	"1280x800",
	"1366x768",
	"1280x960",
	"1440x900",
	"1600x900",
	"1280x1024",
	"1600x1200",
	"1680x1050",
	"1920x1080",
	"1920x1200",
	"2560x1440",
	"2560x1600",
	"3840x1080",
	"5760x1080",
	"3840x2160",
	"5120x2880",
}

func getAbove(res string) []string {
	res = strings.Replace(res, "+", "", 1)
	index := -1
	for k, v := range availRes {
		if res == v {
			index = k
		}
	}

	if index == -1 {
		return []string{}
	} else {
		return availRes[index:]
	}
}

func sanitize(str string) string {
	return strings.Replace(strings.TrimSpace(str), "*", "x", 1)
}

func ParseResolutions(str string) []string {
	parsed := []string{}

	split := strings.Split(str, ",")
	for _, v := range split {
		v = sanitize(v)

		if strings.Contains(v, "+") == true {
			parsed = append(parsed, getAbove(v)...)
		} else {
			parsed = append(parsed, v)
		}
	}

	return parsed
}
