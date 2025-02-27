package internal

import (
	"regexp"
	"strings"
)

func StringJoin(s ...string) string {
	return strings.Join(s, " ")
}

// https://regex101.com/
var uriReg = regexp.MustCompile(`(^https?:\/\/(www\.))?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#()?&//=]*)$`)

func IsUri(uri string) bool {
	return uriReg.MatchString(uri)
}

func PickArgs(args []string, ptr int) string {
	if len(args) > ptr {
		return args[ptr]
	}
	panic("Invalid ptr")
}

func Uniq(input []string) []string {
	r := make([]string, 0, len(input))
	bucket := make(map[string]bool)
	for _, each := range input {
		if _, v := bucket[each]; !v {
			bucket[each] = true
			r = append(r, each)
		}
	}
	return r
}
