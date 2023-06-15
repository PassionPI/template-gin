package json

import (
	"strings"
)

type object = map[string]interface{}

func findValue(json any, keys string) (interface{}, bool) {
	current, ok := json.(object)
	if ok {
		segment := strings.Split(keys, ".")
		last := len(segment) - 1

		for index, part := range segment {
			value, exists := current[part]
			if exists {
				if index == last {
					return value, true
				}
				if next, ok := value.(object); ok {
					current = next
					continue
				}
			}
			break
		}
	}
	return nil, false
}

// FindValue  returns the value of the given keys(eg: "a.b.c") & type T in the json object
func FindValue[T any](json any, keys string) (T, bool) {
	x, ok := findValue(json, keys)
	t, ok := x.(T)
	return t, ok
}
