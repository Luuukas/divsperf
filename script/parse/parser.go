package parse

import (
	"io/ioutil"
	"strconv"
	"sync"
)

const (
	INT string = "int"
	STR string = "string"
	FLO string = "float"
	SSB string = "sub-square-brackets"
	SBR string = "sub-brace"
)

var (
	// Addon包括addon, reporter, analyzer, conner
	Addons = make(map[string] Addon)
	Scripts = make(map[string] *Script)
	Levels [10]*LevelCard
)

type LevelCard struct {
	SittingB *Brace
	Next *LevelCard
}

type Addon interface {
	Action(*sync.WaitGroup, *SquareBrackets) error
	CanReturn() bool
	Action_and_Return(*sync.WaitGroup, *SquareBrackets) (*[]rune, error)
	Name() string
}

type Token struct {
	Ctype string
	Content *[]rune
	Sb *SquareBrackets
}

type SquareBrackets struct {
	Name string
	Tokens *[]Token
}

func (sb *SquareBrackets) LetAction(wg *sync.WaitGroup) error {
	return Addons[sb.Name].Action(wg, sb)
}

func (sb *SquareBrackets) LetActionandReturn(wg *sync.WaitGroup) (*[]rune, error) {
	if !Addons[sb.Name].CanReturn() {
		return nil, TrytoUseAddonWithoutReturn{}
	}
	return Addons[sb.Name].Action_and_Return(wg, sb)
}

// 一个外层大括号的内容
// 一个大括号内容全部由中括号即可运行插件组成
type Brace struct {
	Sbs []SquareBrackets
	Rangelo int
	Rangehi int
	Level int
}

type Script struct {
	Name string
	Braces []Brace
}

type Parser struct {
	filename string
	scri []rune
	scri_len int
	cursor int
	runner Script
	outputfilename string
}

func (parser *Parser) New(filename string) (error) {
	parser.filename = filename
	crvfbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	parser.scri = []rune(string(crvfbytes))
	parser.scri_len = len(parser.scri)
	parser.cursor = 0
	return nil
}

func (parser *Parser) meetNewLine(cursor int) bool {
	if parser.scri[cursor] == rune(13) {
		if cursor+1 < parser.scri_len && parser.scri[cursor+1] == rune(10) {
			return true
		}
	}
	return false
}

func (parser *Parser) skipBlank() {
	var r rune
	for parser.cursor < parser.scri_len {
		r = parser.scri[parser.cursor]
		parser.cursor++
		if r == rune(32){
			continue
		} else if parser.meetNewLine(parser.cursor-1) {
			parser.cursor++
			continue
		}
		parser.cursor--
		break
	}
}

func (parser *Parser) skipAnnotation() {
	var r rune
	space := false
	for parser.cursor < parser.scri_len {
		r = parser.scri[parser.cursor]
		parser.cursor++
		if r == ' ' {
			space = true
		} else if r == '#' {
			if space {
				break
			} else {
				space = false
			}
		} else if parser.meetNewLine(parser.cursor-1) {
			parser.cursor++
			break
		} else {
			space = false
		}
	}
}

func (parser *Parser) getWord() *[]rune{
	parser.skipBlank()
	var rs []rune
	var r rune
	for parser.cursor < parser.scri_len {
		r = parser.scri[parser.cursor]
		parser.cursor++
		if r == ' ' {
			break
		} else if parser.meetNewLine(parser.cursor-1) {
			parser.cursor++
			break
		} else {
			rs = append(rs, r)
		}
	}
	return &rs
}

func (parser *Parser) parseSettings() error {
	for parser.cursor < parser.scri_len {
		rsp := parser.getWord()
		rs_str := string(*rsp)
		switch rs_str {
		case "#":
			parser.skipAnnotation()
		case "Script":
			arg := parser.getWord()
			parser.runner.Name = string(*arg)
			if parser.runner.Name == "" {
				return MissingParameter{}
			}
		case "}":
			return nil
		default:
			return UnknownKeyword{}
		}
	}
	return MissingRightBrace{}
}

func (parser *Parser) getRange() (lo int, hi int, err error) {
	parser.skipBlank()
	r := parser.scri[parser.cursor]
	parser.cursor++
	if r == '[' {
		parser.skipBlank()
		var rangelo []rune
		for parser.cursor < parser.scri_len {
			r = parser.scri[parser.cursor]
			parser.cursor++
			if r == ' ' {
				break
			} else if parser.meetNewLine(parser.cursor-1) {
				parser.cursor++
				break
			}
			rangelo = append(rangelo, r)
		}
		lo, err = strconv.Atoi(string(rangelo))
		if err != nil {
			return
		}
		parser.skipBlank()
		r = parser.scri[parser.cursor]
		if r == ']' {
			hi = lo
			err = nil
			parser.cursor++
			return
		}
		var rangehi []rune
		for parser.cursor < parser.scri_len {
			r = parser.scri[parser.cursor]
			parser.cursor++
			if r == ' ' {
				break
			} else if parser.meetNewLine(parser.cursor-1) {
				parser.cursor++
				break
			}
			rangehi = append(rangehi, r)
		}
		hi, err = strconv.Atoi(string(rangehi))
		if err != nil {
			return
		}
		parser.skipBlank()
		r := parser.scri[parser.cursor]
		parser.cursor++
		if r == ']' {
			return
		} else {
			err = BraceCntSBError{}
			return
		}
	} else {
		err = BraceCntSBError{}
		return
	}
}

