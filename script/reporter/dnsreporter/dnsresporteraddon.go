package dnsreporter

import (
	"divsperf/script"
	"divsperf/script/parse"
	"sync"
)

type DnsResporterAddon struct {

}

func init() {
	dnsresporteraddon := DnsResporterAddon{}
	script.Register(&dnsresporteraddon)
}

func (*DnsResporterAddon) CanReturn() bool {
	return false
}

func (*DnsResporterAddon) Action_and_Return(wg *sync.WaitGroup, sb *parse.SquareBrackets) (*[]rune, error) {
	return nil, nil
}

func (*DnsResporterAddon) Action(wg *sync.WaitGroup, sb *parse.SquareBrackets) error {
	if (*sb.Tokens)[0].Ctype != parse.STR {
		return parse.UnmatchError_parameter_type{
			sb.Name,
			1,
			parse.STR,
			(*sb.Tokens)[0].Ctype,
		}
	}
	analyzername := string(*(*sb.Tokens)[0].Content)
	if (*sb.Tokens)[1].Ctype != parse.STR {
		return parse.UnmatchError_parameter_type{
			sb.Name,
			1,
			parse.STR,
			(*sb.Tokens)[0].Ctype,
		}
	}
	filename := string(*(*sb.Tokens)[1].Content)
	dnsreporter := DnsReporter{}
	return dnsreporter.rp(analyzername, filename)
}

func (*DnsResporterAddon) Name() string {
	return "DnsResporterAddon"
}