// parser_test_export.go is needed to avoid circular references

package parse_test

import (
	_ "divsperf/randval/create/rand"
	"divsperf/randval/parse"
	"testing"
)

func TestParser_New(t *testing.T) {
	parse.TestParser_New(t)
}

func TestParser_skipBlank(t *testing.T) {
	parse.TestParser_skipBlank(t)
}

func TestParser_skipAnnotation(t *testing.T) {
	parse.TestParser_skipAnnotation(t)
}

func TestParser_getRange(t *testing.T) {
	parse.TestParser_getRange(t)
}

func TestParser_getWord(t *testing.T) {
	parse.TestParser_getWord(t)
}

func TestParser_parseSettings(t *testing.T) {
	parse.TestParser_parseSettings(t)
}

func TestParser_getToken(t *testing.T) {
	parse.TestParser_getToken(t)
}

func TestParser_parseSquareBrackets(t *testing.T) {
	parse.TestParser_parseSquareBrackets(t)
}

func TestParser_parseBrace(t *testing.T) {
	parse.TestParser_parseBrace(t)
}

func TestParser_Parse(t *testing.T) {
	parse.TestParser_Parse(t)
}