package translator

import (
	"regexp"
	"strings"
)

var rule1 = regexp.MustCompile("^[^aeiouAEIOU][aeiouAEIOU]")
var rule2 = regexp.MustCompile("^[^aeiouAEIOU][^aeiouAEIOU]")
var rule3 = regexp.MustCompile("^[aeiouAEIOU]")
var punctuation = regexp.MustCompile(`[\.,?!;:\(\)]$`)

func TranslatePig(translateMe string) string {
	var tokens = strings.Fields(translateMe)
	var fin = []string{}
	var shouldCaps = true

	for _, token := range tokens {

		token, quoteStart, quoteEnd := stripQuotes(token)

		token = strings.ToLower(token)
		var punc = ""
		var willCapsNext = false
		if punctuation.MatchString(token) {
			punc = token[len(token)-1:]
			token = token[:len(token)-1]
			if punc == "." || punc == "!" || punc == "?" {
				willCapsNext = true
			}
		}

		if rule1.MatchString(token) {
			var moveMe = string(token[0])
			var editMe = token[1:]
			if shouldCaps {
				editMe = capitalize(editMe)
				shouldCaps = false
			}
			fin = append(fin, replaceQuotes(editMe+moveMe+"ay"+punc, quoteStart, quoteEnd))
		}

		if rule2.MatchString(token) {
			var moveMe = string(token[0:2])
			var editMe = token[2:]
			if shouldCaps {
				editMe = capitalize(editMe)
				shouldCaps = false
			}
			fin = append(fin, replaceQuotes(editMe+moveMe+"ay"+punc, quoteStart, quoteEnd))
		}

		if rule3.MatchString(token) {
			if shouldCaps {
				token = capitalize(token)
				shouldCaps = false
			}
			fin = append(fin, replaceQuotes(token+"way"+punc, quoteStart, quoteEnd))
		}

		if willCapsNext {
			shouldCaps = true
		}
	}

	return strings.Join(fin, " ")
}

func capitalize(capsMe string) string {
	if len(capsMe) == 0 {
		return capsMe
	}
	var cap = capsMe[0:1]
	cap = strings.ToUpper(cap)
	capsMe = cap + capsMe[1:]
	return capsMe
}

func stripQuotes(token string) (string, bool, bool) {
	var quoteStart = false
	var quoteEnd = false

	if string(token[0]) == "'" || string(token[0]) == "\"" || string(token[0]) == "`" {
		quoteStart = true
		token = token[1:]
	}

	if string(token[len(token)-1]) == "'" || string(token[len(token)-1]) == "\"" || string(token[0]) == "`" {
		quoteEnd = true
		token = token[:len(token)-1]
	}

	return token, quoteStart, quoteEnd
}

func replaceQuotes(token string, start bool, end bool) string {
	if start {
		token = "\"" + token
	}

	if end {
		token += "\""
	}

	return token
}
