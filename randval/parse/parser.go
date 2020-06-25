package parse

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type UnmatchError_number_of_parameters struct {
	Key_name string
	Valid_num int
	Present_num int
}
func (err UnmatchError_number_of_parameters) Error() string {
	return fmt.Sprintf("%s - the number of parameters does not match: it should be %d, but get %d.\n", err.Key_name, err.Valid_num, err.Present_num)
}

type UnmatchError_parameter_type struct {
	Key_name string
	Idx int
	Valid_type string
	Present_type string
}
func (err UnmatchError_parameter_type) Error() string {
	return fmt.Sprintf("%s - the parameter type is incorrect: %d-th should be %s, but get %s\n", err.Key_name, err.Idx, err.Valid_type, err.Present_type)
}

const (
	INT string = "int"
	STR string = "string"
	FLO string = "float"
	SSB string = "sub-square-brackets"
	SBR string = "sub-brace"
)

type Token struct {
	Ctype string
	Content *[]rune
	Sb *SquareBrackets
}

type SquareBrackets struct {
	Name string
	Tokens *[]Token
}

type Containing interface {
	GetContent() (*[]rune, error)
}

// 换行
type Br struct {
	rangelo int
	rangehi int
}

func (br *Br) GetContent() ( *[]rune, error) {
	var res []rune
	rand.Seed(time.Now().Unix())
	cnt := br.rangelo + rand.Intn(br.rangehi - br.rangelo + 1)
	for i:=0; i<cnt; i++ {
		res = append(res, 13)
		res = append(res, 10)
	}
	return &res, nil
}

// 连续固定的文本内容
type ConstStr struct {
	str []rune
}

func (cs *ConstStr) GetContent() (*[]rune, error) {
	return &cs.str, nil
}

// 一个外层大括号的内容
// 一个大括号内容由 生成器、文本内容、换行符
type Brace struct {
	Cts []Containing
	rangelo int
	rangehi int
}

// 一个Template由多个Brace组成
type Template struct {
	Name string
	Braces []Brace
	Repeat int
}

func (tpl *Template) Generate() (*[]rune, error) {
	var res []rune
	for bt:=0;bt<tpl.Repeat;bt++ {
		for _, brace := range tpl.Braces {
			rand.Seed(time.Now().Unix())
			cnt := brace.rangelo + rand.Intn(brace.rangehi-brace.rangelo+1)
			for i := 0; i < cnt; i++ {
				for _, ct := range brace.Cts {
					ctn, err := ct.GetContent()
					if err != nil {
						return nil, err
					}
					res = append(res, *ctn...)
				}
			}
		}
	}
	return &res, nil
}

// todo: 如果生成的数据太多，会占用太多内存，直接边生成边保存在文件中
func (tpl *Template) output(filename string) error {
	return nil
}

func (sb *SquareBrackets) GetContent() (*[]rune, error) {
	ger := Rvgenerators[sb.Name]
	ren, err := ger.Generate(*sb)
	if err != nil {
		return nil, err
	}
	return ren, nil
}

type Rvgenerator interface {
	Generate(SquareBrackets) (*[]rune, error)
	Name() string
}

var (
	Rvgenerators = make(map[string] Rvgenerator)
	Templates = make(map[string] *Template)
)

type Parser struct {
	filename string
	crvf []rune
	crvf_len int
	cursor int
	tpl Template
	outputfilename string
}

func (parser *Parser) New(filename string) error {
	parser.filename = filename
	crvfbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	parser.crvf = []rune(string(crvfbytes))
	parser.crvf_len = len(parser.crvf)
	parser.cursor = 0
	return nil
}

func (parser *Parser) meetNewLine(cursor int) bool {
	if parser.crvf[cursor] == rune(13) {
		if cursor+1 < parser.crvf_len && parser.crvf[cursor+1] == rune(10) {
			return true
		}
	}
	return false
}

