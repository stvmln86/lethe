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
		return "", fmt.Errorf("cannot read variable %q - does not exist", name)
	case data == "":
		return "", fmt.Errorf("cannot read variable %q - is empty", name)
	default:
		return data, nil
	}
}

// Parse returns an parameter:argument map from parameter and argument slices. If a
// parameter contains a colon, the text after the colon is used as a default value.
func Parse(paras, argus []string) (map[string]string, error) {
	var amap = make(map[string]string)
	for i, para := range paras {
		name, dflt, ok := strings.Cut(para, ":")
		switch {
		case i >= len(argus) && ok:
			amap[name] = dflt
		case i >= len(argus) && !ok:
			return nil, fmt.Errorf("cannot parse arguments - %q missing", name)
		default:
			amap[name] = argus[i]
		}
	}

	return amap, nil
}

// Split returns a command name and argument slice from an argument slice. If no
// command name is specified, Split returns the default "help".
func Split(argus []string) (string, []string) {
	switch len(argus) {
	case 0:
		return "help", nil
	case 1:
		return strings.ToLower(argus[0]), nil
	default:
		return strings.ToLower(argus[0]), argus[1:]
	}
}
