package parse

import (
	"divsperf/randval/parse"
	"fmt"
	"testing"
)

const (
	Settings_scri = "test_scri/Settings.scri"
	level_scri = "test_scri/level.scri"
	brace_scri = "test_scri/brace.scri"
	parse_scri = "test_scri/parse.scri"
)

func TestParser_parseSettings(t *testing.T) {
	parser := Parser{}
	err := parser.New(Settings_scri)
	if err != nil {
		t.Error(err)
		return
	}
	err = parser.parseSettings()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parser.runner.Name)
}

func TestParser_parserLevel(t *testing.T) {
	parser := Parser{}
	err := parser.New(level_scri)
	if err != nil {
		t.Error(err)
		return
	}
	Lv, err := parser.parseLevel()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(Lv)
	Lv, err = parser.parseLevel()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(Lv)
}

func TestParser_parseBrace(t *testing.T) {
	parser := Parser{}
	err := parser.New(brace_scri)
	if err != nil {
		t.Error(err)
		return
	}
	b, err := parser.parseBrace()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(b.Rangelo, b.Rangehi)
	fmt.Println(b.Level)
	for _, sb := range b.Sbs {
		fmt.Println(sb.Name)
		for _, tk := range *sb.Tokens {
			fmt.Print(tk.Ctype, tk.Content)
		}
		fmt.Println()
	}
}

func TestParser_Parse(t *testing.T) {
	parser := Parser{}
	err := parser.New(parse_scri)
	if err != nil {
		t.Error(err)
		return
	}
	err = parser.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	for _, b := range parser.runner.Braces {
		fmt.Println(b.Rangelo, b.Rangehi)
		fmt.Println(b.Level)
		for _, sb := range b.Sbs {
			fmt.Println(sb.Name)
			for _, tk := range *sb.Tokens {
				if tk.Ctype == parse.SSB {
					ssb := tk.Sb
					fmt.Println(ssb.Name)
					for _, ttk := range *ssb.Tokens {
						fmt.Print(ttk.Ctype, ttk.Content)
					}
				} else {
					fmt.Print(tk.Ctype, tk.Content)
				}
			}
			fmt.Println()
		}
	}
}