package udp

import (
	"fmt"
	"testing"
)

func TestUdp_sr(t *testing.T) {
	udp := Udp{}
	udp.st("127.0.0.1", 11110)

	datars := []rune("hello world!")
	err := udp.sr(500, datars, 500, 1)
	if err != nil {
		t.Error(err)
		return
	}
	ssrp := <- udp.Succs
	fmt.Println(ssrp.Sent_t)
	fmt.Println(ssrp.Recv_t[0])
	fmt.Println(string(ssrp.Datas[0]))
}