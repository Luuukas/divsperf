package parallel

// need to run divsperf\script\conner\udp\test_udp_echo.exe first

import (
	_ "divsperf/randval/create/rand"
	parse_crvf "divsperf/randval/parse"
	_ "divsperf/script/addon/use"
	_ "divsperf/script/conner/udp"
	parse_scri "divsperf/script/parse"
	"sync"
	"testing"
)

const (
	test_crvf = "test_parallel.crvf"
	test_scri = "test_parallel.scri"
)

func TestParallelAddon_Action(t *testing.T) {
	parser_crvf := parse_crvf.Parser{}
	err := parser_crvf.New(test_crvf)
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

	parser_scri := parse_scri.Parser{}
	err = parser_scri.New(test_scri)
	if err != nil {
		t.Error(err)
		return
	}
	err = parser_scri.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	parser_scri.RegisterScript()

	scrip := parse_scri.Scripts["testparalleladdonscri"]

	wg := &sync.WaitGroup{}
	err = scrip.Braces[0].Sbs[0].LetAction(wg)
	if err != nil {
		t.Error(err)
		return
	}
	err = scrip.Braces[0].Sbs[1].LetAction(wg)
	if err != nil {
		t.Error(err)
		return
	}
	wg.Wait()
}