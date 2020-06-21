package rand

import (
	"divsperf/randval/create/tools"
	"divsperf/randval/parse"
	"testing"
)

func TestRand_Generate(t *testing.T) {
	var randToken_1_rs []rune
	l := 5
	randToken_1_rs = append(randToken_1_rs, rune(5))
	randToken_1 := parse.Token{
		parse.INT,
		&randToken_1_rs,
		nil,
	}
	var randToken_2_rs []rune
	h := 33
	randToken_2_rs = append(randToken_2_rs, rune(3))
	randToken_2_rs = append(randToken_2_rs, rune(3))
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
	if !(l<=rsint&&rsint<=h) {
		t.Errorf("Error: get.Generate() return a int that out of range.")
	}
}
