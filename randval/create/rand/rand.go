package rand

import (
	"divsperf/randval"
	"divsperf/randval/create/tools"
	"divsperf/randval/parse"
	mrand "math/rand"
	"strconv"
	"time"
)

type Rand struct {

}

func init() {
	rand := Rand{}
	randval.Register(&rand)
}

func (*Rand) Generate(sb parse.SquareBrackets) (*[]rune, error) {
	if len(*sb.Tokens) != 2 {
		return nil, parse.UnmatchError_number_of_parameters{
			sb.Name,
			2,
			len(*sb.Tokens),
		}
	}
	l, err := tools.TKtoInt((*sb.Tokens)[0], sb.Name, 0)
	if err != nil {
		return nil, err
	}
	h, err := tools.TKtoInt((*sb.Tokens)[1], sb.Name, 1)
	if err != nil {
		return nil, err
	}
	mrand.Seed(time.Now().Unix())
	r := l + mrand.Intn(h+1)
	r_str := strconv.Itoa(r)
	var res []rune
	for _, c := range r_str {
		res = append(res, c-'0')
	}
	return &res, nil
}

func (*Rand) Name() string {
	return "rand"
}