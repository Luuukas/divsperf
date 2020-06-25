package udp

import (
	_ "divsperf/randval/create/rand"
	crvf_parse "divsperf/randval/parse"
	_ "divsperf/script/addon/use"
	"divsperf/script/parse"
	"fmt"
	"sync"
	"testing"
)

// [udp st uuddpp 127.0.0.1 11110 ]
// [udp sr uuddpp 500 {hello, world!} 500 1 ]
//func TestUdpAddon_Action(t *testing.T) {
//	udpaddon := UdpAddon{}
//	udpaddon.Udps = make(map[string]*Udp)
//
//	st_tk1_content := []rune("st")
//	st_tk1 := parse.Token{Ctype: parse.STR, Content: &st_tk1_content}
//
//	st_tk2_content := []rune("uuddpp")
//	st_tk2 := parse.Token{Ctype: parse.STR, Content: &st_tk2_content}
//
//	st_tk3_content := []rune("127.0.0.1")
//	st_tk3 := parse.Token{Ctype: parse.STR, Content: &st_tk3_content}
//
//	st_tk4_content := []rune("11110")
//	st_tk4 := parse.Token{Ctype: parse.INT, Content: &st_tk4_content}
//
//	st_tks := []parse.Token{st_tk1, st_tk2, st_tk3, st_tk4}
//	st_sb := parse.SquareBrackets{Name: "udp", Tokens: &st_tks}
//
//	sr_tk1_content := []rune("sr")
//	sr_tk1 := parse.Token{Ctype: parse.STR, Content: &sr_tk1_content}
//
//	sr_tk2_content := []rune("uuddpp")
//	sr_tk2 := parse.Token{Ctype: parse.STR, Content: &sr_tk2_content}
//
//	sr_tk3_content := []rune("500")
//	sr_tk3 := parse.Token{Ctype: parse.INT, Content: &sr_tk3_content}
//
//	sr_tk4_content := []rune("hello, world!")
//	sr_tk4 := parse.Token{Ctype: parse.SBR, Content: &sr_tk4_content}
//
//	sr_tk5_content := []rune("500")
//	sr_tk5 := parse.Token{Ctype: parse.INT, Content: &sr_tk5_content}
//
//	sr_tk6_content := []rune("1")
//	sr_tk6 := parse.Token{Ctype: parse.INT, Content: &sr_tk6_content}
//
//	sr_tks := []parse.Token{sr_tk1, sr_tk2, sr_tk3, sr_tk4, sr_tk5, sr_tk6}
//	sr_sb := parse.SquareBrackets{Name: "udp", Tokens: &sr_tks}
//
//	wg := &sync.WaitGroup{}
//
//	err := udpaddon.Action(wg, &st_sb)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	err = udpaddon.Action(wg, &sr_sb)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	ssrp := <- udpaddon.Udps["uuddpp"].Succs
//	fmt.Println(ssrp.Sent_t)
//	fmt.Println(ssrp.Recv_t[0])
//	fmt.Println(string(ssrp.Datas[0]))
//}

// [udp st uuddpp 127.0.0.1 11110 ]
// [udp sr uuddpp 500 [use templatename ] 500 1 ]
func TestUdpAddon_Action2(t *testing.T) {
	udpaddon := UdpAddon{}
	udpaddon.Udps = make(map[string]*Udp)

	st_tk1_content := []rune("st")
	st_tk1 := parse.Token{Ctype: parse.STR, Content: &st_tk1_content}

	st_tk2_content := []rune("uuddpp")
	st_tk2 := parse.Token{Ctype: parse.STR, Content: &st_tk2_content}

	st_tk3_content := []rune("127.0.0.1")
	st_tk3 := parse.Token{Ctype: parse.STR, Content: &st_tk3_content}

	st_tk4_content := []rune("11110")
	st_tk4 := parse.Token{Ctype: parse.INT, Content: &st_tk4_content}

	st_tks := []parse.Token{st_tk1, st_tk2, st_tk3, st_tk4}
	st_sb := parse.SquareBrackets{Name: "udp", Tokens: &st_tks}

	sr_tk1_content := []rune("sr")
	sr_tk1 := parse.Token{Ctype: parse.STR, Content: &sr_tk1_content}

	sr_tk2_content := []rune("uuddpp")
	sr_tk2 := parse.Token{Ctype: parse.STR, Content: &sr_tk2_content}

	sr_tk3_content := []rune("500")
	sr_tk3 := parse.Token{Ctype: parse.INT, Content: &sr_tk3_content}

	crvf_parser := crvf_parse.Parser{}
	err := crvf_parser.New("test_udpaddon.crvf")
	if err != nil {
		t.Error(err)
		return
	}
	err = crvf_parser.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	crvf_parser.RegisterTemplate()

	var tplname []rune
	tplname = []rune("testudpaddon")
	tk := parse.Token{Ctype: parse.STR, Content: &tplname}
	var tks []parse.Token
	tks = append(tks, tk)
	use_sb := parse.SquareBrackets{Name: "use", Tokens: &tks}

	sr_tk4 := parse.Token{Ctype: parse.SSB, Sb: &use_sb}

	sr_tk5_content := []rune("500")
	sr_tk5 := parse.Token{Ctype: parse.INT, Content: &sr_tk5_content}

	sr_tk6_content := []rune("1")
	sr_tk6 := parse.Token{Ctype: parse.INT, Content: &sr_tk6_content}

	sr_tks := []parse.Token{sr_tk1, sr_tk2, sr_tk3, sr_tk4, sr_tk5, sr_tk6}
	sr_sb := parse.SquareBrackets{Name: "udp", Tokens: &sr_tks}

	wg := &sync.WaitGroup{}

	err = udpaddon.Action(wg, &st_sb)
	if err != nil {
		t.Error(err)
		return
	}
	err = udpaddon.Action(wg, &sr_sb)
	if err != nil {
		t.Error(err)
		return
	}
	ssrp := <- udpaddon.Udps["uuddpp"].Succs
	fmt.Println(ssrp.Sent_t)
	fmt.Println(ssrp.Recv_t[0])
	fmt.Println(string(ssrp.Datas[0]))
}