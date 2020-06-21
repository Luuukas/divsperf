package tools

import (
	"divsperf/script/parse"
)

func RstoInt(rs []rune) (int, error) {
	res := 0
	for _, c := range rs {
		res = res * 10 + int(c)
	}
	return res, nil
}

func TKtoInt(tk parse.Token, name string, idx int) (int, error) {
	if tk.Ctype != parse.INT {
		return 0, parse.UnmatchError_parameter_type{
			name,
			idx+1,
			parse.INT,
			tk.Ctype,
		}
	}
	return RstoInt(*tk.Content)
}

func TkContent(tk parse.Token) (*[]rune, error) {
	switch tk.Ctype {
	case parse.SSB:
		fallthrough
	case parse.INT:
		fallthrough
	case parse.FLO:
		fallthrough
	case parse.STR:
		return tk.Content, nil
	default:
		data, err := tk.Sb.LetActionandReturn()
		if err!= nil {
			return nil, err
		}
		return data, nil
	}
}