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
	for t:=0;t<pt;t++ {
		go func() {
			defer wg.Done()
			for _, tk := range (*(sb.Tokens))[1:] {
				fmt.Println(tk.Sb.Name)
				if err=tk.Sb.LetAction(wg); err != nil {
					fmt.Println("error in parallel: ",err)
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