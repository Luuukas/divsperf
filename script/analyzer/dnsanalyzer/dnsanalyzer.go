package dnsanalyzer

import (
	"divsperf/script/conner/udp"
	"divsperf/script/parse"
	"github.com/miekg/dns"
	"math"
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
	RTT_min float64
	RTT_max float64
	RTT_tot float64
	Rcode_cnt map[int] int
}

func (daer *DnsAnalyzer) analyze(ssr *udp.SuccessSR) error {
	var RTT float64 = float64(ssr.Recv_t[0].UnixNano()-ssr.Sent_t.UnixNano())/1000000000
	if daer.RTT_max==-1 {
		daer.RTT_max = RTT
	}else {
		daer.RTT_max = math.Max(daer.RTT_max, RTT)
	}
	if daer.RTT_min==-1 {
		daer.RTT_min = RTT
	}else {
		daer.RTT_min = math.Max(daer.RTT_min, RTT)
	}
	daer.RTT_tot += RTT
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
	daer.done = make(chan struct{})
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
	close(daer.done)
	return nil
}

func (daer *DnsAnalyzer) ce() error {
	*daer.Srsucccntp = 0
	*daer.Srcntp = 0
	daer.RTT_max = -1
	daer.RTT_min = -1
	daer.RTT_tot = 0
	daer.Rcode_cnt = make(map[int] int)
	return nil
}