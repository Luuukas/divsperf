package udp

import (
	"divsperf/script"
	"divsperf/script/parse"
	"divsperf/script/tools"
	"sync"
)

const (
	BUFSIZE = 1024
)

type UdpAddon struct {
	Udps map[string] *Udp
}

func init() {
	udpaddon := UdpAddon{}
	script.Register(&udpaddon)
}

func (*UdpAddon) CanReturn() bool {
	return false
}

func (*UdpAddon) Action_and_Return(wg *sync.WaitGroup, sb *parse.SquareBrackets) (*[]rune, error) {
	return nil, nil
}

func (udpap *UdpAddon) Action(wg *sync.WaitGroup, sb *parse.SquareBrackets) error{
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
		if len(*sb.Tokens) != 4 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				4,
				len(*sb.Tokens),
			}
		}
		if _, ok:=udpap.Udps[name]; ok {
			return DuplicatedUdp{sb.Name, name}
		}
		if (*sb.Tokens)[2].Ctype != parse.STR {
			return parse.UnmatchError_parameter_type{
				sb.Name,
				3,
				parse.STR,
				(*sb.Tokens)[2].Ctype,
			}
		}
		ip := string(*(*sb.Tokens)[2].Content)
		port, err := tools.TKtoInt((*sb.Tokens)[3],sb.Name,3)
		if err != nil {
			return err
		}
		udp := Udp{Name:name}
		err = udp.st(ip, port)
		if err != nil {
			return err
		}
		udpap.Udps[name] = &udp
		return nil
	case "sr":
		if len(*sb.Tokens) != 6 {
			return parse.UnmatchError_number_of_parameters{
				sb.Name,
				6,
				len(*sb.Tokens),
			}
		}
		var udpp *Udp
		var ok bool
		if udpp, ok =udpap.Udps[name]; !ok {
			return CannotFindUdp{sb.Name, name}
		}
		wtimeout, err := tools.TKtoInt((*sb.Tokens)[2],sb.Name,2)
		if err != nil {
			return err
		}
		data, err := tools.TkContent((*sb.Tokens)[3])
		if err != nil {
			return err
		}
		rtimeout, err := tools.TKtoInt((*sb.Tokens)[2],sb.Name,2)
		if err != nil {
			return err
		}
		rtimes, err := tools.TKtoInt((*sb.Tokens)[2],sb.Name,2)
		if err != nil {
			return err
		}
		err = udpp.sr(wtimeout, *data, rtimeout, rtimes)
		return err
	default:
		return UnknownDirective{
			sb.Name,
			dir,
		}
	}
}

func (*UdpAddon) Name() string {
	return "UdpAddon"
}