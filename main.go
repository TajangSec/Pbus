package main

import (
	"Pbus/home"
	"Pbus/preprocessing"
	"Pbus/scan"
	"flag"
)

func main() {
	if home.Help() || len(home.Url()) == 0 {
		flag.Usage()
		return
	}
	Url := preprocessing.UrlParse(home.Url())
	wd := preprocessing.WdParse(home.Wordlist())
	thread_num := preprocessing.ThreadParse(home.Thread())
	scan.Scan(Url, wd, thread_num)
}
