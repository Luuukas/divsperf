package parse

import "fmt"

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

type BraceCntSBError struct {}

func (BraceCntSBError) Error() string {
	return "error: there should be a [x ] or [x y ] following the brace.\n"
}

type MissingRightSquareBracket struct {}

func (MissingRightSquareBracket) Error() string {
	return "error: missing closing squarebracket results in incomplete generator.\n"
}

type LevelGettingError struct {}

func (LevelGettingError) Error() string {
	return "error: there should be a [Lv x ] following the loop range.\n"
}

type TooMuchSettings struct {}

func (TooMuchSettings) Error() string {
	return "error: there are too much Settings blocks in the .scri.\n"
}

type FLParseError struct {}

func (FLParseError) Error() string {
	return "error: additional keywords that should not appear are encountered at the first level of the .scri\n"
}

type BraceParseError struct {}

func (BraceParseError) Error() string {
	return "error: a character that does not conform to Addon syntax has been encountered.\n"
}

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

type TrytoUseAddonWithoutReturn struct {}

func (err TrytoUseAddonWithoutReturn) Error() string {
	return "error: trying to get content from an addon which without such ability.\n"
}