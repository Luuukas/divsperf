package dnsreporter

import (
	"divsperf/script/analyzer/dnsanalyzer"
	"divsperf/script/parse"
	"fmt"
)

type DnsReporter struct {

}

func (*DnsReporter) rp(analyzername string, filename string) error {
	analyzer := parse.Addons["DnsAnalyzerAddon"].(*dnsanalyzer.DnsAnalyzerAddon).DnsAnalyzers[analyzername]
	fmt.Printf("Queries sent:         %d queries\n", *analyzer.Srcntp)
	fmt.Printf("Queries completed:    %d queries\n", *analyzer.Srsucccntp)
	fmt.Printf("Queries lost:         %d queries\n", *analyzer.Srcntp-*analyzer.Srsucccntp)
	fmt.Println()
	fmt.Printf("RTT max:              %f sec\n", analyzer.RTT_max)
	fmt.Printf("RTT min:              %f sec\n", analyzer.RTT_min)
	fmt.Printf("RTT average:          %f sec\n", analyzer.RTT_tot/float64(*analyzer.Srcntp))
	fmt.Println()
	fmt.Printf("Percentage completed:  %f%%\n", float64(*analyzer.Srsucccntp)/float64(*analyzer.Srcntp))
	fmt.Printf("Percentage lost:       %f%%\n", 1.0-float64(*analyzer.Srsucccntp)/float64(*analyzer.Srcntp))
	fmt.Println()
	fmt.Printf("Started at:           %s\n", analyzer.Begin_t.Format("2006-01-02 15:04:05"))
	fmt.Printf("Finished at:          %s\n", analyzer.End_t.Format("2006-01-02 15:04:05"))
	fmt.Println()
	for key, rcnt := range analyzer.Rcode_cnt {
		fmt.Printf("Rcode - %b             %f%%\n", key, float64(rcnt)/float64(*analyzer.Srcntp))
	}
	return nil
}