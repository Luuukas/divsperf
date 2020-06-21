package dnsanalyzer

import "fmt"

type UnknownDirective struct {
	Key_name string
	unknonw_dir string
}

func (ukd UnknownDirective) Error() string {
	return fmt.Sprintf("error: %s does not have directive named %s\n", ukd.Key_name, ukd.unknonw_dir)
}

type DuplicatedTarget struct {
	Key_name string
	target_name string
	Dptarget_name string
}

func (dpt DuplicatedTarget) Error() string {
	return fmt.Sprintf("error: the analyzer %s had been used to analyze %s, %s will cause duplicated target.\n", dpt.Key_name, dpt.target_name, dpt.Dptarget_name)
}