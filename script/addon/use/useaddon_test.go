package use

import (
	_ "divsperf/randval/create/rand"
	crvf_parse "divsperf/randval/parse"
	scri_parse "divsperf/script/parse"
	"fmt"
	"sync"
	"testing"
)

func TestUseAddon_Action_and_Return(t *testing.T) {
	wg := &sync.WaitGroup{}

	crvf_parser := crvf_parse.Parser{}
	err := crvf_parser.New("test_useaddon.crvf")
	if err != nil {
		t.Error(err)
		return
	}
	err = crvf_parser.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	crvf_parser.RegisterTemplate()

	var tplname []rune
	tplname = []rune("testuseaddon")
	tk := scri_parse.Token{Ctype: scri_parse.STR, Content: &tplname}
	var tks []scri_parse.Token
	tks = append(tks, tk)
	sb := scri_parse.SquareBrackets{Name: "use", Tokens: &tks}

	useaddon := UseAddon{}

	rs, err := useaddon.Action_and_Return(wg, &sb)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(rs)
}