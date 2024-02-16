package helper

import (
	"strings"
)

func RemoveTimezone(time string) string {
	t := strings.Replace(time, "T", "", -1)
	return t
}
