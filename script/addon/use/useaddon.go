package use

import (
	crvf_parse "divsperf/randval/parse"
	"divsperf/script"
	scri_parse "divsperf/script/parse"
	"sync"
)

type UseAddon struct {

}

func init() {
	useaddon := UseAddon{}
	script.Register(&useaddon)
}

func (*UseAddon) CanReturn() bool {
	return true
}

func (*UseAddon) Action_and_Return(wg *sync.WaitGroup, sb *scri_parse.SquareBrackets) (*[]rune, error) {
	if (*sb.Tokens)[0].Ctype != scri_parse.STR {
		return nil, scri_parse.UnmatchError_parameter_type{
			sb.Name,
			1,
			scri_parse.STR,
			(*sb.Tokens)[0].Ctype,
		}
	}
	tplname := string(*(*sb.Tokens)[0].Content)
	tpl := crvf_parse.Templates[tplname]
	return tpl.Generate()
}

func (*UseAddon) Action(wg *sync.WaitGroup, sb *scri_parse.SquareBrackets) error {
	return nil
}

func (*UseAddon) Name() string {
	return "UseAddon"
}