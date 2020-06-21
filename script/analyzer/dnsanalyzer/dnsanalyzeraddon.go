package dnsanalyzer

import (
	"divsperf/script"
	"divsperf/script/parse"
)

type DnsAnalyzerAddon struct {
	DnsAnalyzers map[string] *DnsAnalyzer
}

func init() {
	daeraddon := DnsAnalyzerAddon{}
	script.Register(&daeraddon)
}

func (*DnsAnalyzerAddon) CanReturn() bool {
	return false
}

func (*DnsAnalyzerAddon) Action_and_Return(sb *parse.SquareBrackets) (*[]rune, error) {
	return nil, nil
}

func (daeraddon *DnsAnalyzerAddon) Action(sb *parse.SquareBrackets) error {
	if (*sb.Tokens)[0].Ctype != parse.STR {
		return parse.UnmatchError_parameter_type{
			sb.Name,
			1,
			parse.STR,
			(*sb.Tokens)[0].Ctype,
		}
	}
	dir := string(*(*sb.Tokens)[0].Content)
	if (*sb.Tokens)[1].Ctype != parse.STR {
		return parse.UnmatchError_parameter_type{
			sb.Name,
			2,
			parse.STR,
			(*sb.Tokens)[1].Ctype,
		}
	}
	name := string(*(*sb.Tokens)[1].Content)
	switch dir {
	case "st":
		if len(*sb.Tokens) != 3 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				3,
				len(*sb.Tokens),
			}
		}
	case "be":
		if len(*sb.Tokens) != 2 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				2,
				len(*sb.Tokens),
			}
		}
	case "ed":
		if len(*sb.Tokens) != 2 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				2,
				len(*sb.Tokens),
			}
		}
	case "ce":
		if len(*sb.Tokens) != 2 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				2,
				len(*sb.Tokens),
			}
		}
	default:
		return UnknownDirective{
			sb.Name,
			dir,
		}
	}
}

func (*DnsAnalyzerAddon) Name() string {
	return "DnsAnalyzerAddon"
}