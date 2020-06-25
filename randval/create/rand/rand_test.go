package rand

import (
	"divsperf/randval/create/tools"
	"divsperf/randval/parse"
	"fmt"
	"testing"
)

func TestRand_Generate(t *testing.T) {
	randToken_1_rs := []rune{5+48}
	randToken_1 := parse.Token{
		parse.INT,
		&randToken_1_rs,
		nil,
	}
	randToken_2_rs := []rune{3+48,3+48}
	randToken_2 := parse.Token{
		parse.INT,
		&randToken_2_rs,
		nil,
	}
	var Tokens []parse.Token
	Tokens = append(Tokens, randToken_1)
	Tokens = append(Tokens, randToken_2)
	sb := parse.SquareBrackets{
		"rand",
		&Tokens,
	}
	var ger parse.Rvgenerator = &Rand{}
	rs, err := ger.Generate(sb)
	if err != nil {
		t.Errorf("Error: get.Generate() gives rise to error.")
	}
	rsint, err := tools.RstoInt(*rs)
	if err != nil {
		t.Errorf("Error: rstoInt() gives rise to error.")
	}
	fmt.Println(rsint)
	if !(5<=rsint&&rsint<=33) {
		t.Errorf("Error: get.Generate() return a int that out of range.")
	}
}

func TestRand_Generate2(t *testing.T) {
	ssb1_tk1_content := []rune("1")
	ssb1_tk1 := parse.Token{Ctype: parse.INT, Content: &ssb1_tk1_content}
	ssb1_tk2_content := []rune("2")
	ssb1_tk2 := parse.Token{Ctype: parse.INT, Content: &ssb1_tk2_content}
	var tks1 []parse.Token
	tks1 = append(tks1, ssb1_tk1)
	tks1 = append(tks1, ssb1_tk2)
	ssb1 := parse.SquareBrackets{Name: "rand", Tokens: &tks1}

	ssb2_tk1_content := []rune("5")
	ssb2_tk1 := parse.Token{Ctype: parse.INT, Content: &ssb2_tk1_content}
	ssb2_tk2_content := []rune("9")
	ssb2_tk2 := parse.Token{Ctype: parse.INT, Content: &ssb2_tk2_content}
	var tks2 []parse.Token
	tks2 = append(tks2, ssb2_tk1)
	tks2 = append(tks2, ssb2_tk2)
	ssb2 := parse.SquareBrackets{Name: "rand", Tokens: &tks2}

	tk1 := parse.Token{Ctype: parse.SSB, Sb: &ssb1}
	tk2 := parse.Token{Ctype: parse.SSB, Sb: &ssb2}

	var tks []parse.Token
	tks = append(tks, tk1)
	tks = append(tks, tk2)

	sb := parse.SquareBrackets{Name: "rand", Tokens: &tks}

	var ger parse.Rvgenerator = &Rand{}
	rs, err := ger.Generate(sb)
	if err != nil {
		t.Errorf("Error: get.Generate() gives rise to error.")
	}
	rsint, err := tools.RstoInt(*rs)
	if err != nil {
		t.Errorf("Error: rstoInt() gives rise to error.")
	}
	fmt.Println(rsint)
}