package service

import (
	"strings"
)

var tf = map[string]interface{}{
	"cmd": func(s string) string {
		// Put command in single quotes, otherwise special characters like dollar ($) sign will be interpreted.
		return `'` + strings.Replace(s, `'`, `'"'"'`, -1) + `'`
	},
	"cmdSystemD": func(s string) string {
		s = strings.Replace(s, `%`, `%%`, -1)
		s = `"` + strings.Replace(s, `"`, `\"`, -1) + `"`
		return s
	},
	"cmdEscape": func(s string) string {
		return strings.Replace(s, " ", `\x20`, -1)
	},
	"envKey": func(env string) string {
		return strings.Split(env, "=")[0]
	},
	"envValue": func(env string) string {
		return strings.Join(strings.Split(env, "=")[1:], "=")
	},
}