func (parser *Parser) parseLevel() (int, error) {
	parser.skipBlank()
	var r rune
	if parser.cursor < parser.scri_len {
		r = parser.scri[parser.cursor]
		parser.cursor++
		if r != '[' {
			return -1, LevelGettingError{}
		}
	} else {
		return -1, LevelGettingError{}
	}
	Lv := parser.getWord()
	if string(*Lv) == "" {
		return -1, LevelGettingError{}
	}
	level := parser.getWord()
	if string(*level) == "" {
		return -1, LevelGettingError{}
	}
	level_i, err := strconv.Atoi(string(*level))
	if err != nil {
		return -1, LevelGettingError{}
	}
	parser.skipBlank()
	if parser.cursor < parser.scri_len {
		r = parser.scri[parser.cursor]
		parser.cursor++
		if r != ']' {
			return -1, LevelGettingError{}
		}
	} else {
		return -1, LevelGettingError{}
	}
	return level_i, nil
}

func (parser *Parser) parseSquareBrackets() (*SquareBrackets, error) {
	var sb SquareBrackets
	sb.Name = string(*parser.getWord())
	sb.Tokens = new([]Token)
	for {
		token, err := parser.getToken()
		if err != nil {
			return nil, err
		}
		if token == nil {
			return &sb, nil
		}
		*sb.Tokens = append(*sb.Tokens, *token)
	}
}

func (parser *Parser) getToken() (*Token, error) {
	parser.skipBlank()
	var token Token
	token.Content = new([]rune)
	var err error
	var r rune
	r = parser.scri[parser.cursor]
	parser.cursor++
	if r == '{' {
		left_brace_cnt := 1
		token.Ctype = SBR
		for parser.cursor < parser.scri_len {
			r = parser.scri[parser.cursor]
			parser.cursor++
			switch r {
			case '}':
				if !(parser.cursor>1 && parser.scri[parser.cursor-2] == '\\') {
					left_brace_cnt--
				}
				if left_brace_cnt == 0 {
					return &token, nil
				} else {
					*token.Content = append(*token.Content, '}')
				}
			case '{':
				if !(parser.cursor>1 && parser.scri[parser.cursor-2] == '\\') {
					left_brace_cnt++
				}
				*token.Content = append(*token.Content, '{')
			default:
				*token.Content = append(*token.Content, r)
			}
		}
		return nil, MissingRightBrace{}
	} else if r == '[' {
		token.Ctype = SSB
		token.Sb, err = parser.parseSquareBrackets()
		if err != nil {
			return nil, err
		}
	} else if r == ']' {
		return nil, nil
	} else {
		parser.cursor--
		isFirst := true
		isInt := false
		isFlo := -1
		for parser.cursor < parser.scri_len {
			r = parser.scri[parser.cursor]
			parser.cursor++
			*token.Content = append(*token.Content, r)
			if r == ' ' {
				*token.Content = (*token.Content)[:len(*token.Content)-1]
				break
			}  else if parser.meetNewLine(parser.cursor-1) {    // '\'跟个回车，会把下一行拼接到后面，而后忽略掉'\'
				if parser.scri[parser.cursor-2] == '\\' {
					*token.Content = (*token.Content)[:len(*token.Content)-2]
					parser.cursor++
				} else {
					*token.Content = (*token.Content)[:len(*token.Content)-1]
					parser.cursor++
					break
				}
			} else if r == '.' {
				if isInt {
					isInt = false
					isFlo = 0
				} else {
					isFlo = -1
				}
			} else if '0' <= r && r <= '9' {
				if isFirst {
					isInt = true
					isFirst = false
				}
				if isFlo == 0 {
					isFlo = 1
				}
			} else {
				isInt = false
				isFlo = -1
			}
		}
		if isInt {
			token.Ctype = INT
		} else if isFlo == 1 {
			token.Ctype = FLO
		}else {
			token.Ctype = STR
		}
	}
	return &token, nil
}

func (parser *Parser) parseBrace() (*Brace, error) {
	var brace Brace
	var r rune
	var err error
	for parser.cursor < parser.scri_len {
		parser.skipBlank()
		r = parser.scri[parser.cursor]
		parser.cursor++
		if r == '[' {
			sb, err := parser.parseSquareBrackets()
			if err != nil {
				return nil, err
			}
			brace.Sbs = append(brace.Sbs, *sb)
		} else if r == '}' {
			brace.Rangelo, brace.Rangehi, err = parser.getRange()
			if err != nil {
				return nil, err
			}
			brace.Level, err = parser.parseLevel()
			if err != nil {
				return nil, err
			}
			return &brace, nil
		} else {
			return nil, BraceParseError{}
		}
	}
	return nil, MissingRightBrace{}
}

func (parser *Parser) Parse() error {
	Settings_cnt := 0
	for parser.cursor < parser.scri_len {
		kw := parser.getWord()
		kw_str := string(*kw)
		switch kw_str {
		case "#":
			parser.skipAnnotation()
		case "Settings{":
			Settings_cnt++
			if Settings_cnt>1 {
				return TooMuchSettings{}
			}
			err := parser.parseSettings()
			if err != nil {
				return err
			}
		case "{":
			bracep, err := parser.parseBrace()
			if err != nil {
				return err
			}
			LvC := LevelCard{bracep, Levels[bracep.Level]}
			Levels[bracep.Level] = &LvC
			parser.runner.Braces = append(parser.runner.Braces, *bracep)
		default:
			return FLParseError{}
		}
	}
	return nil
}

func (parser *Parser) RegisterScript() {
	if _, ok := Scripts[parser.runner.Name]; !ok {
		Scripts[parser.runner.Name] = &parser.runner
	}
}