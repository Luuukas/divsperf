package binary

import (
	"divsperf/randval"
	"divsperf/randval/parse"
)

type Binary struct {

}

func init() {
	binary := Binary{}
	randval.Register(&binary)
}

type MultipleError struct {

}

func (MultipleError) Error() string {
	return "error: not an integer multiple of 8.\n"
}

func (*Binary) Generate(sb parse.SquareBrackets) (*[]rune, error) {
	if len(*sb.Tokens) != 1 {
		return nil, parse.UnmatchError_number_of_parameters{
			sb.Name,
			1,
			len(*sb.Tokens),
		}
	}
	if (*sb.Tokens)[0].Ctype != parse.SBR {
		return nil, parse.UnmatchError_parameter_type{
			sb.Name,
			1,
			parse.SBR,
			(*sb.Tokens)[0].Ctype,
		}
	}
	if len(*(*sb.Tokens)[0].Content)%8 != 0 {
		return nil, MultipleError{}
	}
	var brs []rune
	for i:=0;i<len(*(*sb.Tokens)[0].Content);i+=8 {
		var r rune = 0
		for j:=0;j<8;j++ {
			r = r*2+(*(*sb.Tokens)[0].Content)[i+j]-48
		}
		brs = append(brs, r)
	}
	return &brs, nil
}

func (*Binary) Name() string {
	return "binary"
}