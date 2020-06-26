package parallel

import (
	"divsperf/script"
	"divsperf/script/parse"
	"divsperf/script/tools"
	"fmt"
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
	wg.Add(pt)
	wg_in := &sync.WaitGroup{}
	wg_in.Add(pt)
	for t:=0;t<pt;t++ {
		tks := (*(sb.Tokens))[1:]
		go func() {
			defer func() {
				wg.Done()
				wg_in.Done()
			}()
			for _, tk := range tks {
				if err=tk.Sb.LetAction(wg_in); err != nil {
					fmt.Println("error in parallel: ",err)
					return
				}
			}
		}()
	}
	wg_in.Wait()
	return nil
}

func (*ParallelAddon) Name() string {
	return "parallel"
}