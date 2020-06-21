package script

import "divsperf/script/parse"

func Register(addon parse.Addon) {
	if _, ok := parse.Addons[addon.Name()]; !ok {
		parse.Addons[addon.Name()] = addon
	}
}