package tools

import "divsperf/randval/parse"

func RstoInt(rs []rune) (int, error) {
	res := 0
	for _, c := range rs {
		res = res * 10 + (int(c) - 48)
	}
	return res, nil
}

func IsInt(rs []rune) bool {
	for _, c := range rs {
		if int(c)<48 || int(c)>57 {
			return false
		}
	}
	return true
}

func TKtoInt(tk parse.Token, name string, idx int) (int, error) {
	if tk.Ctype == parse.SSB {
		rs, err := tk.Sb.GetContent()
		if err != nil {
			return -1, err
		}
		if !IsInt(*rs) {
			return 0, parse.UnmatchError_parameter_type{
				name,
				idx+1,
				parse.INT,
				tk.Ctype+"(which won't return an integer)",
			}
		}
		return RstoInt(*rs)
	}
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
