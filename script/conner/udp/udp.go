package udp

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type SuccessSR struct {
	Sent_t time.Time
	Recv_t []time.Time
	Datas [][]byte
}

type Udp struct {
	Name string
	IP    string
	Port  int
	Succs chan *SuccessSR
	Srsucccnt int
	Srcnt int
}

func (udp *Udp) st (ip string, port int) error {
	udp.IP = ip
	udp.Port = port
	udp.Succs = make(chan *SuccessSR, 100)
	return nil
}

func (udp *Udp) sr (wtimeout int, datars []rune, rtimeout int, rtimes int) error {
	ssr := SuccessSR{}
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", udp.IP, udp.Port))
	if err != nil {
		return err
	}
	err = conn.SetWriteDeadline(time.Now().Add(time.Duration(wtimeout)*time.Millisecond))
	if err != nil {
		return err
	}
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read([]byte(string(datars)))
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			return err
		}
	}
	ssr.Sent_t = time.Now()
	for rt:=0;rt<rtimes;rt++ {
		err = conn.SetReadDeadline(time.Now().Add(time.Duration(rtimeout)*time.Millisecond))
		if err != nil {
			return err
		}
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		ssr.Recv_t = append(ssr.Recv_t, time.Now())
		ssr.Datas = append(ssr.Datas, buf[:n])
	}
	udp.Succs <- &ssr
	return nil
}

