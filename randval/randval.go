package randval

import (
	"divsperf/randval/parse"
)

func Register(ger parse.Rvgenerator) {
	if _, ok := parse.Rvgenerators[ger.Name()]; !ok {
		parse.Rvgenerators[ger.Name()] = ger
	}
}