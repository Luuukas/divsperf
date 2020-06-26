package command

import (
	"bufio"
	crvf_parse "divsperf/randval/parse"
	scri_parse "divsperf/script/parse"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	_ "divsperf/randval/setup"
	_ "divsperf/script/setup"
)

func Processing() {
	for {
		fmt.Scan()
		// 从stdin中取内容直到遇到换行符，停止
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println("你输入的内容是：", strings.TrimSpace(input))
		words := strings.Fields(input)
		switch words[0] {
		case "readcrvf":
			if len(words) != 2 {
				fmt.Println("usage: readcrvf xxx.crvf")
				break
			}
			crvf_parser := crvf_parse.Parser{}
			err = crvf_parser.New(words[1])
			if err != nil {
				fmt.Println("fail to read the .crvf")
				break
			}
			err = crvf_parser.Parse()
			if err != nil {
				fmt.Println(err)
			}
		case "readscri":
			if len(words) != 2 {
				fmt.Println("usage: readscri xxx.scri")
				break
			}
			scri_parser := scri_parse.Parser{}
			err = scri_parser.New(words[1])
			if err != nil {
				fmt.Println("fail to read the .scri")
				break
			}
			err = scri_parser.Parse()
			if err != nil {
				fmt.Println(err)
			}
		case "run":
			if len(words) != 1 {
				fmt.Println("usage: run")
				break
			}
			RunScripts()
		default:
			fmt.Println("invalid command")
		}
		fmt.Print("> ")
	}
}

func RunScripts() {
	for l, LCp := range scri_parse.Levels {
		wg := &sync.WaitGroup{}
		for LCp != nil {
			go func() {
				wg.Add(1)
				defer wg.Done()
				rand.Seed(time.Now().Unix())
				rt := LCp.SittingB.Rangelo + rand.Intn(LCp.SittingB.Rangehi+1)
				for t:=0;t<rt;t++ {
					for _, sb := range LCp.SittingB.Sbs {
						err := sb.LetAction(wg)
						log.Printf("error: runscript - level: %d %s - %v",l,sb.Name,err)
					}
				}
			}()
			LCp = LCp.Next
		}
		wg.Wait()
	}
}