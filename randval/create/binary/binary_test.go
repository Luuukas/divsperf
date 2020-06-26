package binary

import (
	"divsperf/randval/parse"
	"fmt"
	"testing"
)

func TestBinary_Generate(t *testing.T) {
	content := []rune("11011011")
	tk := parse.Token{
		Ctype: parse.SBR,
		Content: &content,
	}
	tks := []parse.Token{tk}
	sb := parse.SquareBrackets{
		"binary",
		&tks,
	}

	ger := &Binary{}

	rs, err := ger.Generate(sb)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(rs)
}