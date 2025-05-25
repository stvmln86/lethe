// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"strings"
)

// Env returns the value of an existing environment variable.
func Env(name string) (string, error) {
	data, ok := os.LookupEnv(name)
	data = strings.TrimSpace(data)
	switch {
	case !ok:
		return "", fmt.Errorf("environment variable %q is not set", name)
	case data == "":
		return "", fmt.Errorf("environment variable %q is blank", name)
	default:
		return data, nil
	}
}

// Parse returns a parsed argument map from a parameter slice and argument slice.
// Parameters containing a colon will use the text after the colon as a default.
func Parse(paras, elems []string) (map[string]string, error) {
	var pairs = make(map[string]string)
	for i, para := range paras {
		name, dflt, ok := strings.Cut(para, ":")

		switch {
		case i >= len(elems) && ok:
			pairs[name] = dflt
		case i >= len(elems) && !ok:
			return nil, fmt.Errorf("argument %q was not provided", name)
		default:
			pairs[name] = elems[i]
		}
	}

	return pairs, nil
}

// Split returns a lowercase command name and argument slice from an argument slice.
func Split(elems []string) (string, []string) {
	switch len(elems) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(elems[0]), nil
	default:
		return strings.ToLower(elems[0]), elems[1:]
	}
}