func (parser *Parser) skipBlank() {
	var r rune
	for parser.cursor < parser.crvf_len {
		r = parser.crvf[parser.cursor]
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

func (parser *Parser) getWord() *[]rune{
	parser.skipBlank()
	var rs []rune
	var r rune
	for parser.cursor < parser.crvf_len {
		r = parser.crvf[parser.cursor]
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

type UnknownKeyword struct {}

func (UnknownKeyword) Error() string {
	return "error: an undefined keyword was encountered.\n"
}

type MissingParameter struct {}

func (MissingParameter) Error() string {
	return "error: missing parameter in Settings.\n"
}

type MissingRightBrace struct {}

func (MissingRightBrace) Error() string {
	return "error: missing closing brace results in incomplete blocks.\n"
}

func (parser *Parser) parseSettings() error {
	for parser.cursor < parser.crvf_len {
		rsp := parser.getWord()
		rs_str := string(*rsp)
		switch rs_str {
		case "#":
			parser.skipAnnotation()
		case "Template":
			arg := parser.getWord()
			parser.tpl.Name = string(*arg)
			if parser.tpl.Name == "" {
				return MissingParameter{}
			}
		case "Outputfile":
			arg := parser.getWord()
			parser.outputfilename = string(*arg)
			if parser.outputfilename == "" {
				return MissingParameter{}
			}
		case "Repeat":
			arg := parser.getWord()
			var err error
			parser.tpl.Repeat, err = strconv.Atoi(string(*arg))
			if err != nil {
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

func (parser *Parser) skipAnnotation() {
	var r rune
	space := false
	for parser.cursor < parser.crvf_len {
		r = parser.crvf[parser.cursor]
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

type BraceCntSBError struct {}

func (BraceCntSBError) Error() string {
	return "error: there should be a [ x ] or [ x y ] following the brace.\n"
}

func (parser *Parser) getRange() (lo int, hi int, err error) {
	parser.skipBlank()
	r := parser.crvf[parser.cursor]
	parser.cursor++
	if r == '[' {
		parser.skipBlank()
		var rangelo []rune
		for parser.cursor < parser.crvf_len {
			r = parser.crvf[parser.cursor]
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
		r = parser.crvf[parser.cursor]
		if r == ']' {
			hi = lo
			err = nil
			parser.cursor++
			return
		}
		var rangehi []rune
		for parser.cursor < parser.crvf_len {
			r = parser.crvf[parser.cursor]
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
		r := parser.crvf[parser.cursor]
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

func (parser *Parser) parseBrace() (*Brace, error) {
	var brace Brace
	var r rune
	var err error
	for parser.cursor < parser.crvf_len {
		r = parser.crvf[parser.cursor]
		parser.cursor++
		if r == '[' {
			sb, err := parser.parseSquareBrackets()
			if err != nil {
				return nil, err
			}
			brace.Cts = append(brace.Cts, sb)
		} else if r == '}' {
			brace.rangelo, brace.rangehi, err = parser.getRange()
			if err != nil {
				return nil, err
			}
			return &brace, nil
		} else {
			var cstr ConstStr
			cstr.str = append(cstr.str, r)
			for parser.cursor < parser.crvf_len {
				r = parser.crvf[parser.cursor]
				parser.cursor++
				if r == '[' {
					if parser.crvf[parser.cursor-2] == '\\' {
						cstr.str = cstr.str[:len(cstr.str)-1]
						cstr.str = append(cstr.str, '[')
					} else {
						parser.cursor--
						break
					}
				} else if parser.meetNewLine(parser.cursor-1) {
					if parser.crvf[parser.cursor-2] == '\\' {
						cstr.str = cstr.str[:len(cstr.str)-1]
					} else {
						cstr.str = append(cstr.str, '\n')
					}
					parser.cursor++
				} else if r == '}' {
					if parser.crvf[parser.cursor-2] == '\\' {
						cstr.str = cstr.str[:len(cstr.str)-1]
						cstr.str = append(cstr.str, '}')
					} else {
						parser.cursor--
						brace.Cts = append(brace.Cts, &cstr)
						break
					}
				} else {
					cstr.str = append(cstr.str, r)
				}
			}
		}
	}
	return nil, MissingRightBrace{}
}

type MissingRightSquareBracket struct {}

func (MissingRightSquareBracket) Error() string {
	return "error: missing closing squarebracket results in incomplete generator.\n"
}

func (parser *Parser) getToken() (*Token, error) {
	parser.skipBlank()
	var token Token
	token.Content = new([]rune)
	var err error
	var r rune
	r = parser.crvf[parser.cursor]
	parser.cursor++
	if r == '{' {
		left_brace_cnt := 1
		token.Ctype = SBR
		for parser.cursor < parser.crvf_len {
			r = parser.crvf[parser.cursor]
			parser.cursor++
			switch r {
			case '}':
				if !(parser.cursor>1 && parser.crvf[parser.cursor-2] == '\\') {
					left_brace_cnt--
				}
				if left_brace_cnt == 0 {
					return &token, nil
				} else {
					*token.Content = append(*token.Content, '}')
				}
			case '{':
				if !(parser.cursor>1 && parser.crvf[parser.cursor-2] == '\\') {
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
		for parser.cursor < parser.crvf_len {
			r = parser.crvf[parser.cursor]
			parser.cursor++
			*token.Content = append(*token.Content, r)
			if r == ' ' {
				*token.Content = (*token.Content)[:len(*token.Content)-1]
				break
			}  else if parser.meetNewLine(parser.cursor-1) {    // '\'跟个回车，会把下一行拼接到后面，而后忽略掉'\'
				if parser.crvf[parser.cursor-2] == '\\' {
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

type TooMuchSettings struct {}

func (TooMuchSettings) Error() string {
	return "error: there are too much Settings blocks in the .crvf.\n"
}

type FLParseError struct {}

func (FLParseError) Error() string {
	return "error: additional keywords that should not appear are encountered at the first level of the .crvf"
}


func (parser *Parser) Parse() error {
	Settings_cnt := 0
	var preBr *Br = nil
	for parser.cursor < parser.crvf_len {
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
		case "Br":
			lo, hi, err := parser.getRange()
			if err != nil {
				return err
			}
			if preBr != nil {
				preBr.rangelo += lo
				preBr.rangehi += hi
			} else {
				var br Br
				br.rangelo = lo
				br.rangehi = hi
				preBr = &br
			}
		case "{":
			if preBr != nil {
				var brbrace Brace
				brbrace.Cts = append(brbrace.Cts, preBr)
				brbrace.rangelo = 1
				brbrace.rangehi = 1

				parser.tpl.Braces = append(parser.tpl.Braces, brbrace)
				preBr = nil
			}
			bracep, err := parser.parseBrace()
			if err != nil {
				return err
			}
			parser.tpl.Braces = append(parser.tpl.Braces, *bracep)
		default:
			return FLParseError{}
		}
	}
	if preBr != nil {
		var brbrace Brace
		brbrace.Cts = append(brbrace.Cts, preBr)
		brbrace.rangelo = 1
		brbrace.rangehi = 1

		parser.tpl.Braces = append(parser.tpl.Braces, brbrace)
		preBr = nil
	}
	return nil
}

func (parser *Parser) RegisterTemplate() {
	if _, ok := Templates[parser.tpl.Name]; !ok {
		Templates[parser.tpl.Name] = &parser.tpl
	}
}