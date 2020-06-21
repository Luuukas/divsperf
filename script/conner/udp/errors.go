package udp

import "fmt"

type UnknownDirective struct {
	Key_name string
	unknonw_dir string
}

func (ukd UnknownDirective) Error() string {
	return fmt.Sprintf("error: %s does not have directive named %s\n", ukd.Key_name, ukd.unknonw_dir)
}

type DuplicatedUdp struct {
	Key_name string
	udp_name string
}

func (dudp DuplicatedUdp) Error() string {
	return fmt.Sprintf("the udp named %s already exists.\n", dudp.udp_name)
}

type CannotFindUdp struct {
	Key_name string
	udp_name string
}

func (cnfudp CannotFindUdp) Error() string {
	return fmt.Sprintf("the udp named %s not exists.\n", cnfudp.udp_name)
}