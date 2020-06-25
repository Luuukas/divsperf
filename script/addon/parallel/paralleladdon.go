package parallel

import (
	"divsperf/script"
	"divsperf/script/parse"
	"divsperf/script/tools"
	"sync"
)

type ParallelAddon struct {

}

func init() {
	paralleladdon := ParallelAddon{}
	script.Register(&paralleladdon)
}

func (*ParallelAddon) CanReturn() bool {
	return false
}

func (*ParallelAddon) Action_and_Return(wg *sync.WaitGroup, sb *parse.SquareBrackets) (*[]rune, error) {
	return nil, nil
}

func (*ParallelAddon) Action(wg *sync.WaitGroup, sb *parse.SquareBrackets) error {
	pt, err := tools.TKtoInt((*sb.Tokens)[0], sb.Name, 0)
	if err != nil {
		return err
	}
	for t:=0;t<pt;t++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for _, tk := range (*(sb.Tokens))[1:] {
				if err=tk.Sb.LetAction(wg); err != nil {
					return
				}
			}
		}()
	}
	return nil
}

func (*ParallelAddon) Name() string {
	return "parallel"
}