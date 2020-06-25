package dnsanalyzer

import (
	"divsperf/script"
	"divsperf/script/parse"
	"sync"
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

func (*DnsAnalyzerAddon) Action_and_Return(wg *sync.WaitGroup, sb *parse.SquareBrackets) (*[]rune, error) {
	return nil, nil
}

func (daeraddon *DnsAnalyzerAddon) Action(wg *sync.WaitGroup, sb *parse.SquareBrackets) error {
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
		udpname := string(*(*sb.Tokens)[2].Content)
		var daer DnsAnalyzer
		err := daer.st(name, udpname)
		if err != nil {
			return err
		}
		daeraddon.DnsAnalyzers[name] = &daer
	case "be":
		if len(*sb.Tokens) != 2 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				2,
				len(*sb.Tokens),
			}
		}
		err := daeraddon.DnsAnalyzers[name].be()
		if err != nil {
			return err
		}
	case "ed":
		if len(*sb.Tokens) != 2 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				2,
				len(*sb.Tokens),
			}
		}
		err := daeraddon.DnsAnalyzers[name].ed()
		if err != nil {
			return err
		}
	case "ce":
		if len(*sb.Tokens) != 2 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				2,
				len(*sb.Tokens),
			}
		}
		err := daeraddon.DnsAnalyzers[name].ce()
		if err != nil {
			return err
		}
	default:
		return UnknownDirective{
			sb.Name,
			dir,
		}
	}
}

func (*DnsAnalyzerAddon) Name() string {
	return "analyzer"
}