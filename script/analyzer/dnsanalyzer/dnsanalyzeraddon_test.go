package dnsanalyzer

import (
	_ "divsperf/randval/create/binary"
	_ "divsperf/randval/create/rand"
	parse_crvf "divsperf/randval/parse"
	_ "divsperf/script/addon/parallel"
	_ "divsperf/script/addon/use"
	_ "divsperf/script/conner/udp"
	parse_scri "divsperf/script/parse"
	"fmt"
	"sync"
	"testing"
)

func TestDnsAnalyzerAddon_Action(t *testing.T) {
	parser_crvf := parse_crvf.Parser{}
	err := parser_crvf.New("test_dnsanalyzeraddon.crvf")
	if err != nil {
		t.Error(err)
		return
	}
	err = parser_crvf.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	parser_crvf.RegisterTemplate()
	rs, err := parse_crvf.Templates["dnsbinary"].Generate()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(rs)

	parser_scri := parse_scri.Parser{}
	err = parser_scri.New("test_dnsanalyzeraddon.scri")
	if err != nil {
		t.Error(err)
		return
	}
	err = parser_scri.Parse()
	if err != nil {
		return
	}
	parser_scri.RegisterScript()
	wg := &sync.WaitGroup{}
	err = parse_scri.Scripts["testdnsanalyzeraddonscri"].Braces[0].Sbs[0].LetAction(wg)
	if err != nil {
		return
	}
	err = parse_scri.Scripts["testdnsanalyzeraddonscri"].Braces[0].Sbs[1].LetAction(wg)
	if err != nil {
		return
	}
	wg.Wait()
}