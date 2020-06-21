package dnsanalyzer

import (
	"divsperf/script/conner/udp"
	"divsperf/script/parse"
	"github.com/miekg/dns"
	"time"
)

type DnsAnalyzer struct {
	Name string
	done chan struct{}
	dnsServer *udp.Udp
	Srcntp *int
	Srsucccntp *int
	Begin_t time.Time
	End_t time.Time
	RTT_min float32
	RTT_max float32
	Rcode_cnt map[int] int
}

func (daer *DnsAnalyzer) analyze(ssr *udp.SuccessSR) error {
	msg := new(dns.Msg)
	msg.Unpack(ssr.Datas[0])
	if _, ok := daer.Rcode_cnt[msg.Rcode]; !ok {
		daer.Rcode_cnt[msg.Rcode] = 0
	}
	daer.Rcode_cnt[msg.Rcode]++
	return nil
}

func (daer *DnsAnalyzer) st(name string, udpname string) error {
	if daer.dnsServer != nil {
		return DuplicatedTarget{
			"DnsAnalyzer",
			daer.dnsServer.Name,
			name,
		}
	}
	daer.dnsServer = parse.Addons["UdpAddon"].(*udp.UdpAddon).Udps[udpname]
	daer.Srcntp = &daer.dnsServer.Srcnt
	daer.Srsucccntp = &daer.dnsServer.Srsucccnt
	return nil
}

func (daer *DnsAnalyzer) be() error {
	daer.Begin_t = time.Now()
	go func() {
		select {
		case <-daer.done:
			return
		case ssr:=<-daer.dnsServer.Succs:
			daer.analyze(ssr)
		}
	}()
	return nil
}

func (daer *DnsAnalyzer) ed() error {
	daer.End_t = time.Now()
	daer.done<- struct{}{}
	return nil
}

func (daer *DnsAnalyzer) ce() error {

}