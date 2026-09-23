package translator

import (
	"regexp"
	"strings"
)

var thornStart = regexp.MustCompile(`^th(.+)`)
var thornEnd = regexp.MustCompile(`(.+)th$`)
var eth = regexp.MustCompile(`(.+)th(.+)`)
var ash = regexp.MustCompile(`(.*)[aA]([^mn].*)`)

func TranslateOldEnglish(translateMe string) string {
	var tokens = strings.Fields(translateMe)
	var fin = []string{}

	for _, token := range tokens {
		token = thornStart.ReplaceAllString(token, "Þ${1}")
		token = thornEnd.ReplaceAllString(token, "${1}þ")
		token = eth.ReplaceAllString(token, "${1}ð${2}")
		token = ash.ReplaceAllString(token, "${1}Æ${2}")
		token = strings.ReplaceAll(token, "w", "ƿ")
		token = strings.ReplaceAll(token, "W", "Ƿ")
		fin = append(fin, token)
	}

	return strings.Join(fin, " ")
}
