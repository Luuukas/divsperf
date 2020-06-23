package parse

import (
	"fmt"
	"testing"
)

const (
	test_crvf = "test_crvfs/example.crvf"
	blank_crvf = "test_crvfs/blank.crvf"
	annotation_crvf = "test_crvfs/annotation.crvf"
	range_crvf = "test_crvfs/range.crvf"
	getWord_crvf = "test_crvfs/getWord.crvf"
	Settings_crvf = "test_crvfs/Settings.crvf"
	brace_crvf = "test_crvfs/brace.crvf"
	squarebrackets_crvf = "test_crvfs/squarebrackets.crvf"
	token_crvf = "test_crvfs/token.crvf"
)

func TestParser_New(t *testing.T) {
	parser := Parser{}
	err := parser.New(test_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parser.filename)
	fmt.Println(parser.crvf)
}

func TestParser_skipBlank(t *testing.T) {
	parser := Parser{}
	err := parser.New(blank_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parser.crvf[parser.cursor:])
	parser.skipBlank()
	fmt.Println(parser.cursor)
	fmt.Println(parser.crvf[parser.cursor:])
}

func TestParser_skipAnnotation(t *testing.T) {
	parser := Parser{}
	err := parser.New(annotation_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parser.crvf[parser.cursor:])
	parser.skipAnnotation()
	fmt.Println(parser.cursor)
	fmt.Println(parser.crvf[parser.cursor:])

	parser.skipAnnotation()
	fmt.Println(parser.cursor)
	fmt.Println(parser.crvf[parser.cursor:])
	parser.skipBlank()

	parser.skipAnnotation()
	fmt.Println(parser.cursor)
	fmt.Println(parser.crvf[parser.cursor:])
}

func TestParser_getRange(t *testing.T) {
	parser := Parser{}
	err := parser.New(range_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parser.crvf[parser.cursor:])
	lo, hi, err := parser.getRange()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(lo, hi)
	fmt.Println(parser.crvf[parser.cursor:])
}

func TestParser_getWord(t *testing.T) {
	parser := Parser{}
	err := parser.New(getWord_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(*parser.getWord())
	fmt.Println(*parser.getWord())
	fmt.Println(*parser.getWord())
	fmt.Println(*parser.getWord())
	fmt.Println(*parser.getWord())
	fmt.Println(*parser.getWord())
	fmt.Println(*parser.getWord())
}

func TestParser_parseSettings(t *testing.T) {
	parser := Parser{}
	err := parser.New(Settings_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	err = parser.parseSettings()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parser.tpl.Name)
	fmt.Println(parser.outputfilename)
	fmt.Println(parser.tpl.Repeat)
}

func TestParser_getToken(t *testing.T) {
	parser := Parser{}
	err := parser.New(token_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	tk, err := parser.getToken()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(tk.Ctype)
	fmt.Println(string(*tk.Content))
}

func TestParser_parseSquareBrackets(t *testing.T) {
	parser := Parser{}
	err := parser.New(squarebrackets_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	sb, err := parser.parseSquareBrackets()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(sb.Name)
	fmt.Println((*sb.Tokens)[0].Ctype)
	fmt.Println(string(*(*sb.Tokens)[0].Content))
	fmt.Println((*sb.Tokens)[1].Ctype)
	fmt.Println(string(*(*sb.Tokens)[1].Content))
}

func TestParser_parseBrace(t *testing.T) {
	parser := Parser{}
	err := parser.New(brace_crvf)
	if err != nil {
		t.Error(err)
		return
	}
	bracep, err := parser.parseBrace()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(bracep.rangelo)
	fmt.Println(bracep.rangehi)
}